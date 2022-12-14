package main

import (
	"fmt"
	"park/goproject/first/controller"
	"park/goproject/first/database"
	"park/goproject/first/middleware"
	"park/goproject/first/models"
	"park/goproject/first/repository"
	routes "park/goproject/first/routes"
	"park/goproject/first/service"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("test")
	app := fiber.New()

	err := database.Connect()
	database.RedisConnect()
	if err != nil {
		panic(err)
	}

	database.Database.AutoMigrate(&models.User{}, &models.CreditCard{})

	test := database.NewRedisTest(database.Rds)

	//ret, _ := json.Marshal(user_dto)
	//app.Use(middleware.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRepo := repository.NewUserRepository(database.Database)
	userService := service.NewUserService(userRepo, test)
	userController := controller.NewUserController(userService)

	cardRepo := repository.NewCardRepository(database.Database)
	cardService := service.NewCardService(cardRepo, userRepo)
	cardController := controller.NewCardController(cardService)

	userCachMid := middleware.NewUserCacheMiddleware(test)
	r := routes.NewRouter(userCachMid)
	r.SetUpUserRoutes(app, userController)
	routes.SetRedis(app, test)
	routes.SetUpCardRoutes(app, cardController)
	routes.SetUpTokenRoutes(app)
	app.Listen(":3000")
}
