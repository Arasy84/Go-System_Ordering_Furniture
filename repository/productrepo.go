package repository

import (
	// "errors"
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	"furniture/utils/respons"

	"gorm.io/gorm"
)

// ProductRepository adalah interface yang mendefinisikan operasi yang dapat dilakukan pada entitas produk (Product).
type ProductRepository interface {
	AddProduct(product *domain.Product) (*domain.Product, error)
	Update(product *domain.Product, id int) (*domain.Product, error)
	Delete(id int) error
	GetId(id int) (*domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetProductByCategory(category string) (*domain.Product, error)
}

// ProductRepositoryImpl adalah implementasi dari ProductRepository.
type ProductRepositoryImpl struct {
	DB *gorm.DB
	// store map[uint]*domain.User
}

// NewProductRepository adalah konstruktor untuk membuat instance ProductRepositoryImpl.
func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: DB,
		// store map[uint]*domain.User
	}
}

// AddProduct digunakan untuk menambahkan produk baru ke database.
func (Repository *ProductRepositoryImpl) AddProduct(product *domain.Product) (*domain.Product, error) {
	productDB := request.ProductDomainToProductSchema(*product)
	result := Repository.DB.Create(&productDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := respons.ProductSchemaToProductDomain(productDB)

	return results, nil
}

// Update digunakan untuk memperbarui data produk berdasarkan ID.
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

// Delete digunakan untuk menghapus data produk berdasarkan ID.
func (Repository *ProductRepositoryImpl) Delete(id int) error {
	result := Repository.DB.Delete(&schema.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetById digunakan untuk mendapatkan produk berdasarkan ID.
func (Repository *ProductRepositoryImpl) GetId(id int) (*domain.Product, error) {
	var product domain.Product
	if err := Repository.DB.Unscoped().Where("id = ? AND deleted_at IS NULL", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll digunakan untuk mendapatkan semua produk yang tersedia.
func (Repository *ProductRepositoryImpl) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	result := Repository.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// GetProductByCategory digunakan untuk mendapatkan produk berdasarkan kategori.
func (Repository *ProductRepositoryImpl) GetProductByCategory(category string) (*domain.Product, error) {
    product := domain.Product{}
    result := Repository.DB.Where("category = ?", category).First(&product)
    if result.Error != nil {
        return nil, result.Error
    }
    return &product, nil
}

