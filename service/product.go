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

type ProductService interface {
	AddProduct(service echo.Context, request modelsrequest.AddProductRequest) (*domain.Product, error)
	UpdateProduct(service echo.Context, request modelsrequest.ProductUpdateRequest, id int) (*domain.Product, error)
	DeleteProduct(service echo.Context, id int) error
	GetProductId(service echo.Context, id int) (*domain.Product, error)
	GetAllProduct(service echo.Context) ([]domain.Product, error)
    GetProductByCategory(service echo.Context, Category string) (*domain.Product, error)
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(ProductRepository repository.ProductRepository, Validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
		Validate:          Validate,
	}
}

func (s *ProductServiceImpl) AddProduct(service echo.Context, request modelsrequest.AddProductRequest) (*domain.Product, error) {
    err := s.Validate.Struct(request)
    if err != nil {
        return nil, helper.ValidationError(service, err)
    }

    product := req.AddProductRequestToProductDomain(request)

    result, err := s.ProductRepository.AddProduct(product)
    if err != nil {
        return nil, fmt.Errorf("error adding product: %s", err.Error())
    }

    return result, nil
}

func (s *ProductServiceImpl) UpdateProduct(service echo.Context, request modelsrequest.ProductUpdateRequest, id int) (*domain.Product, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}

	existingProduct, _ := s.ProductRepository.GetId(id)
	if existingProduct == nil {
		return nil, fmt.Errorf("product not found")
	}

	product := req.ProductUpdateRequestToProductDomain(request)

	result, err := s.ProductRepository.Update(product, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating product: %s", err.Error())
	}
	return result, nil
}

func (s *ProductServiceImpl) DeleteProduct(service echo.Context, id int) error {
    existingProduct, _ := s.ProductRepository.GetId(id)
    if existingProduct == nil {
        return fmt.Errorf("product not found")
    }

    err := s.ProductRepository.Delete(id)
    if err!= nil {
        return fmt.Errorf("error when deleting product: %s", err)
    }

    return nil
}

func (s *ProductServiceImpl) GetProductId(service echo.Context, id int) (*domain.Product, error) {
    product, err := s.ProductRepository.GetId(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetAllProduct(service echo.Context) ([]domain.Product, error) {
    product, err := s.ProductRepository.GetAll()
    if err != nil {
        return nil, fmt.Errorf("products not found")
    }
	
    return product, nil
}

func (s *ProductServiceImpl) GetProductByCategory(service echo.Context, Category string) (*domain.Product, error) {
	product, _ := s.ProductRepository.GetProductByCategory(Category)
	if product == nil {
		return nil, fmt.Errorf("category not found")
	}

	return product, nil
}