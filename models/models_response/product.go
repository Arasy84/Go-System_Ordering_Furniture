package modelsrespons

type ProductResponse struct {
	ID          uint    `json:"id" form:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
}

type ProductCreate struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
}

type UpdateProduct struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Stock       int     `json:"stock" form:"stock"`
	Category    string  `json:"Category" form:"category"`
}