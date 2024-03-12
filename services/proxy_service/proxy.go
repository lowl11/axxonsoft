package proxy_service

import (
	"axxonsoft/adapter"
	"axxonsoft/data/dto"
	"axxonsoft/data/entity"
	"axxonsoft/data/enums/statuses"
	"axxonsoft/data/event"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/storage/sql"
	"net/http"
	"time"
)

func (service Service) Task(ctx context.Context, task *dto.Task) (*uuid.UUID, error) {
	var requestBody []byte
	var err error

	if task.Body != nil {
		requestBody, err = json.Marshal(task.Body)
		if err != nil {
			return nil, errors.
				New("Parse request body error").
				SetType("Proxy_ParseRequestBody").
				SetError(err)
		}
	}

	taskID := uuid.New()
	if err = service.task.Add(ctx, entity.Task{
		Entity: sql.Entity{
			ID:        taskID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Status: statuses.New,

		RequestMethod:  task.Method,
		RequestURL:     task.URL,
		RequestHeaders: adapter.JoinHeaders(task.Headers),
		RequestBody:    requestBody,
	}); err != nil {
		return nil, errors.
			New("Add new task error").
			SetType("Proxy_AddTask").
			SetError(err)
	}

	if err = service.dispatcher.Dispatch(ctx, event.CallTask{
		TaskID: taskID,
	}); err != nil {
		return nil, err
	}

	return &taskID, nil
}

func (service Service) CallTask(ctx context.Context, taskID uuid.UUID) error {
	// get task, throws Not Found
	task, err := service.GetByID(ctx, taskID)
	if err != nil {
		return err
	}

	// update status
	task.Status = statuses.InProgress
	if err = service.task.Update(ctx, *task); err != nil {
		return errors.
			New("Update task status error").
			SetType("Task_UpdateStatus").
			SetError(err).
			AddContext("task-id", taskID)
	}

	// call third party service
	response, err := service.client.
		R().
		SetContext(ctx).
		SetHeaders(adapter.SplitHeaders(task.RequestHeaders)).
		Do(task.RequestMethod, task.RequestURL)
	if err != nil {
		errMessage := errors.
			New("Do task error").
			SetType("Task_Do").
			SetError(err).
			SetContext(map[string]any{
				"method":  task.RequestMethod,
				"url":     task.RequestURL,
				"headers": task.RequestHeaders,
			}).Error()

		task.ResponseErrorReason = &errMessage
		task.Status = statuses.Error
	} else {
		joinedHeaders := adapter.JoinHeaders(response.Headers())
		bodyLen := len(response.Body())
		status := response.Status()
		statusCode := response.StatusCode()

		task.ResponseBody = response.Body()
		task.ResponseHeaders = &joinedHeaders
		task.ResponseLength = &bodyLen
		task.ResponseStatus = &status
		task.ResponseStatusCode = &statusCode

		task.Status = statuses.Done
	}

	// update task response info
	if err = service.task.Update(ctx, *task); err != nil {
		return errors.
			New("Update task response info error").
			SetType("Task_UpdateResponseInfo").
			SetError(err).
			AddContext("task-id", taskID)
	}

	return nil
}

func (service Service) GetByID(ctx context.Context, taskID uuid.UUID) (*entity.Task, error) {
	task, err := service.task.FindByID(ctx, taskID)
	if err != nil {
		return nil, errors.
			New("Get task by id error").
			SetType("Task_GetByID").
			SetError(err).
			AddContext("task-id", task)
	}

	if task == nil {
		return nil, errors.
			New("Task not found").
			SetType("Task_NotFound").
			SetHttpCode(http.StatusNotFound).
			AddContext("task-id", taskID)
	}

	return task, nil
}
