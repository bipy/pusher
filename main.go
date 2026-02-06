package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"pusher/pkg/routes"
)

func main() {
	app := echo.New()

	// Middleware
	app.Use(middleware.RequestLogger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		UnsafeAllowOriginFunc: func(c *echo.Context, origin string) (allowedOrigin string, allowed bool, err error) {
			return origin, true, nil
		}}))

	// Routes
	routes.PublicRoutes(app)

	// Server configuration with defaults
	host := getEnv("SERVER_HOST", "0.0.0.0")
	port := getEnv("SERVER_PORT", "3333")
	addr := fmt.Sprintf("%s:%s", host, port)

	// Create context that listens for the interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Start server with graceful shutdown
	sc := echo.StartConfig{
		Address:         addr,
		HideBanner:      true,
		GracefulTimeout: 10, // Graceful shutdown timeout in seconds
	}

	slog.Info("Starting server", "address", addr)
	if err := sc.Start(ctx, app); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
	slog.Info("Server shutdown completed")
}

// getEnv returns environment variable value or default if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
