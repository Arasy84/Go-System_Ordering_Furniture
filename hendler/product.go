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

// ProductHandler adalah interface yang mendefinisikan operasi yang dapat dilakukan pada entitas produk (product).
type ProductHandler interface {
	AddProduct(ctx echo.Context) error
	ProductUpdate(ctx echo.Context) error
	ProductDelete(ctx echo.Context) error
	ProductGetById(ctx echo.Context) error
	ProductGetAll(ctx echo.Context) error
	GetByCategory(ctx echo.Context) error
}

// HandlerProduct adalah implementasi dari ProductHandler.
type HandlerProduct struct {
	Service service.ProductService
}

// NewHandlerProduct adalah konstruktor untuk membuat instance HandlerProduct.
func NewHandlerProduct(productService service.ProductService) HandlerProduct {
	return HandlerProduct{Service: productService}
}

// AddProduct digunakan untuk menambahkan produk baru.
func (h *HandlerProduct) AddProduct(ctx echo.Context) error {
	// Bind request ke objek AddProductRequest
	AddProductRequest := modelsrequest.AddProductRequest{}
	err := ctx.Bind(&AddProductRequest)
	if err != nil {
		 // Jika binding gagal, kirim respons dengan status 400 (Bad Request)
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	// Coba menambahkan produk baru ke database
	result, err := h.Service.AddProduct(ctx, AddProductRequest)
	if err!= nil {
		if strings.Contains(err.Error(), "validation failed") {
			 // Jika validasi gagal, kirim respons dengan status 400 (Bad Request)
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
        }
		// Jika terjadi kesalahan lain, kirim respons dengan status 500 (Internal Server Error)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Product Creation Failed"))
	}
	 // Konversi hasil operasi ke respons produk dan kirim respons dengan status 201 (Created)
	response := res.AddProductRequestToProductResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully created product", response))
}

// ProductUpdate digunakan untuk memperbarui data produk.
func (h *HandlerProduct) ProductUpdate(ctx echo.Context) error {
	// Ambil ID produk dari parameter URL
	productId := ctx.Param("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		// Jika ID tidak valid, kirim respons dengan status 400 (Bad Request)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	 // Bind request ke objek ProductUpdateRequest
	productUpdateRequest := modelsrequest.ProductUpdateRequest{}
	err = ctx.Bind(&productUpdateRequest)
	if err != nil {
		// Jika binding gagal, kirim respons dengan status 400 (Bad Request)
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	 // Coba memperbarui produk berdasarkan ID
	result, err := h.Service.UpdateProduct(ctx, productUpdateRequest, productIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			// Jika validasi gagal, kirim respons dengan status 400 (Bad Request)
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		
		if strings.Contains(err.Error(), "Product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Product failed"))
	}

	// Konversi hasil operasi ke respons produk yang diperbarui dan kirim respons dengan status 201 (Created)
	response := res.ProductUpdateRequestToProductDomain(result)
	fmt.Print(result)
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully updated product", response))
}

// ProductDelete digunakan untuk menghapus data produk.
func (h *HandlerProduct) ProductDelete(ctx echo.Context) error {
	// Ambil ID produk dari parameter URL
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		// Jika ID tidak valid, kirim respons dengan status 400 (Bad Request)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	// Coba menghapus produk berdasarkan ID
	err = h.Service.DeleteProduct(ctx, ProductIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Product not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Product Failed"))
	}

	// Jika penghapusan produk berhasil, kirim respons dengan status 201 (Created)
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully deleted product", nil))
}

// ProductGetById digunakan untuk mendapatkan data produk berdasarkan ID.
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

// ProductGetAll digunakan untuk mendapatkan semua data produk.
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

	// Mengambil produk berdasarkan kategori dengan memanggil layanan ProductService.
	result, err := h.Service.GetProductByCategory(ctx, category)
	if err != nil {
		if strings.Contains(err.Error(), "category Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("category Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(("Get Product by Category error")))
	}

	 // Mengonversi hasil yang diperoleh menjadi respons produk dan mengembalikannya dalam respons JSON.
	response := res.ProductDomainToProductResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Product by Category", response))
}

