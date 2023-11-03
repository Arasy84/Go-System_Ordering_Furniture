package service

import (
	"fmt"
	"furniture/helper"
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/repository"
	req "furniture/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type OrderService interface {
	CreateOrder(ctx echo.Context, request modelsrequest.CreateOrder) (*domain.Order, error)
	FindByID(ctx echo.Context, id int) (*domain.Order, error)
	FindAll(ctx echo.Context) ([]domain.Order, error)
	Delete(ctx echo.Context, id int) error
}

type OrderServiceImpl struct {
    OrderRepository repository.OrderRepository
	Validate 		*validator.Validate
}

func NewOrderService(OrderRepository repository.OrderRepository, validate *validator.Validate) OrderService {
    return &OrderServiceImpl{
		OrderRepository: OrderRepository,
		Validate:         validate,
	}
}

func (s *OrderServiceImpl) CreateOrder(ctx echo.Context, request modelsrequest.CreateOrder) (*domain.Order, error) {
	err := s.Validate.Struct(request)
    if err!= nil {
        return nil, helper.ValidationError(ctx, err)
    }

	order := req.OrderCreateRequestToOrderDomain(&request)

	result, err := s.OrderRepository.CreateOrder(order)
	if err!= nil {
        return nil, fmt.Errorf("error create order: %s", err.Error())
    }
	return result, nil
}

func (s *OrderServiceImpl) FindByID(ctx echo.Context, id int) (*domain.Order, error) {
	order, err := s.OrderRepository.FindByID(id)
	if err!= nil {
        return nil, err
    }
	return order, nil
}

func (s *OrderServiceImpl) FindAll(ctx echo.Context) ([]domain.Order, error) {
	order, err := s.OrderRepository.FindAll()
    if err!= nil {
        return nil, fmt.Errorf("orders not found")
    }
    return order, nil
}

func (s *OrderServiceImpl) Delete(ctx echo.Context, id int) error {
	existingProduct, _ := s.OrderRepository.FindByID(id)
	if existingProduct == nil {
        return fmt.Errorf("order not found")
    }
	err := s.OrderRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting order: %s", err)
	}

	return nil
}