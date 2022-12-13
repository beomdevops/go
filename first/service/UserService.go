package service

import (
	"park/goproject/first/models"
	"park/goproject/first/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(user_repo repository.UserRepository) *UserService {

	return &UserService{userRepo: user_repo}
}

func (userService *UserService) CreateUser(name string) *models.UserDto {

	user := &models.User{Name: name}

	user, err := userService.userRepo.CreateUser(user)

	if err != nil {
		return nil
	}
	return user.ToDto()
}

func (userService *UserService) FindById(id int) *models.UserDto {
	user, err := userService.userRepo.FindById(id)

	if err != nil {
		return nil
	}

	return user.ToDto()
}

func (userService *UserService) FindByName(name string) *models.UserDto {
	user, err := userService.userRepo.FindByName(name)

	if err != nil {
		return nil
	}

	return user.ToDto()
}
