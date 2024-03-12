package proxy_service

import "github.com/lowl11/boost/pkg/web/requests"

type Service struct {
	client *requests.Service
}

func New() *Service {
	return &Service{
		client: requests.New(),
	}
}
