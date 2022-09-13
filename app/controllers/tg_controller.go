package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"pusher/app/models"
	"pusher/pkg/config"
	"pusher/pkg/utils"
	"pusher/platform"
)

func GetSend(c echo.Context) error {
	if config.Key != "" {
		err := Authorize(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
		}
	}

	params := c.QueryParams()

	text, err := url.QueryUnescape(params.Get("text"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	if text == "" {
		msg, err := url.QueryUnescape(params.Get("msg"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		}
		if msg == "" {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("empty payload", nil))
		}
		text = msg
	}

	disableLinkPreview := !params.Has("preview")

	useMarkdown := params.Has("markdown")

	if !useMarkdown {
		text = utils.EscapeMarkdown(text)
	}

	ip := utils.EscapeMarkdown(c.RealIP())

	err = platform.Push([]rune(text), disableLinkPreview, ip)
	if err != nil {
		return c.JSON(http.StatusBadGateway, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}

func PostSend(c echo.Context) error {
	if config.Key != "" {
		err := Authorize(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
		}
	}

	msg := &models.ReqMessage{}
	err := c.Bind(msg)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	if msg.Text == "" {
		if msg.Msg == "" {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("empty payload", nil))
		}
		msg.Text = msg.Msg
	}

	if !msg.Markdown {
		msg.Text = utils.EscapeMarkdown(msg.Text)
	}

	ip := utils.EscapeMarkdown(c.RealIP())

	err = platform.Push([]rune(msg.Text), !msg.Preview, ip)
	if err != nil {
		return c.JSON(http.StatusBadGateway, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(nil))
}
