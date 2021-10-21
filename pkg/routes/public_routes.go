package routes

import (
    "github.com/gofiber/fiber/v2"
    "pusher/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
    // Create routes group.
    route := a.Group("/")

    route.Get("/", controllers.Send)

    route.Post("/", controllers.Send)

    route.Get("/pulse", controllers.Pulse)

}
