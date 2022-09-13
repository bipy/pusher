package routes

import (
	"github.com/labstack/echo/v4"
	"pusher/app/controllers"
)

func PublicRoutes(a *echo.Echo) {
	a.GET("/", controllers.GetSend)
	a.POST("/", controllers.PostSend)

	a.GET("/pulse", controllers.Pulse)

}
