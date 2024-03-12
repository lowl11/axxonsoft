package proxy_controller

import (
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

	proxy.POST("/task", controller.Task)
}
