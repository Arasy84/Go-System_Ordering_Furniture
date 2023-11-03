package helper

import (
	// "errors"
	// "fmt"
	modelrespons "furniture/models/models_response"
	"os"
	// "strconv"
	// "strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAdminToken digunakan untuk menghasilkan token JWT untuk admin.
func GenerateAdminToken(adminLoginResponse *modelrespons.AdminLogin, id uint) (string, error) {
	// Menghitung waktu kedaluwarsa token (expire time) sekitar 7 hari dari saat ini
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	 // Membuat klaim (claims) JWT yang berisi ID, email, waktu kedaluwarsa, dan peran "admin"
	claims := jwt.MapClaims{
		"id":    id,
		"email": adminLoginResponse.Email,
		"exp":   expireTime,
		"role":  "admin",
	}

	// Membuat token JWT dengan metode penandatanganan HS256 (HMAC-SHA256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Menandatangani token menggunakan kunci rahasia dari environment variable
	validToken, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

// GenerateUserToken digunakan untuk menghasilkan token JWT untuk pengguna (user).
func GenerateUserToken(UserLoginResponse *modelrespons.UserLogin, id uint) (string, error) {
	 // Menghitung waktu kedaluwarsa token (expire time) sekitar 7 hari dari saat ini
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	// Membuat klaim (claims) JWT yang berisi ID, email, waktu kedaluwarsa, dan peran "user"
	claims := jwt.MapClaims{
		"id":    id,
		"email": UserLoginResponse.Email,
		"exp":   expireTime,
		"role":  "user",
	}

	// Membuat token JWT dengan metode penandatanganan HS256 (HMAC-SHA256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Menandatangani token menggunakan kunci rahasia dari environment variable
	validToken, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

// func ExtractToken(tokenString string) (int, error) {
// 	parts := strings.Split(tokenString, " ")
// 	if len(parts) != 2 || parts[0] != "Bearer" {
// 		return 0, errors.New("invalid token")
// 	}
// 	jwtToken := parts[1]

// 	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("SECRET")), nil
// 	})
// 	if err != nil {
// 		return 0, errors.New("invalid token")
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		noTableString := fmt.Sprint(claims["no_table"])
// 		noTable, _ := strconv.Atoi(noTableString)
// 		return noTable, nil
// 	}

// 	return 0, errors.New("claims not found")
// }
