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
	userDb := request.UserDomainToUserSchema(*user)
	result := Repository.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}
	results := respons.UserSchemaToUserDomain(userDb)
	fmt.Println(result)
	return results, nil

}

func (Repository *UserRepositoryImpl) UpdateUser(user *domain.User, id int) (*domain.User, error) {
	result := Repository.DB.Table("users").Where("id = ?", id).Updates(domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,})
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (Repository *UserRepositoryImpl) Delete(id int) error {
	result := Repository.DB.Delete(&schema.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (Repository *UserRepositoryImpl) GetId(id int) (*domain.User, error) {
	var user domain.User
	result := Repository.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (Repository *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	var users []domain.User
	result := Repository.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (Repository *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	result := Repository.DB.Where("email =?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
