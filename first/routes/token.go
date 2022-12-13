package routes

import (
	"park/goproject/first/tokenHandler"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpTokenRoutes(app *fiber.App) {

	token := app.Group("/token", logger.New())
	token.Get("/jwt", tokenHandler.GetJwtToken)
	token.Get("/jwk", tokenHandler.GetJwk)
}
