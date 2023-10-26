package schema

import (
    "time"
)

type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          float64
	Stock          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}