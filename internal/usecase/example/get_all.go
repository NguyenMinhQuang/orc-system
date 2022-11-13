package example

import (
	"context"
	"errors"
)

func (u *UseCase) GetAllUser(ctx context.Context) ([]GetAllUserOutput, error) {
	var res = make([]GetAllUserOutput, 0)
	uL, err := u.ExampleRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	for _, val := range uL {
		dt := GetAllUserOutput{
			ID:       val.ID,
			UserName: val.UserName,
			Sex:      val.Sex,
			Addr:     val.Addr,
		}
		res = append(res, dt)
	}

	return res, errors.New("todo some thing")
}
