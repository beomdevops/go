package tokenHandler

import (
	token "park/goproject/first/jwt"

	fiber "github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {

	jwt, err := token.GenJwt()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": jwt})
}

func GetJwk(c *fiber.Ctx) error {

	jwk, err := token.GenJwk()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	return c.Status(200).SendString(jwk)
}
