package repository

import (
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	"furniture/utils/respons"

	"gorm.io/gorm"
)

// AdminRepository adalah interface yang mendefinisikan operasi yang dapat dilakukan pada entitas Admin.
type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, id int) (*domain.Admin, error)
	GetId(id int) (*domain.Admin, error)
	GetAll() ([]domain.Admin, error)
	GetByEmail(email string) (*domain.Admin, error)
	Delete(id int) error
}

// AdminRepositoryImpl adalah implementasi dari AdminRepository.
type AdminRepositoryImpl struct {
	DB *gorm.DB
}

// NewAdminRepository adalah konstruktor untuk membuat instance AdminRepositoryImpl.
func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: db}
}

// Create digunakan untuk membuat admin baru di database.
func (Repository *AdminRepositoryImpl) Create(Admin *domain.Admin) (*domain.Admin, error) {
	adminDb := request.AdminDomaintoAdminSchema(*Admin)
	result := Repository.DB.Create(&adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := respons.AdminSchemaToAdminDomain(adminDb)

	return results, nil
}

// Update digunakan untuk memperbarui data admin berdasarkan ID.
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

// Delete digunakan untuk menghapus admin berdasarkan ID.
func (Repository *AdminRepositoryImpl) Delete(id int) error {
	result := Repository.DB.Delete(&schema.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetId digunakan untuk mendapatkan admin berdasarkan ID.
func (Repository *AdminRepositoryImpl) GetId(id int) (*domain.Admin, error) {
	var admin domain.Admin
	result := Repository.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

// GetAll digunakan untuk mendapatkan semua admin yang tersedia.
func (Repository *AdminRepositoryImpl) GetAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := Repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

// GetByEmail digunakan untuk mendapatkan admin berdasarkan alamat email.
func (repository *AdminRepositoryImpl) GetByEmail(email string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
