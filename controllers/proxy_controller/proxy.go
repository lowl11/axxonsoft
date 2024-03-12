package proxy_controller

import (
	"axxonsoft/adapter"
	"axxonsoft/data/dto"
	"axxonsoft/data/event"
	"github.com/lowl11/boost"
)

func (controller Controller) Task(ctx boost.Context) error {
	task := dto.Task{}
	if err := ctx.Parse(&task); err != nil {
		return ctx.Error(err)
	}

	createdID, err := controller.proxy.Task(ctx.Context(), &task)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.CreatedID(*createdID)
}

func (controller Controller) GetByID(ctx boost.Context) error {
	taskID, err := ctx.Param("task-id").UUID()
	if err != nil {
		return ctx.Error(err)
	}

	task, err := controller.proxy.GetByID(ctx.Context(), taskID)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Ok(adapter.Task(task))
}

func (controller Controller) CallTask(ctx boost.EventContext) error {
	call := event.CallTask{}
	if err := ctx.Parse(&call); err != nil {
		return err
	}

	if err := controller.proxy.CallTask(ctx.Context(), call.TaskID); err != nil {
		return err
	}

	return nil
}
