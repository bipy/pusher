package controllers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v5"
	"pusher/app/models"
	"pusher/pkg/config"
	"pusher/pkg/utils"
	"pusher/platform"
)

// GetSend handles GET requests to send messages to Telegram
func GetSend(c *echo.Context) error {
	// Check authentication if enabled
	if config.Key != "" {
		if err := Authorize(c); err != nil {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
		}
	}

	params := c.QueryParams()

	// Parse message text
	text, err := url.QueryUnescape(params.Get("text"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid text parameter", nil))
	}

	// Fallback to 'msg' parameter if 'text' is empty
	if text == "" {
		msg, err := url.QueryUnescape(params.Get("msg"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid msg parameter", nil))
		}
		if msg == "" {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("empty payload: text or msg required", nil))
		}
		text = msg
	}

	// Parse options
	disableLinkPreview := !params.Has("preview")
	useMarkdown := params.Has("markdown")

	// Escape markdown if not explicitly enabled
	if !useMarkdown {
		text = utils.EscapeMarkdown(text)
	}

	// Get and escape client IP
	ip := utils.EscapeMarkdown(c.RealIP())

	// Send to Telegram
	if err := platform.Push([]rune(text), disableLinkPreview, ip); err != nil {
		return c.JSON(http.StatusBadGateway, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}

// PostSend handles POST requests to send messages to Telegram
func PostSend(c *echo.Context) error {
	// Check authentication if enabled
	if config.Key != "" {
		if err := Authorize(c); err != nil {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
		}
	}

	// Parse request body
	msg := &models.ReqMessage{}
	if err := c.Bind(msg); err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid request body", nil))
	}

	// Fallback to 'msg' field if 'text' is empty
	if msg.Text == "" {
		if msg.Msg == "" {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("empty payload: text or msg required", nil))
		}
		msg.Text = msg.Msg
	}

	// Escape markdown if not explicitly enabled
	if !msg.Markdown {
		msg.Text = utils.EscapeMarkdown(msg.Text)
	}

	// Get and escape client IP
	ip := utils.EscapeMarkdown(c.RealIP())

	// Send to Telegram
	if err := platform.Push([]rune(msg.Text), !msg.Preview, ip); err != nil {
		return c.JSON(http.StatusBadGateway, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}
