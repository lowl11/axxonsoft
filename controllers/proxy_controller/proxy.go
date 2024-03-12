package proxy_controller

import (
	"axxonsoft/data/dto"
	"github.com/lowl11/boost"
)

func (controller Controller) Task(ctx boost.Context) error {
	task := dto.Task{}
	if err := ctx.Parse(&task); err != nil {
		return ctx.Error(err)
	}

	response, err := controller.proxy.Task(ctx.Context(), &task)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.
		SetContentType(response.ContentType).
		Status(response.StatusCode).
		Bytes(response.Body)
}
