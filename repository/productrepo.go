package repository

import (
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	"furniture/utils/respons"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product *domain.Product) (*domain.Product, error)
	Update(product *domain.Product, id int) (*domain.Product, error)
	Delete(id int) error
	GetId(id int) (*domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetProductByCategory(category string) (*domain.Product, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}

func (Repository *ProductRepositoryImpl) AddProduct(product *domain.Product) (*domain.Product, error) {
	productDB := request.ProductDomainToProductSchema(*product)
	result := Repository.DB.Create(&productDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := respons.ProductSchemaToProductDomain(productDB)

	return results, nil
}

func (Repository *ProductRepositoryImpl) Update(product *domain.Product, id int) (*domain.Product, error) {
	result := Repository.DB.Table("Product").Where("id = ?", id).Updates(domain.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Category:    product.Category,})
	if result.Error != nil {
		return nil, result.Error
	}
	
	return product, nil
}

func (Repository *ProductRepositoryImpl) Delete(id int) error {
	result := Repository.DB.Delete(&schema.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (Repository *ProductRepositoryImpl) GetId(id int) (*domain.Product, error) {
	var product domain.Product
	if err := Repository.DB.Unscoped().Where("id = ? AND deleted_at IS NULL", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (Repository *ProductRepositoryImpl) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	result := Repository.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (Repository *ProductRepositoryImpl) GetProductByCategory(category string) (*domain.Product, error) {
    product := domain.Product{}
    result := Repository.DB.Where("category = ?", category).First(&product)
    if result.Error != nil {
        return nil, result.Error
    }
    return &product, nil
}
