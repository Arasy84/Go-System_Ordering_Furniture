package domain

type Order struct {
	ID         uint `gorm:"primaryKey"`
	ProductID    uint
	Product    string
	UserID     uint
	Quantity   int
	TotalPrice float64
}
