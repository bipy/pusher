package controllers

import (
    "github.com/gofiber/fiber/v2"
    "net/url"
    "pusher/app/models"
    "pusher/pkg/config"
    "pusher/platform/client"
)

func Send(c *fiber.Ctx) error {
    msg := &models.Message{}

    if config.Key != "" && c.Get("Secure-Key", "") != config.Key {
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    if c.Method() == fiber.MethodGet {
        msg.Text, _ = url.QueryUnescape(c.Query("text", ""))
        msg.DisableLinkPreview = c.Query("preview", "0") == "0"
    } else if c.Method() == fiber.MethodPost {
        if err := c.BodyParser(msg); err != nil {
            return c.SendStatus(fiber.StatusBadRequest)
        }
    } else {
        return c.SendStatus(fiber.StatusMethodNotAllowed)
    }

    if msg.Text == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code": 1,
            "msg": "empty payload",
        })
    }

    msg.ChatId = config.ChatId
    msg.ParseMode = config.ParseMode

    code, _, err := client.HttpClient.Post(config.ApiUrl).JSON(msg).Bytes()
    if err != nil || code != fiber.StatusOK {
        return c.SendStatus(fiber.StatusBadGateway)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code": 0,
        "msg": "OK",
    })
}
