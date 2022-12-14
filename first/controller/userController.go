package controller

import (
	"park/goproject/first/models"
	"park/goproject/first/service"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(di_userService *service.UserService) *UserController {
	return &UserController{userService: di_userService}
}

func (userController *UserController) FindById(ctx *fiber.Ctx) error {
	parma := ctx.Params("userId")
	id, err := strconv.Atoi(parma)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	find_user := userController.userService.FindById(id)

	if find_user == nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_user})

}

func (userController *UserController) FindByName(ctx *fiber.Ctx) error {
	name := ctx.Params("userName")
	find_user := userController.userService.FindByName(name)

	if find_user == nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": find_user})
}

type UserCreateRequest struct {
	User_Name string `json:"name"`
}

func (user_dto *UserCreateRequest) toEntity() *models.User {
	return models.NewUser(user_dto.User_Name)
}

func (userController *UserController) CreateUser(ctx *fiber.Ctx) error {
	p := new(UserCreateRequest)
	err := ctx.BodyParser(p)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": err, "data": err})
	}

	user := userController.userService.CreateUser(p.toEntity())
	if user == nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "not create", "data": nil})
	}
	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": user})

}
