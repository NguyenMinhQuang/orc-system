package example

import (
	"context"
	"orc-system/internal/model"
)

type (
	GetExampleByIDInput struct {
		ID int
	}
)

type IRepository interface {
	GetExampleByID(param GetExampleByIDInput, ctx context.Context) (model.Example, error)
	GetAllUser(ctx context.Context) ([]model.User, error)
}
