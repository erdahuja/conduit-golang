package helpers

import "github.com/gofiber/fiber/v2"

// ParseBody derives body as per destination from context
func ParseBody(ctx *fiber.Ctx, dst interface{}) {
	err := ctx.BodyParser(dst)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}

func HashPassword() {
	
}