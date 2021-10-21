package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "log"
    "os"
    "pusher/pkg/routes"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    routes.PublicRoutes(app)
    routes.NotFoundRoute(app)

    listen := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

    log.Fatal(app.Listen(listen))
}
