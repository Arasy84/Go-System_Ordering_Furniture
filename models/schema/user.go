package schema

import (
	"time"

	"gorm.io/gorm"
)

// User adalah model data yang mewakili pengguna dalam database.
type User struct {
	ID 				uint `gorm:"PrimaryKey"` 			// Kolom ID sebagai Primary Key
	Name 			string `json:"name"`				// Nama pengguna
	Email 			string `json:"email"`				// Alamat email pengguna
	Password 		string `json:"password"`			// Kata sandi pengguna
	Address 		string `json:"address"`				// Alamat pengguna
	Phone 			uint `json:"phone"`					// Nomor telepon pengguna
	CreatedAt 		time.Time `gorm:"autocreateTime"`	// Waktu pembuatan rekaman
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`	// Waktu pembaruan rekaman
	DeletedAt 		gorm.DeletedAt `gorm:"index"`		// Waktu penghapusan rekaman (soft delete)
}