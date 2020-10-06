package repository

import (
	db "github.com/antonioalfa22/go-grpc-template/config"
	"github.com/antonioalfa22/go-grpc-template/internal/models"
)

type UserRepository struct{}
var userRepository *UserRepository

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (r *UserRepository) Get(id uint64) (*models.User, error) {
	var user models.User
	where := models.User{ID: id}
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{}, "id asc")
	return &users, err
}

func (r *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{}, "id asc")
	return &users, err
}

func (r *UserRepository) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	return Save(&user)
}

func (r *UserRepository) Delete(user *models.User) error {
	return db.GetDB().Unscoped().Delete(&user).Error
}
