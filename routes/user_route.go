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

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	// Membuat repository untuk entitas User.
	userRepository := repository.NewUserRepository(db)

	// Membuat layanan (service) untuk entitas User dan menginisialisasi dengan repository dan validator.
	userService := service.NewUserService(userRepository, validate)

	// Membuat handler untuk entitas User dan menginisialisasi dengan layanan.
	userHandler := hendler.NewHandlerUser(userService)

	// Membuat grup rute untuk entitas User di Echo.
	usersGroup := e.Group("user")

	// Rute POST untuk membuat pengguna baru.
	usersGroup.POST("", userHandler.UserCreate)

	// Rute POST untuk proses login pengguna.
	usersGroup.POST("/login", userHandler.UserLogin)

	// Middleware: Menggunakan JWT untuk otentikasi.
	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	
	usersGroup.GET("", userHandler.UserGetAll)
	usersGroup.GET("/:id", userHandler.UserGetById)
	usersGroup.PUT("/:id", userHandler.UserUpdate)
	usersGroup.DELETE("/:id", userHandler.UserDelete)
}
