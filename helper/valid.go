package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// ValidationError adalah fungsi yang digunakan untuk menangani error validasi dan mengembalikan error yang sesuai.
// Fungsi ini memeriksa apakah error yang diterima adalah tipe error validasi, dan jika iya, itu akan mengembalikan pesan error yang sesuai.
func ValidationError(ctx echo.Context, err error) error {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		messages := make([]string, 0)
		for _, e := range validationError {
			messages = append(messages, fmt.Sprintf("Validation error on field %s, tag %s", e.Field(), e.Tag()))
		}

		return fmt.Errorf("validation failed: %s", strings.Join(messages, "; "))
	}

	return nil
}