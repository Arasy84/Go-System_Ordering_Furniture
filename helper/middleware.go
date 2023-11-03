package helper

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware adalah middleware untuk otorisasi berdasarkan peran (role) pengguna.
// Parameter roles adalah daftar peran yang diizinkan untuk mengakses endpoint tertentu.
func AuthMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Mendapatkan token dari konteks
			token := c.Get("user").(*jwt.Token)
			// Mendapatkan klaim (claims) dari token
			claims := token.Claims.(jwt.MapClaims)
			// Mendapatkan peran (role) dari klaim
			role := claims["role"].(string)

			
			_, ok := claims["role"].(string)
			if !ok {
				return c.JSON(http.StatusForbidden, "Role not found in token claims")
			}

			 // Memeriksa apakah peran pengguna ada di daftar peran yang diizinkan	
			for _, allowedRole := range roles {
				if role == allowedRole {
					return next(c)
				}
			}
			// Jika peran pengguna tidak ada dalam daftar peran yang diizinkan, mengembalikan respon Status Forbidden
			return c.JSON(http.StatusForbidden, "Unauthorized")
		}
	}
}
