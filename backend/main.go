package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/conduit-golang/backend/router"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
