package controllers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"pusher/pkg/config"
)

// Authorize validates the request using the Secure-Key header
func Authorize(c echo.Context) error {
	key := c.Request().Header.Get("Secure-Key")
	if key == "" || key != config.Key {
		return errors.New("unauthorized")
	}
	return nil
}
