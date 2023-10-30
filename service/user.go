package service

import (
	"fmt"
	"furniture/helper"
	"furniture/models/domain"
	req "furniture/utils/request"
	modelsrequest "furniture/models/models_request"
	"furniture/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(service echo.Context, request modelsrequest.UserCreateRequest) (*domain.User, error)
	Login(sercvice echo.Context, request modelsrequest.UserLogin) (*domain.User, error)
	Update(service echo.Context, request modelsrequest.UserUpdate, id int) (*domain.User, error)
	Delete(service echo.Context, id int) error
	GetId(service echo.Context, id int) (*domain.User, error)
	GetAll(service echo.Context) ([]domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate *validator.Validate
}

func NewUserService(UserRepository repository.UserRepository, Validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		Validate: Validate,
	}
}


func (s *UserServiceImpl) CreateUser(service echo.Context, request modelsrequest.UserCreateRequest) (*domain.User, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}

	existingUser, _ := s.UserRepository.GetByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("email Already Exists")
	}

	user := req.UserCreateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := s.UserRepository.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("error when creating user: %s", err.Error())
	}

	return result, nil

}

func (s *UserServiceImpl) Login(service echo.Context, request modelsrequest.UserLogin) (*domain.User, error) {
	err := s.Validate.Struct(request)
    if err!= nil {
        return nil, helper.ValidationError(service, err)
    }

    existingUser, err := s.UserRepository.GetByEmail(request.Email)
    if err!= nil {
        return nil, fmt.Errorf("invalid Email, or Password")
    }
    user := req.UserLoginRequestToUserDomain(request)
    err = helper.ComparePassword(existingUser.Password, user.Password)
    if err!= nil {
        return nil, fmt.Errorf("invalid Id, Email, or Password")
    }
    return existingUser, nil
}

func (s *UserServiceImpl) Update(service echo.Context, request modelsrequest.UserUpdate, id int) (*domain.User, error) {
	err := s.Validate.Struct(request)
    if err!= nil {
        return nil, helper.ValidationError(service, err)
    }

    existingUser, _ := s.UserRepository.GetId(id)
    if existingUser == nil {
        return nil, fmt.Errorf("user not found")
    }

    user := req.UserUpdateRequestToUserDomain(request)
    user.Password = helper.HashPassword(user.Password)
    result, err := s.UserRepository.UpdateUser(user, id)
    if err!= nil {
        return nil, fmt.Errorf("error when updating user: %s", err.Error())
    }
    return result, nil
}

func (s *UserServiceImpl) Delete(service echo.Context, id int) error {
	existingUser, _ := s.UserRepository.GetId(id)
    if existingUser == nil {
        return fmt.Errorf("user not found")
    }
    err := s.UserRepository.Delete(id)
    if err!= nil {
        return fmt.Errorf("error when deleting user: %s", err.Error())
    }
    return nil
}

func (s *UserServiceImpl) GetId(service echo.Context, id int) (*domain.User, error) {

	existingUser, _ := s.UserRepository.GetId(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	return existingUser, nil
}

func (s *UserServiceImpl) GetAll(sercvice echo.Context) ([]domain.User, error) {
	user, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("users not found")
	}

	return user, nil
}