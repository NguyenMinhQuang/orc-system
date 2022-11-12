package example

import "context"

type IUseCase interface {
	GetByID(ctx context.Context, param GetByIDInput) (GetByIDOutput, error)
}
