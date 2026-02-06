package controllers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"pusher/pkg/config"
)

// Authorize validates the request using either the Secure-Key header or key query parameter
// This allows authentication without custom headers for easier integration
func Authorize(c echo.Context) error {
	// Try to get key from header first
	key := c.Request().Header.Get("Secure-Key")

	// If not in header, try query parameter
	if key == "" {
		key = c.QueryParam("key")
	}

	// Validate the key
	if key == "" || key != config.Key {
		return errors.New("unauthorized")
	}
	return nil
}
