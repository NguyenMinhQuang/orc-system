package codetype

import "strings"

type StageType string

const (
	StageTypeProd    StageType = "PROD"
	StageTypeStaging StageType = "STG"
	StageTypeDev     StageType = "DEV"
	StageTypeLocal   StageType = "LOCAL"
)

func (s *StageType) UpCase() {
	*s = StageType(strings.ToUpper(string(*s)))
}

func (s *StageType) ToString() string {
	return string(*s)
}
