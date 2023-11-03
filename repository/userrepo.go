package repository

import (
	"fmt"
	"furniture/models/domain"
	"furniture/models/schema"
	"furniture/utils/request"
	"furniture/utils/respons"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(user *domain.User, Id int) (*domain.User, error)
	Delete(id int) error
	GetId(id int) (*domain.User, error)
	GetAll() ([]domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (Repository *UserRepositoryImpl) CreateUser(user *domain.User) (*domain.User, error) {
	// Mengonversi dari domain.User ke schema.User untuk penyimpanan di database.
	userDb := request.UserDomainToUserSchema(*user)
	result := Repository.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}
	// Mengonversi dari schema.User kembali ke domain.User untuk respons.
	results := respons.UserSchemaToUserDomain(userDb)
	fmt.Println(result)
	return results, nil

}

func (Repository *UserRepositoryImpl) UpdateUser(user *domain.User, id int) (*domain.User, error) {
	// Menggunakan operasi Update pada database untuk memperbarui data pengguna berdasarkan ID.
	result := Repository.DB.Table("users").Where("id = ?", id).Updates(domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone})
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (Repository *UserRepositoryImpl) Delete(id int) error {
	// Menggunakan operasi Delete pada database untuk menghapus pengguna berdasarkan ID.
	result := Repository.DB.Delete(&schema.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (Repository *UserRepositoryImpl) GetId(id int) (*domain.User, error) {
	// Menggunakan operasi First pada database untuk mendapatkan pengguna berdasarkan ID.
	var user domain.User
	result := Repository.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (Repository *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	// Menggunakan operasi Find pada database untuk mendapatkan semua pengguna.
	var users []domain.User
	result := Repository.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (Repository *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	// Menggunakan operasi First pada database untuk mendapatkan pengguna berdasarkan alamat email.
	user := domain.User{}
	result := Repository.DB.Where("email =?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}