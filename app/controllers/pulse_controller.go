package controllers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// Pulse is a health check endpoint
func Pulse(c *echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
