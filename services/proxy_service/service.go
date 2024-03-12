package proxy_service

import (
	"axxonsoft/repositories/task_repo"
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/pkg/web/requests"
)

type Service struct {
	dispatcher boost.Dispatcher
	task       *task_repo.Repo
	client     *requests.Service
}

func New(dispatcher boost.Dispatcher, task *task_repo.Repo) *Service {
	return &Service{
		dispatcher: dispatcher,
		task:       task,
		client:     requests.New(),
	}
}
