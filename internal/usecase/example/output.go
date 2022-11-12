package example

type GetByIDOutput struct {
	ID   int
	Name string
}

type GetAllUserOutput struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Sex      string `json:"sex"`
	Addr     string `json:"addr"`
}
