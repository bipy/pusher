package controllers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"pusher/pkg/config"
)

// Authorize validates the request using either the Secure-Key header or key query parameter.
//
// Two authentication methods are supported:
// 1. Secure-Key header (recommended for production - more secure)
// 2. key query parameter (easier for integrations but exposes credential in logs)
//
// Security Note: Query parameters are logged in URLs and may be exposed in server logs,
// browser history, and referrer headers. Use header-based authentication in production
// environments. Query parameter auth is provided for convenience in trusted environments
// or when custom headers are not feasible (e.g., simple webhooks, URL-based triggers).
func Authorize(c echo.Context) error {
	// Try to get key from header first (more secure)
	key := c.Request().Header.Get("Secure-Key")

	// If not in header, try query parameter (easier but less secure)
	if key == "" {
		key = c.QueryParam("key")
	}

	// Validate the key
	if key == "" || key != config.Key {
		return errors.New("unauthorized")
	}
	return nil
}
