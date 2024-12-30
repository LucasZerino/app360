package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/register", controllers.Register)
}
