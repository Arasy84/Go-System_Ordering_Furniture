package schema

import (
    "time"
)

type Admin struct {
	ID             uint `gorm:"PrimaryKey"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}