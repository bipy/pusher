package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"pusher/pkg/config"
)

func Authorize(c echo.Context) error {
	if c.Request().Header.Get("Secure-Key") != config.Key {
		return errors.New("unauthorized")
	}
	return nil
}
