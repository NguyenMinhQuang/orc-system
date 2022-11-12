package service

import (
	"orc-system/config"
	"orc-system/internal/service/example"
)

type Service struct {
	Example example.IExample
}

func NewService() *Service {
	cfg := config.GetConfig()
	return &Service{
		Example: example.NewExampleService(cfg.EndPoint),
	}
}
