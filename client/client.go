package client

import (
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func (e *Client) DB() *gorm.DB {
	return e.db
}

func NewClient(db *gorm.DB) *Client {
	return &Client{
		db: db,
	}
}
