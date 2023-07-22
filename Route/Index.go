package Route

import (
	"github.com/gofiber/fiber/v2"
)

func Init(c *fiber.App) {
	var (
		ApiPrefix = c.Group("/api/")
	)

	InitCake(ApiPrefix)
}
