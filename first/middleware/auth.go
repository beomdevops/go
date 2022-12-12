package middleware

import (
	"fmt"
	"park/goproject/first/jwt"

	fiber "github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Auth")
		if header == "" {
			token := jwt.GenJwt()
			return c.Status(400).SendString(token)
		}
		fmt.Println(header)
		fmt.Println("this is middleware")
		return c.Next()
	}
}
