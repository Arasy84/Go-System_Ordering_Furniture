package schema

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             uint `gorm:"primaryKey"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          float64 `gorm:"type:int" json:"price"`
	Stock          int `json:"stock"`
	Category	   string `json:"Category"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}