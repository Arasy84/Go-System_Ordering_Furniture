package routes

import (
	"furniture/hendler"
	"furniture/repository"
	"furniture/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	AdminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(AdminRepository, validate)
	adminController := hendler.NewHandlerAdmin(adminService)

	adminsGroup := e.Group("/admin")

	adminsGroup.POST("", adminController.AdminRegister)
	adminsGroup.POST("/login", adminController.AdminLogin)

	adminsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	adminsGroup.GET("/:id", adminController.AdminGetById)
	adminsGroup.GET("", adminController.AdminGetAll)
	adminsGroup.PUT("/:id", adminController.AdminUpdate)
	adminsGroup.DELETE("/:id", adminController.AdminDelete)
}
