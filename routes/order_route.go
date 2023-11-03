package routes

import (
	"furniture/helper"
	"furniture/hendler"
	"furniture/repository"
	"furniture/service"
	"os"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OrderRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository, validate)
	orderHandler := hendler.NewOrderHendler(orderService)

	ordersGroup := e.Group("")
	ordersGroup.GET("/order", orderHandler.FindAll)
	ordersGroup.GET("/order/:id", orderHandler.FindByID)
	// admin := e.Group("", helper.VerifySignIn("ADMIN_JWT_SECRET"))
	ordersGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	ordersGroup.POST("/user/order", orderHandler.CreateOrder, helper.AuthMiddleware("user"))
	ordersGroup.DELETE("/admin/order/:id", orderHandler.Delete, helper.AuthMiddleware("admin"))

}
