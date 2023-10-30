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
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userHandler := hendler.NewHandlerUser(userService)

	usersGroup := e.Group("user")

	usersGroup.POST("", userHandler.UserCreate)
	usersGroup.POST("/login", userHandler.UserLogin)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.GET("", userHandler.UserGetAll)
	usersGroup.GET("/:id", userHandler.UserGetById)
	usersGroup.PUT("/:id", userHandler.UserUpdate)
	usersGroup.DELETE("/:id", userHandler.UserDelete)
}
