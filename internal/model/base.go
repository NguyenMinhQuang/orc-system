package model

type ResponseModel struct {
	Err         error       `json:"err"`
	Message     string      `json:"message"`
	ResultCount int         `json:"result_count"`
	Data        interface{} `json:"data"`
}
