package modelsrespons

type OrderCreateResponse struct {
	ID         uint `json:"id"`
	ProductID    uint `json:"product_id"`
	Product    string `json:"product"`
	UserID     uint `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}