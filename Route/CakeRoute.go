package Route

import (
	"github.com/gofiber/fiber/v2"
	"ralali-crud-cake-test/Config"
)

func InitCake(c fiber.Router) {
	DI := CakeDI(Config.DB)
	c.Get("/cakes", DI.GetCake)
	c.Get("/cakes/:id", DI.GetCakeById)
	c.Post("/cakes", DI.AddCake)
	c.Put("/cakes/:id", DI.UpdateCake)
	c.Delete("/cakes/:id", DI.DeleteCake)
}
