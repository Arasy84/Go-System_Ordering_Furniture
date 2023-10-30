package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID 				uint `gorm:"PrimaryKey"`
	Name 			string `json:"name"`
	Email 			string `json:"email"`
	Password 		string `json:"password"`
	Address 		string `json:"address"`
	Phone 			uint `json:"phone"`
	CreatedAt 		time.Time `gorm:"autocreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
}