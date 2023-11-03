package schema

import (
    "time"
)

// Admin adalah model data yang mewakili admin dalam database.
type Admin struct {
	ID             uint `gorm:"PrimaryKey"`			// Kolom ID sebagai Primary Key
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}