package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

func AdminSchemaToAdminDomain(Admin *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password}
}

func AdminDomainToAdminSchema(Admin *domain.Admin) *schema.Admin {
	return &schema.Admin{
		ID:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password}
}

// func AdminDomainToAdminRespons(Admin *domain.Admin) modelsrespons.AdminReponse {
// 	return modelsrespons.AdminReponse{
// 		Id:       Admin.ID,
// 		Name:     Admin.Name,
// 		Email:    Admin.Email,
// 		Password: Admin.Password}
// }

func AdminDomainToAdminLoginResponse(Admin *domain.Admin) modelsrespons.AdminLogin {
	return modelsrespons.AdminLogin{
		Name:  Admin.Name,
		Email: Admin.Email,
	}
}

func AdminDomaintoAdminResponse(Admin *domain.Admin) modelsrespons.AdminReponse {
	return modelsrespons.AdminReponse{
		Id:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password,
	}
}

func ConvertAdminResponse(Admin []domain.Admin) []modelsrespons.AdminReponse {
	var results []modelsrespons.AdminReponse
	for _, Admin := range Admin {
		userResponse := modelsrespons.AdminReponse{
			Id:       Admin.ID,
			Name:     Admin.Name,
			Email:    Admin.Email,
			Password: Admin.Password,
		}
		results = append(results, userResponse)
	}
	return results
}
