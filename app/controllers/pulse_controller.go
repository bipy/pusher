package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Pulse(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
