package service

import (
	"fmt"
	"furniture/helper"
	"furniture/models/domain"
	modelsrequest "furniture/models/models_request"
	"furniture/repository"
	req "furniture/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminService interface {
	Create(service echo.Context, request modelsrequest.AdminCreateRequest) (*domain.Admin, error)
	Login(service echo.Context, request modelsrequest.AdminLoginRequest) (*domain.Admin, error)
	GetId(service echo.Context, id int) (*domain.Admin, error)
	GetAll(service echo.Context) ([]domain.Admin, error)
	Update(service echo.Context, request modelsrequest.AdminUpdateRequest, id int) (*domain.Admin, error)
	Delete(service echo.Context, id int) error
}

type ServiceAdmin struct {
	RepositoryAdmin repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminService(AdminService repository.AdminRepository, validate *validator.Validate) *ServiceAdmin {
	return &ServiceAdmin{
		RepositoryAdmin: AdminService,
		Validate:        validate,
	}
}

func (Service *ServiceAdmin) Create(service echo.Context, request modelsrequest.AdminCreateRequest) (*domain.Admin, error) {
	err := Service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}
	existingAdmin, _ := Service.RepositoryAdmin.GetByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("admin already exists")
	}
	admin := req.AdminCreateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	result, err := Service.RepositoryAdmin.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating admin: %s", err.Error())
	}

	return result, nil
}

func (Service *ServiceAdmin) Login(service echo.Context, request modelsrequest.AdminLoginRequest) (*domain.Admin, error) {
	err := Service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}
	existingAdmin, err := Service.RepositoryAdmin.GetByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid Id, Email, or Password")
	}
	admin := req.AdminLoginRequestToAdminDomain(request)
	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid Id, Email, or Password")
	}
	return existingAdmin, nil
}

func (Service *ServiceAdmin) Update(service echo.Context, request modelsrequest.AdminUpdateRequest, id int) (*domain.Admin, error) {
	err := Service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}

	existingAdmin, _ := Service.RepositoryAdmin.GetId(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)
	result, err := Service.RepositoryAdmin.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}
	return result, nil
}

func (Service *ServiceAdmin) Delete(service echo.Context, id int) error {
	existingAdmin, _ := Service.RepositoryAdmin.GetId(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}
	err := Service.RepositoryAdmin.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting admin: %s", err.Error())
	}
	return nil
}

func (Service *ServiceAdmin) GetId(service echo.Context, id int) (*domain.Admin, error) {
	existingAdmin, _ := Service.RepositoryAdmin.GetId(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}
	return existingAdmin, nil
}

func (Service *ServiceAdmin) GetAll(service echo.Context) ([]domain.Admin, error) {
	admin, err := Service.RepositoryAdmin.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error when getting all admins")
	}
	return admin, nil
}
