package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

func UserCreateRequestToUserDomain(req modelsrequest.UserCreateRequest) *domain.User {
	return &domain.User{
        Name: req.Name,
        Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
    }
}

func UserLoginRequestToUserDomain(req modelsrequest.UserLogin) *domain.User {
	return &domain.User{
        Email: req.Email,
        Password: req.Password,
    }
}

func UserUpdateRequestToUserDomain(req modelsrequest.UserUpdate) *domain.User {
	return &domain.User{
		Name: req.Name,
		Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
	}
}

func UserDomainToUserSchema(req domain.User) *schema.User {
	return &schema.User{
        Name: req.Name,
        Email: req.Email,
        Password: req.Password,
        Address: req.Address,
        Phone: req.Phone,
    }
}