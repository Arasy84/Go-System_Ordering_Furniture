package modelsrequest

type CreateOrder struct {
	UserID     uint    `json:"user_id" form:"userid"`
	ProductID    uint    `json:"product_id" form:"productid"`
	Product		string  `json:"product" form:"product"`
	Quantity   int     `json:"quantity" form:"quantity" validate:"required,min=1"`
	TotalPrice float64 `json:"totalPrice" form:"totalPrice" validate:"required,min=1"`
}