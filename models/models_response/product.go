package modelsrespons

// ProductResponse adalah model respons yang digunakan untuk mengirim data produk ke klien.
type ProductResponse struct {
	ID          uint    `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
	Image       string 	`json:"image"`
}

// ProductCreate adalah model respons yang digunakan untuk mengirim data produk yang berhasil dibuat ke klien.
type ProductCreate struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
	Image       string 	`json:"image"`
}

// UpdateProduct adalah model respons yang digunakan untuk mengirim data produk yang berhasil diperbarui ke klien.
type UpdateProduct struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
	Image       string 	`json:"image"`
}