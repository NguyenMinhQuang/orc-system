package example

import (
	"orc-system/config"
	rp "orc-system/internal/repository/example"
	sv "orc-system/internal/service/example"
)

type UseCase struct {
	Config      *config.Config
	ExampleRepo rp.IRepository
	Service     sv.IExample
}

func NewExampleUseCase(repo rp.IRepository, svEx sv.IExample) IUseCase {
	return &UseCase{
		Config:      config.GetConfig(),
		ExampleRepo: repo,
		Service:     svEx,
	}
}
