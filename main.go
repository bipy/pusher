package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pusher/pkg/routes"
)

func main() {
	app := echo.New()

	// Hide banner for cleaner output
	app.HideBanner = true

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	// Routes
	routes.PublicRoutes(app)

	// Server configuration with defaults
	host := getEnv("SERVER_HOST", "0.0.0.0")
	port := getEnv("SERVER_PORT", "3333")
	addr := fmt.Sprintf("%s:%s", host, port)

	// Start server
	go func() {
		app.Logger.Info("Starting server on ", addr)
		if err := app.Start(addr); err != nil {
			app.Logger.Info("Shutting down the server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}

// getEnv returns environment variable value or default if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
