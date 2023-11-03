package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

// UserCreateRequestToUserDomain mengonversi objek modelsrequest.UserCreateRequest menjadi domain.User.
func UserCreateRequestToUserDomain(req modelsrequest.UserCreateRequest) *domain.User {
	return &domain.User{
        Name: req.Name,
        Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
    }
}

// UserLoginRequestToUserDomain mengonversi objek modelsrequest.UserLogin menjadi domain.User.
func UserLoginRequestToUserDomain(req modelsrequest.UserLogin) *domain.User {
	return &domain.User{
        Email: req.Email,
        Password: req.Password,
    }
}

// UserUpdateRequestToUserDomain mengonversi objek modelsrequest.UserUpdate menjadi domain.User.
func UserUpdateRequestToUserDomain(req modelsrequest.UserUpdate) *domain.User {
	return &domain.User{
		Name: req.Name,
		Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
	}
}

// UserDomainToUserSchema mengonversi objek domain.User menjadi schema.User.
func UserDomainToUserSchema(req domain.User) *schema.User {
	return &schema.User{
        Name: req.Name,
        Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
    }
}