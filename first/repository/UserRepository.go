package repository

import (
	"errors"
	"park/goproject/first/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(u *models.User) (*models.User, error)
	FindById(id int) (*models.User, error)
	FindByName(name string) (*models.User, error)
}

func NewUserRepository(cdb *gorm.DB) UserRepository {
	return &userRepository{db: cdb}
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) CreateUser(u *models.User) (*models.User, error) {

	result := repo.db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil

}

func (repo *userRepository) FindById(id int) (*models.User, error) {

	data := &models.User{}

	result := repo.db.Find(data, "id = ?", id)

	if result.RowsAffected < 1 {
		return nil, errors.New("not found")
	}
	return data, nil
}
func (repo *userRepository) FindByName(name string) (*models.User, error) {
	data := &models.User{}
	result := repo.db.Where("name = ?", name).First(data)
	if result.RowsAffected < 1 {
		return nil, errors.New("not found")
	}
	return data, nil

}
