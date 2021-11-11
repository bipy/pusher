package routes

import (
	"github.com/gofiber/fiber/v2"
	"pusher/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")

	route.Get("/", controllers.GetSend)

	route.Post("/", controllers.PostSend)

	route.Get("/pulse", controllers.Pulse)

}
