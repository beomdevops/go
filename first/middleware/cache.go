package middleware

import (
	"fmt"
	"park/goproject/first/database"

	fiber "github.com/gofiber/fiber/v2"
)

type UserCacheMiddleware struct {
	redisTest *database.RedisTest
}

func NewUserCacheMiddleware(di *database.RedisTest) *UserCacheMiddleware {
	return &UserCacheMiddleware{redisTest: di}
}

func (u *UserCacheMiddleware) GetUserId() fiber.Handler {
	return func(c *fiber.Ctx) error {

		parma := c.Params("userId")

		data, err := u.redisTest.GetUser(parma)

		if err != nil {
			fmt.Println(parma)
			return c.Next()
		}
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "success", "data": data})
	}
}
