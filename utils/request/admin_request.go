package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

// AdminCreateRequestToAdminDomain mengonversi objek modelsrequest.AdminCreateRequest menjadi domain.Admin.
func AdminCreateRequestToAdminDomain(request modelsrequest.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

// AdminLoginRequestToAdminDomain mengonversi objek modelsrequest.AdminLoginRequest menjadi domain.Admin.
func AdminLoginRequestToAdminDomain(request modelsrequest.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Email:    request.Email,
		Password: request.Password,
	}
}

// AdminDomaintoAdminSchema mengonversi objek domain.Admin menjadi schema.Admin.
func AdminDomaintoAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

// AdminUpdateRequestToAdminDomain mengonversi objek modelsrequest.AdminUpdateRequest menjadi domain.Admin.
func AdminUpdateRequestToAdminDomain(request modelsrequest.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}