package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Pulse is a health check endpoint
func Pulse(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
