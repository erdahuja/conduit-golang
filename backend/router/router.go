package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/conduit-golang/backend/controllers"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api/v1", logger.New())
	api.Use(requestid.New())
	api.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4100",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	api.Use(recover.New())

	// Auth
	auth := api.Group("/users")
	auth.Post("/", controllers.Register)
	auth.Post("/login", controllers.Login)

	// User
	user := api.Group("/user")
	user.Get("/", controllers.GetUser)
	user.Put("/", controllers.UpdateUser)

	// fallback route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	// // Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
