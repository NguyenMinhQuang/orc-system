package example

import "orc-system/pkg/utils"

type GetByIDInput struct {
	ID int `json:"id" query:"id"`
}

func (i GetByIDInput) Validate() error {
	return utils.ValidateStruct(i)
}
