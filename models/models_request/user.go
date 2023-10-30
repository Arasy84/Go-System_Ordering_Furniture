package modelsrequest

type UserCreateRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
	Address string `json:"address" form:"address" validate:"required,min=1,max=255"`
	Phone uint `json:"phone" form:"phone"`
}

type UserLogin struct {
 	Email string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

type UserUpdate struct {
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email string `json:"email" form:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
	Address string `json:"address" form:"address" validate:"required,min=1,max=255"`
	Phone uint `json:"phone" form:"phone"`
}