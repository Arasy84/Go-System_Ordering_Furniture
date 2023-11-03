package schema

import (
	"time"

	"gorm.io/gorm"
)

// Product adalah model data yang mewakili produk dalam database.
type Product struct {
	ID             uint `gorm:"primaryKey"`					// Kolom ID sebagai Primary Key
	Name           string `json:"name"`						// Nama produk
	Description    string `json:"description"`				// Deskripsi produk
	Price          float64 `gorm:"type:int" json:"price"`	// Harga produk
	Stock          int `json:"stock"`						// Stok produk
	Category	   string `json:"Category"`					// Kategori produk
	Image          string 	`json:"image"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`		// Waktu pembuatan rekaman
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`		// Waktu pembaruan rekaman
	DeletedAt      gorm.DeletedAt `gorm:"index"`			// Waktu penghapusan rekaman (soft delete)
}
