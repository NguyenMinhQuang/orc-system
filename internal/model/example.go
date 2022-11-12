package model

import "time"

type Example struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
