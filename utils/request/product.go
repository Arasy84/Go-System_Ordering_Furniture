package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

func AddProductRequestToProductDomain(req modelsrequest.AddProductRequest) *domain.Product {
    return &domain.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
    }
}

func ProductUpdateRequestToProductDomain(req modelsrequest.ProductUpdateRequest) *domain.Product {
    return &domain.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
    }
}

func ProductDomainToProductSchema(req domain.Product) *schema.Product {
    return &schema.Product{
        Name: req.Name,
        Description: req.Description,
        Price: req.Price,
        Stock: req.Stock,
        Category: req.Category,
    }
}