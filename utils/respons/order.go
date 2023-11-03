package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

type OrderResponse struct {
	ID              uint                 `json:"id"`
	UserID          uint                 `json:"user_id"`
	ProductID       uint                 `json:"product_id"`
	Quantity        int                  `json:"quantity"`
	TotalPrice      float64              `json:"total_price"`

}

func ConvertSchemaToResponse(order *schema.Order) *domain.Order {
	return &domain.Order{
        ID:              order.ID,
        UserID:            order.UserID,
        ProductID:         order.ProductID,
		Product:    		order.Product,
        Quantity:          order.Quantity,
        TotalPrice:        order.TotalPrice,
    }
}

func OrderSchemaToOrderDomain(Order *schema.Order) *domain.Order {
	return &domain.Order{
        ID:              Order.ID,
        UserID:            Order.UserID,
        ProductID:         Order.ProductID,
		Product:    	Order.Product,
        Quantity:          Order.Quantity,
        TotalPrice:        Order.TotalPrice,
    }
}

func CreateOrderToOrderResponse(Order *domain.Order)  modelsrespons.OrderCreateResponse {
	return modelsrespons.OrderCreateResponse {
		ID:              Order.ID,
        UserID:          Order.UserID,
        ProductID:       Order.ProductID,
		Product:    	 Order.Product,
        Quantity:        Order.Quantity,
        TotalPrice:      Order.TotalPrice,
	}
}

func ConvertOrderResponse(orders []domain.Order) []modelsrespons.OrderCreateResponse {
	var results []modelsrespons.OrderCreateResponse
	for _, Order := range orders {
		orderResponse := modelsrespons.OrderCreateResponse{
			ID:         Order.ID,
			UserID:     Order.UserID,
			ProductID:  Order.ProductID,
			Product:    Order.Product,
            Quantity:   Order.Quantity,
            TotalPrice: Order.TotalPrice,
		}
		results = append(results, orderResponse)
	}
	return results
}