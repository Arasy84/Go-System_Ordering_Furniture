package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

func ProductSchemaToProductDomain(res *schema.Product) *domain.Product {
	return &domain.Product{
		ID: res.ID,
		Name: res.Name,
		Description: res.Description,
		Price: res.Price,
		Stock: res.Stock,
		Category: res.Category,
	}
}

func ProductDomainToProductResponse(product *domain.Product) modelsrespons.ProductResponse {
	return modelsrespons.ProductResponse{
		ID:       product.ID,
        Name:     product.Name,
        Description: product.Description,
        Price: product.Price,
	}
}

func ConvertProductResponse(products []domain.Product) []modelsrespons.ProductResponse {
	var results []modelsrespons.ProductResponse
	for _, product := range products {
		productResponse := modelsrespons.ProductResponse{
			ID:       		product.ID,
            Name:     		product.Name,
            Description: 	product.Description,
            Price: 			product.Price,
            Stock: 			product.Stock,
		}
		results = append(results, productResponse)
	}
	return results
}

func ProductUpdateRequestToProductDomain(product *domain.Product) modelsrespons.UpdateProduct {
	return modelsrespons.UpdateProduct{
		ID: 			product.ID,
		Name:         	product.Name,
        Description: 	product.Description,
        Price:         	product.Price,
		Stock:        	product.Stock,
		Category:     	product.Category,
	}
}

func AddProductRequestToProductResponse(product *domain.Product) modelsrespons.ProductCreate {
	return modelsrespons.ProductCreate{
		ID:             	product.ID,
        Name:             	product.Name,
        Description:     	product.Description,
        Price:             	product.Price,
        Stock:            	product.Stock,
        Category:         	product.Category,
	}
}