package modelsrequest

type AddProductRequest struct {
    Name        string  `json:"name" form:"name" validate:"required,min=1,max=255"`
    Description string  `json:"description" form:"description" validate:"required,min=1,max=255"`
    Price       float64 `json:"price" form:"price" validate:"required,min=1"`
    Stock       int     `json:"stock" form:"stock" validate:"required,min=1,max=255"`
    Category    string  `json:"Category" form:"category" validate:"required,min=1,max=255"`
	Image       string 	`json:"image" form:"image"`
}


type ProductUpdateRequest struct {
	Name        string  `json:"name" form:"name" validate:"required,min=1,max=255"`
	Description string  `json:"description" form:"description" validate:"required,min=1,max=255"`
	Price       float64 `json:"price" form:"price" validate:"required,min=1"`
	Stock       int     `json:"stock" form:"stock" validate:"required,min=1,max=255"`
	Category    string  `json:"category" form:"category" validate:"required,min=1,max=255"`
	Image       string 	`json:"image" form:"image"`
}
