package proxy_controller

import (
	"axxonsoft/data/event"
	"axxonsoft/services/proxy_service"
	"github.com/lowl11/boost"
)

type Controller struct {
	proxy *proxy_service.Service
}

func New(proxy *proxy_service.Service) *Controller {
	return &Controller{
		proxy: proxy,
	}
}

func (controller Controller) RegisterEndpoints(router boost.Router) {
	proxy := router.Group("/api/v1/proxy")

	proxy.GET("/task/:task-id", controller.GetByID)
	proxy.POST("/task", controller.Task)
}

func (controller Controller) BindEvents(listener boost.Listener) {
	listener.Bind(event.CallTask{}, controller.CallTask)
}
