package repository

import (
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	res "furniture/utils/respons"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	FindByID(id int) (*domain.Order, error)
	FindAll() ([]domain.Order, error)
	Delete(id int) error
}

type OrderRepositoryImpl struct {
    DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &OrderRepositoryImpl{DB: db}
}

func (r *OrderRepositoryImpl) CreateOrder(order *domain.Order) (*domain.Order, error) {
	orderDb := request.OrderDomainToOrderSchema(*order)
	result := r.DB.Create(&orderDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.OrderSchemaToOrderDomain(orderDb)

	return results, nil
}

func (r *OrderRepositoryImpl) FindByID(id int) (*domain.Order, error) {
	var order domain.Order
	if err := r.DB.Unscoped().Where("id= ? AND deleted_at IS NULL", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepositoryImpl) FindAll() ([]domain.Order, error) {
	var order []domain.Order
	if err := r.DB.Where("deleted_at IS NULL").Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepositoryImpl) Delete(id int) error {
	result := r.DB.Delete(&schema.Order{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}