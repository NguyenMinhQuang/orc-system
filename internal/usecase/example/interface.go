package example

import "context"

type IUseCase interface {
	GetAllUser(ctx context.Context) ([]GetAllUserOutput, error)
	GetByID(ctx context.Context, param GetByIDInput) (GetByIDOutput, error)
}
