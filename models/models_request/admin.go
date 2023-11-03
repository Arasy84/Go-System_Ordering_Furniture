package modelsrequest

// AdminCreateRequest adalah model permintaan yang digunakan untuk membuat admin baru.
type AdminCreateRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
}

// AdminLoginRequest adalah model permintaan yang digunakan untuk login admin.
type AdminLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,max=255"`
}

// AdminUpdateRequest adalah model permintaan yang digunakan untuk memperbarui informasi admin.
type AdminUpdateRequest struct {
	Name     string `json:"name" form:"name" validate:"min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"min=8,max=255"`
}