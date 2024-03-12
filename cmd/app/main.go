package main

import (
	"axxonsoft/controllers/proxy_controller"
	"axxonsoft/repositories/task_repo"
	"axxonsoft/services/proxy_service"
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/config"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/storage"
	"github.com/lowl11/boost/storage/sql"
	"time"
)

func main() {
	app := boost.New()

	// database
	storage.RegisterConnect(
		app.Context(),
		config.Get("database").String(),
		storage.WithMaxConnections(10),
		storage.WithMaxIdleConnections(10),
		storage.WithMaxLifetime(time.Second*10),
		storage.WithMaxIdleLifetime(time.Second*10),
	)
	storage.MustPing()
	sql.EnableLog()

	// dispatcher (RMQ)
	boost.RegisterDispatcher(config.Get("amqp").String())

	// proxy
	di.Register[task_repo.Repo](task_repo.New)
	di.Register[proxy_service.Service](proxy_service.New)

	di.MapControllers(
		proxy_controller.New,
	)

	di.Get[proxy_controller.Controller]().BindEvents(app.Listener())

	go app.RunListener(config.Get("amqp").String())
	app.Run(config.Get("port").String())
}
