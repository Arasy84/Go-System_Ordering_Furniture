package schema

import "time"

type Order struct {
	ID      		uint			 	`gorm:"PrimaryKey"`
	UserID        	uint			 	`json:"user_id"`
	ProductID     	uint				`json:"product_id"`
	Product         string              `json:"product"`
	Quantity      	int					`json:"quantity"`
	TotalPrice    	float64				`json:"total_price"`
	CreatedAt     	time.Time			`gorm:"autoCreateTime"`
	UpdatedAt     	time.Time			`gorm:"autoUpdate"`
	DeletedAt     	*time.Time			`gorm:"index"`
}