package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

// AdminSchemaToAdminDomain mengonversi objek schema.Admin menjadi domain.Admin.
func AdminSchemaToAdminDomain(Admin *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password}
}

// AdminDomainToAdminSchema mengonversi objek domain.Admin menjadi schema.Admin.
func AdminDomainToAdminSchema(Admin *domain.Admin) *schema.Admin {
	return &schema.Admin{
		ID:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password}
}

// AdminDomainToAdminLoginResponse mengonversi objek domain.Admin menjadi modelsrespons.AdminLogin.
func AdminDomainToAdminLoginResponse(Admin *domain.Admin) modelsrespons.AdminLogin {
	return modelsrespons.AdminLogin{
		Name:  Admin.Name,
		Email: Admin.Email,
	}
}

// AdminDomainToAdminResponse mengonversi objek domain.Admin menjadi modelsrespons.AdminReponse.
func AdminDomaintoAdminResponse(Admin *domain.Admin) modelsrespons.AdminReponse {
	return modelsrespons.AdminReponse{
		Id:       Admin.ID,
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password,
	}
}

// ConvertAdminResponse mengonversi slice dari objek domain.Admin menjadi slice dari modelsrespons.AdminReponse.
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
