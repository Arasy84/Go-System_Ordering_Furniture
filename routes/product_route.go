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
	ProductRepository := repository.NewProductRepository(db)
	serviceProduct := service.NewProductService(ProductRepository, validate)
	productHendler := hendler.NewHandlerProduct(serviceProduct)

	productsGroup := e.Group("")
	productsGroup.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	
	productsGroup.GET("/product", productHendler.ProductGetAll)
	productsGroup.GET("/product/:id", productHendler.ProductGetById)
	productsGroup.GET("/product/category/:category", productHendler.GetByCategory)
	productsGroup.POST("/admin/product", productHendler.AddProduct, helper.AuthMiddleware("admin"))
	productsGroup.PUT("/admin/product/:id", productHendler.ProductUpdate, helper.AuthMiddleware("admin"))
	productsGroup.DELETE("/admin/product/:id", productHendler.ProductDelete, helper.AuthMiddleware("admin"))
}
