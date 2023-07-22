package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ralali-crud-cake-test/Config"
	"ralali-crud-cake-test/Model"
	"ralali-crud-cake-test/Route"
)

func main() {
	/**
	TODO:
	1. Unit tests
	2. Validator
	3. Docker & Docker Compose
	*/
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Cake",
		AppName:       "Crud Cake Service",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(Model.ErrorResponse("request is invalid", err.Error()))
		},
	})

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	// cors
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,User-Agent,Content-Length",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
	}))

	// add security header
	app.Use(helmet.New())

	// compress
	app.Use(compress.New())

	// add etag
	app.Use(etag.New())

	// add recover
	app.Use(recover.New())

	// init database with model gen
	Config.InitDB()

	// init route
	Route.Init(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
