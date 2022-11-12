package example

import (
	"context"
	"orc-system/internal/repository/example"
	"orc-system/pkg/logger"
)

func (u *UseCase) GetByID(ctx context.Context, param GetByIDInput) (GetByIDOutput, error) {
	var (
		res GetByIDOutput
		err error
	)
	logger.Warn()
	// get data from DB
	input := example.GetExampleByIDInput{
		ID: param.ID,
	}
	data, err := u.ExampleRepo.GetExampleByID(input, ctx)
	if err != nil {
		return res, err
	}
	// get from 3thr
	// u.Service.GetList()

	// xu ly logic
	res.ID = data.ID
	res.Name = data.Name
	return res, nil
}
