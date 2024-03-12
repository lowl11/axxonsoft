package proxy_service

import (
	"axxonsoft/data/dto"
	"context"
	"github.com/lowl11/boost/errors"
)

func (service Service) Task(ctx context.Context, task *dto.Task) (*dto.TaskResponse, error) {
	response, err := service.client.
		R().
		SetContext(ctx).
		Do(task.Method, task.URL)
	if err != nil {
		return nil, errors.
			New("Do task error").
			SetType("Task_Do").
			SetError(err).
			SetContext(map[string]any{
				"method":  task.Method,
				"url":     task.URL,
				"headers": task.Headers,
			})
	}

	return &dto.TaskResponse{
		StatusCode:  response.StatusCode(),
		ContentType: response.ContentType(),
		Body:        response.Body(),
	}, nil
}
