package main

import (
	"fmt"
	"park/goproject/first/middleware"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("test")
	app := fiber.New()

	app.Use(middleware.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
