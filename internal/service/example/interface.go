package service

type IExample interface {
	GetList(param *ExpInput) (*ExpOutPut, error)
}
