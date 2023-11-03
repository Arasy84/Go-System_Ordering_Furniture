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
	// Membuat repository untuk entitas Admin.
	AdminRepository := repository.NewAdminRepository(db)

	// Membuat layanan (service) untuk entitas Admin dan menginisialisasi dengan repository dan validator.
	adminService := service.NewAdminService(AdminRepository, validate)

	// Membuat handler untuk entitas Admin dan menginisialisasi dengan layanan.
	adminController := hendler.NewHandlerAdmin(adminService)

	// Membuat grup rute untuk entitas Admin di Echo.
	adminsGroup := e.Group("/admin")

	// Rute POST untuk mendaftarkan admin baru.
	adminsGroup.POST("", adminController.AdminRegister)

	// Rute POST untuk login admin.
	adminsGroup.POST("/login", adminController.AdminLogin)

	// Middleware: Menggunakan JWT untuk otentikasi.
	adminsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	adminsGroup.GET("/:id", adminController.AdminGetById)
	adminsGroup.GET("", adminController.AdminGetAll)
	adminsGroup.PUT("/:id", adminController.AdminUpdate)
	adminsGroup.DELETE("/:id", adminController.AdminDelete)
}
