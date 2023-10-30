package hendler

import (
	"fmt"
	"furniture/helper"
	modelsrequest "furniture/models/models_request"
	"furniture/service"
	res "furniture/utils/respons"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	AddProduct(ctx echo.Context) error
	ProductUpdate(ctx echo.Context) error
	ProductDelete(ctx echo.Context) error
	ProductGetById(ctx echo.Context) error
	ProductGetAll(ctx echo.Context) error
	GetByCategory(ctx echo.Context) error
}

type HandlerProduct struct {
	Service service.ProductService
}

func NewHandlerProduct(productService service.ProductService) HandlerProduct {
	return HandlerProduct{Service: productService}
}

func (h *HandlerProduct) AddProduct(ctx echo.Context) error {
	AddProductRequest := modelsrequest.AddProductRequest{}
	err := ctx.Bind(&AddProductRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.AddProduct(ctx, AddProductRequest)
	if err!= nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
        }
		
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Product Creation Failed"))
	}
	response := res.AddProductRequestToProductResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully created product", response))
}

func (h *HandlerProduct) ProductUpdate(ctx echo.Context) error {
	productId := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	productUpdateRequest := modelsrequest.ProductUpdateRequest{}
	err = ctx.Bind(&productUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.UpdateProduct(ctx, productUpdateRequest, productIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		
		if strings.Contains(err.Error(), "Product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Product failed"))
	}

	response := res.ProductUpdateRequestToProductDomain(result)
	fmt.Print(result)
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully updated product", response))
}

func (h *HandlerProduct) ProductDelete(ctx echo.Context) error {
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.Service.DeleteProduct(ctx, ProductIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Product Failed"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully deleted product", nil))
}

func (h *HandlerProduct) ProductGetById(ctx echo.Context) error {
	productId := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.Service.GetProductId(ctx, productIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Id Product Error"))
	}

	response := res.AddProductRequestToProductResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfuly Get Product Id", response))
}

func (h *HandlerProduct) ProductGetAll(ctx echo.Context) error {
	result, err := h.Service.GetAllProduct(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Product Error"))
	}
	response := res.ConvertProductResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfuly Get All Product Id", response))
}

func (h *HandlerProduct) GetByCategory(ctx echo.Context) error {
	category := ctx.Param("category")

	result, err := h.Service.GetProductByCategory(ctx, category)
	if err != nil {
		if strings.Contains(err.Error(), "category Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("category Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(("Get Product by Category error")))
	}

	response := res.ProductDomainToProductResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Product by Category", response))
}
