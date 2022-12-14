package routes

import (
	"park/goproject/first/controller"
	"park/goproject/first/database"
	"park/goproject/first/middleware"
	"park/goproject/first/tokenHandler"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	userCacheMiddleware *middleware.UserCacheMiddleware
}

func NewRouter(di_userCacheMiddleware *middleware.UserCacheMiddleware) *Router {

	return &Router{
		userCacheMiddleware: di_userCacheMiddleware,
	}
}

func SetUpTokenRoutes(app *fiber.App) {

	token := app.Group("/token", logger.New())
	token.Get("/jwt", tokenHandler.GetJwtToken)
	token.Get("/jwk", tokenHandler.GetJwk)
}

func (r *Router) SetUpUserRoutes(app *fiber.App, userController *controller.UserController) {
	user := app.Group("/users", logger.New())
	user.Get("/:userId", r.userCacheMiddleware.GetUserId(), userController.FindById)
	user.Get("/:userName", userController.FindByName)
	user.Post("/", userController.CreateUser)
}

func SetUpCardRoutes(app *fiber.App, cardController *controller.CardController) {
	card := app.Group("/cards", logger.New())
	card.Get("/:cardId", cardController.FindByCardId)
	card.Get("/:userId", cardController.FindByUserId)
	card.Post("/", cardController.CreateCard)
}

func SetRedis(app *fiber.App, redisTest *database.RedisTest) {
	redisT := app.Group("/redis", logger.New())
	redisT.Get("/", redisTest.SetDate)
	redisT.Get("/get", redisTest.GetData)
}
