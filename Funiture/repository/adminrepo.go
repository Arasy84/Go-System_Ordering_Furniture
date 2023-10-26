package repository

import (
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	"furniture/utils/respons"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, id int) (*domain.Admin, error)
	GetId(id int) (*domain.Admin, error)
	GetAll() ([]domain.Admin, error)
	GetByEmail(email string) (*domain.Admin, error)
	Delete(id int) error
}

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: db}
}

func (Repository *AdminRepositoryImpl) Create(Admin *domain.Admin) (*domain.Admin, error) {
	adminDb := request.AdminDomaintoAdminSchema(*Admin)
	result := Repository.DB.Create(&adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := respons.AdminSchemaToAdminDomain(adminDb)

	return results, nil
}

func (Repository *AdminRepositoryImpl) Update(Admin *domain.Admin, id int) (*domain.Admin, error) {
	result := Repository.DB.Table("admins").Where("id= ?", id).Updates(domain.Admin{
		Name:     Admin.Name,
		Email:    Admin.Email,
		Password: Admin.Password})
	if result.Error != nil {
		return nil, result.Error
	}
	return Admin, nil
}

func (Repository *AdminRepositoryImpl) Delete(id int) error {
	result := Repository.DB.Delete(&schema.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (Repository *AdminRepositoryImpl) GetId(id int) (*domain.Admin, error) {
	var admin domain.Admin
	result := Repository.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (Repository *AdminRepositoryImpl) GetAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := Repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) GetByEmail(email string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
