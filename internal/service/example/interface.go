package example

type IExample interface {
	GetList(param *ExpInput) (*ExpOutPut, error)
}
