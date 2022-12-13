package main

import (
	"fmt"
	"park/goproject/first/database"
	"park/goproject/first/middleware"
	routes "park/goproject/first/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("test")
	app := fiber.New()

	err := database.Connect()

	if err != nil {
		defer database.cloes()
		panic(err)
	}
	//database.Database.AutoMigrate(&models.User{})

	//ret, _ := json.Marshal(user_dto)
	app.Use(middleware.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetUpTokenRoutes(app)
	app.Listen(":3000")
}
