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
 
// AdminService adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Admin.
type AdminService interface {
	Create(service echo.Context, request modelsrequest.AdminCreateRequest) (*domain.Admin, error)
	Login(service echo.Context, request modelsrequest.AdminLoginRequest) (*domain.Admin, error)
	GetId(service echo.Context, id int) (*domain.Admin, error)
	GetAll(service echo.Context) ([]domain.Admin, error)
	Update(service echo.Context, request modelsrequest.AdminUpdateRequest, id int) (*domain.Admin, error)
	Delete(service echo.Context, id int) error
}

// AdminServiceImpl adalah implementasi dari AdminService.
type ServiceAdmin struct {
	RepositoryAdmin repository.AdminRepository
	Validate        *validator.Validate
}

// Updated code
func NewAdminService(AdminRepository repository.AdminRepository, validate *validator.Validate) AdminService {
    return &ServiceAdmin{
        RepositoryAdmin: AdminRepository,
        Validate:        validate,
    }
}

// NewAdminService adalah konstruktor untuk membuat instance AdminServiceImpl.
// func NewAdminService(AdminService repository.AdminRepository, validate *validator.Validate) AdminService {
// 	return &ServiceAdmin{
// 		RepositoryAdmin: AdminService,
// 		Validate:        validate,
// 	}
//}

// Create digunakan untuk menambahkan admin baru.
func (Service *ServiceAdmin) Create(service echo.Context, request modelsrequest.AdminCreateRequest) (*domain.Admin, error) {
	 // Validasi input dengan menggunakan validator.
	err := Service.Validate.Struct(request)
	if err != nil {
		// Jika terdapat kesalahan validasi, kembalikan pesan kesalahan validasi.
		return nil, helper.ValidationError(service, err)
	}
	// Periksa apakah admin dengan email yang sama sudah ada.
	existingAdmin, _ := Service.RepositoryAdmin.GetByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("admin already exists")
	}
	// Ubah data request menjadi domain.Admin dan hash password sebelum menyimpannya.
	admin := req.AdminCreateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	// Simpan admin baru ke dalam repository.
	result, err := Service.RepositoryAdmin.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating admin: %s", err.Error())
	}

	return result, nil
}

// Login digunakan untuk menglogin admin.
func (Service *ServiceAdmin) Login(service echo.Context, request modelsrequest.AdminLoginRequest) (*domain.Admin, error) {
	// Validasi input dengan menggunakan validator.
	err := Service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}

	 // Cari admin berdasarkan email yang diinputkan.
	existingAdmin, err := Service.RepositoryAdmin.GetByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid Id, Email, or Password")
	}

	// Ubah data request menjadi domain.Admin dan periksa kesesuaian password.
	admin := req.AdminLoginRequestToAdminDomain(request)
	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid Id, Email, or Password")
	}
	return existingAdmin, nil
}

// Update digunakan untuk memperbarui informasi admin.
func (Service *ServiceAdmin) Update(service echo.Context, request modelsrequest.AdminUpdateRequest, id int) (*domain.Admin, error) {
	// Validasi input dengan menggunakan validator.
	err := Service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(service, err)
	}

	// Cari admin berdasarkan ID yang diinputkan.
	existingAdmin, _ := Service.RepositoryAdmin.GetId(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	// Ubah data request menjadi domain.Admin dan periksa kesesuaian password.
	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)
	result, err := Service.RepositoryAdmin.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}
	return result, nil
}

// Delete digunakan untuk menghapus admin berdasarkan ID.
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

// GetByID digunakan untuk mendapatkan informasi admin berdasarkan ID.
func (Service *ServiceAdmin) GetId(service echo.Context, id int) (*domain.Admin, error) {
	existingAdmin, _ := Service.RepositoryAdmin.GetId(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}
	return existingAdmin, nil
}

// GetAll digunakan untuk mendapatkan daftar semua admin.
func (Service *ServiceAdmin) GetAll(service echo.Context) ([]domain.Admin, error) {
	admin, err := Service.RepositoryAdmin.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error when getting all admins")
	}
	return admin, nil
}
