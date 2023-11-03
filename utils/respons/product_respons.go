package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

// ProductSchemaToProductDomain mengonversi objek schema.Product menjadi domain.Product.
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

// ProductDomainToProductResponse mengonversi objek domain.Product menjadi modelsrespons.ProductResponse.
func ProductDomainToProductResponse(product *domain.Product) modelsrespons.ProductResponse {
	return modelsrespons.ProductResponse{
		ID:       product.ID,
        Name:     product.Name,
        Description: product.Description,
        Price: product.Price,
		Stock: product.Stock,
        Category: product.Category,
		Image: product.Image,
	}
}

// ConvertProductResponse mengonversi slice dari objek domain.Product menjadi slice dari modelsrespons.ProductResponse.
func ConvertProductResponse(products []domain.Product) []modelsrespons.ProductResponse {
	var results []modelsrespons.ProductResponse
	for _, product := range products {
		productResponse := modelsrespons.ProductResponse{
			ID:       		product.ID,
            Name:     		product.Name,
            Description: 	product.Description,
            Price: 			product.Price,
            Stock: 			product.Stock,
			Category:       product.Category,
			Image:             product.Image,
		}
		results = append(results, productResponse)
	}
	return results
}

// ProductUpdateRequestToProductDomain mengonversi objek domain.Product menjadi modelsrespons.UpdateProduct.
func ProductUpdateRequestToProductDomain(product *domain.Product) modelsrespons.UpdateProduct {
	return modelsrespons.UpdateProduct{
		ID: 			product.ID,
		Name:         	product.Name,
        Description: 	product.Description,
        Price:         	product.Price,
		Stock:        	product.Stock,
		Category:     	product.Category,
		Image:            product.Image,
	}
}

// AddProductRequestToProductResponse mengonversi objek domain.Product menjadi modelsrespons.ProductCreate.
func AddProductRequestToProductResponse(product *domain.Product) modelsrespons.ProductCreate {
	return modelsrespons.ProductCreate{
		ID:             	product.ID,
        Name:             	product.Name,
        Description:     	product.Description,
        Price:             	product.Price,
        Stock:            	product.Stock,
        Category:         	product.Category,
		Image:             product.Image,
	}
}