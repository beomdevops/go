package service

import (
	"park/goproject/first/database"
	"park/goproject/first/models"
	"park/goproject/first/repository"
)

type UserService struct {
	userRepo repository.UserRepository
	userRds  *database.RedisTest
}

func NewUserService(user_repo repository.UserRepository, di_userRds *database.RedisTest) *UserService {

	return &UserService{
		userRepo: user_repo,
		userRds:  di_userRds,
	}
}

func (userService *UserService) CreateUser(p_user *models.User) *models.UserDto {

	user, err := userService.userRepo.CreateUser(p_user)

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
	userService.userRds.SetUser(user)

	return user.ToDto()
}

func (userService *UserService) FindByName(name string) *models.UserDto {
	user, err := userService.userRepo.FindByName(name)

	if err != nil {
		return nil
	}

	return user.ToDto()
}
