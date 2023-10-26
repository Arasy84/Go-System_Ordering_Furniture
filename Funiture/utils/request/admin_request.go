package request

import (
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/models/schema"
)

func AdminCreateRequestToAdminDomain(request modelsrequest.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request modelsrequest.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminDomaintoAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request modelsrequest.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}