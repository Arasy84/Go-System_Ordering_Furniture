package main

import (
	"furniture/config"
	"furniture/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Membuat instance objek Echo
	app := echo.New()
	// Membuat instance objek validator
	validate := validator.New()
	// Terhubung ke database
	DB := config.ConnectDB()

	// app.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Welcome To Alta furniture")
	// })

	// Menghubungkan rute-rute dengan aplikasi Echo
	routes.OpenAI(app)
	routes.AdminRoutes(app, DB, validate)
	routes.ProductRoute(app, DB, validate)
	routes.UserRoutes(app, DB, validate)
	routes.OrderRoutes(app, DB, validate)

	// Middleware: Menghapus trailing slash dari URL
	app.Pre(middleware.RemoveTrailingSlash())
	// Middleware: Mengaktifkan CORS (Cross-Origin Resource Sharing)
	app.Use(middleware.CORS())
	// Middleware: Logger untuk mencatat permintaan HTTP
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	// Menjalankan server HTTP di port 8000
	app.Logger.Fatal(app.Start(":8000"))
}
