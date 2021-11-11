package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"pusher/app/models"
	"pusher/pkg/config"
	"pusher/pkg/utils"
	"pusher/platform/client"
)

func GetSend(c *fiber.Ctx) error {
	msg := &models.Message{}

	if config.Key != "" && c.Get("Secure-Key", "") != config.Key {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	msg.Text, _ = url.QueryUnescape(c.Query("text", ""))
	msg.DisableLinkPreview = c.Query("preview", "0") == "0"
	msg.Msg, _ = url.QueryUnescape(c.Query("msg", ""))

	if msg.Text == "" {
		if msg.Msg == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": 1,
				"msg":  "empty payload",
			})
		}
		msg.Text = msg.Msg
	}

	msg.ChatId = config.ChatId
	msg.ParseMode = config.ParseMode
	msg.Text = utils.EscapeMarkdown(msg.Text)

	code, _, err := client.HttpClient.Post(config.ApiUrl).JSON(msg).Bytes()
	if err != nil || code != fiber.StatusOK {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": 0,
		"msg":  "OK",
	})
}

func PostSend(c *fiber.Ctx) error {
	msg := &models.Message{}

	if config.Key != "" && c.Get("Secure-Key", "") != config.Key {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err := c.BodyParser(msg); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if msg.Text == "" {
		if msg.Msg == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": 1,
				"msg":  "empty payload",
			})
		}
		msg.Text = msg.Msg
	}

	msg.ChatId = config.ChatId
	msg.ParseMode = config.ParseMode
	msg.Text = utils.EscapeMarkdown(msg.Text)

	code, _, err := client.HttpClient.Post(config.ApiUrl).JSON(msg).Bytes()
	if err != nil || code != fiber.StatusOK {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": 0,
		"msg":  "OK",
	})
}
