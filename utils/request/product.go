package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

// AddProductRequestToProductDomain mengonversi objek modelsrequest.AddProductRequest menjadi domain.Product.
func AddProductRequestToProductDomain(req modelsrequest.AddProductRequest) *domain.Product {
    return &domain.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
        Image: req.Image,
    }
}

// ProductUpdateRequestToProductDomain mengonversi objek modelsrequest.ProductUpdateRequest menjadi domain.Product.
func ProductUpdateRequestToProductDomain(req modelsrequest.ProductUpdateRequest) *domain.Product {
    return &domain.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
        Image: req.Image,
    }
}

// ProductDomainToProductSchema mengonversi objek domain.Product menjadi schema.Product.
func ProductDomainToProductSchema(req domain.Product) *schema.Product {
    return &schema.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
        Image: req.Image,
    }
}