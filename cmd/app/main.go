package main

import (
	"axxonsoft/controllers/proxy_controller"
	"axxonsoft/services/proxy_service"
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/pkg/system/di"
)

func main() {
	app := boost.New()

	di.Register[proxy_service.Service](proxy_service.New)

	di.MapControllers(
		proxy_controller.New,
	)

	app.Run(config.Get("port").String())
}
