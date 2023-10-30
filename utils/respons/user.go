package respons

import (
	"furniture/models/domain"
	modelsrespons "furniture/models/models_response"
	"furniture/models/schema"
)

func UserDomainToUserLoginResponse(user *domain.User) modelsrespons.UserLogin {
	return modelsrespons.UserLogin{
		Name:  user.Name,
		Email: user.Email,
	}
}

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

func ConvertUserResponse(users []domain.User) []modelsrespons.UserResponse {
	var results []modelsrespons.UserResponse
	for _, user := range users {
		userResponse := modelsrespons.UserResponse{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Address:  user.Address,
		}
		results = append(results, userResponse)
	}
	return results
}
