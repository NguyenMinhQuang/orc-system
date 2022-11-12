package model

import "time"

type Example struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Sex      string `json:"sex"`
	Addr     string `json:"addr"`
}
