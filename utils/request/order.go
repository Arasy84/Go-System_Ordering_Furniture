package request

import (
	modelsrequest "furniture/models/models_request"
	"furniture/models/domain"
	"furniture/models/schema"
)

type OrderCreateRequest struct {
	UserID 			  uint		  `json:"user_id"`
	ProductID         uint        `json:"product_id"`
	Quantity          int         `json:"quantity" form:"quantity" validate:"required,min=1"`
	TotalPrice        float64     `json:"total_price" form:"total_price" validate:"required,min=1"`
}

func OrderDomainToOrderSchema(Order domain.Order) *schema.Order {
	return &schema.Order{
        ID:                Order.ID,
        UserID:            Order.UserID,
        ProductID:         Order.ProductID,
		Product:        	Order.Product,
        Quantity:          Order.Quantity,
        TotalPrice:        Order.TotalPrice,
	}
}

func OrderCreateRequestToOrderDomain(request *modelsrequest.CreateOrder) *domain.Order {
	return &domain.Order{
        UserID:            request.UserID,
        ProductID:         request.ProductID,
		Product:        request.Product,
        Quantity:          request.Quantity,
        TotalPrice:        request.TotalPrice,
    }
}