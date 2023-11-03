package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

// UserDomainToUserLoginResponse mengonversi objek domain.User menjadi modelsrespons.UserLogin.
func UserDomainToUserLoginResponse(user *domain.User) modelsrespons.UserLogin {
	return modelsrespons.UserLogin{
		ID:    user.ID,
		Email: user.Email,
		Password: user.Password,
	}
}

// UserDomainToUserResponse mengonversi objek domain.User menjadi modelsrespons.UserResponse.
func UserDomainToUserResponse(user *domain.User) modelsrespons.UserResponse {
	return modelsrespons.UserResponse{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
}

// UserSchemaToUserDomain mengonversi objek schema.User menjadi domain.User.
func UserSchemaToUserDomain(user *schema.User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
}

// ConvertUserResponse mengonversi slice dari objek domain.User menjadi slice dari modelsrespons.UserResponse.
func ConvertUserResponse(users []domain.User) []modelsrespons.UserResponse {
	var results []modelsrespons.UserResponse
	for _, user := range users {
		userResponse := modelsrespons.UserResponse{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Address:  user.Address,
			Phone:    user.Phone,	
		}
		results = append(results, userResponse)
	}
	return results
}
