package routes

import (
	"furniture/helper"
	"furniture/hendler"
	"furniture/repository"
	"furniture/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRoute(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	// Membuat repository untuk entitas Produk.
	ProductRepository := repository.NewProductRepository(db)

	// Membuat layanan (service) untuk entitas Produk dan menginisialisasi dengan repository dan validator.
	serviceProduct := service.NewProductService(ProductRepository, validate)

	// Membuat handler untuk entitas Produk dan menginisialisasi dengan layanan.
	productHendler := hendler.NewHandlerProduct(serviceProduct)

	// Membuat grup rute untuk entitas Produk di Echo.
	productsGroup := e.Group("")

	// Middleware: Menggunakan JWT untuk otentikasi.
	productsGroup.GET("/product", productHendler.ProductGetAll)
	productsGroup.GET("/product/:id", productHendler.ProductGetById)
	productsGroup.GET("/product/category/:category", productHendler.GetByCategory)
	
	productsGroup.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	// productsGroup.GET("/products", productHendler.GetProducts)
	// productsGroup.POST("/recommendation", productHendler.GetRecomendAi)

	// Rute POST untuk menambahkan produk baru (memerlukan otentikasi admin).
	productsGroup.POST("/admin/product", productHendler.AddProduct, helper.AuthMiddleware("admin"))

	// Rute PUT untuk mengubah produk (memerlukan otentikasi admin).
	productsGroup.PUT("/admin/product/:id", productHendler.ProductUpdate, helper.AuthMiddleware("admin"))
	
	// Rute DELETE untuk menghapus produk (memerlukan otentikasi admin).
	productsGroup.DELETE("/admin/product/:id", productHendler.ProductDelete, helper.AuthMiddleware("admin"))
}

