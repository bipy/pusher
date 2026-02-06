package controllers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"pusher/pkg/config"
)

// Authorize validates the request using the Secure-Key header
func Authorize(c echo.Context) error {
	key := c.Request().Header.Get("Secure-Key")
	if key == "" {
		return errors.New("missing Secure-Key header")
	}
	if key != config.Key {
		return errors.New("invalid Secure-Key")
	}
	return nil
}
