package example

import (
	"orc-system/pkg/utils"
)

func (i GetExampleByIDInput) Validate() error {
	return utils.ValidateStruct(i)
}
