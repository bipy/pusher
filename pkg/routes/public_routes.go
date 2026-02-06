package routes

import (
	"github.com/labstack/echo/v5"
	"pusher/app/controllers"
)

// PublicRoutes registers all public API routes
func PublicRoutes(a *echo.Echo) {
	// Message sending endpoints
	a.GET("/", controllers.GetSend)
	a.POST("/", controllers.PostSend)

	// Health check endpoint
	a.GET("/pulse", controllers.Pulse)
}
