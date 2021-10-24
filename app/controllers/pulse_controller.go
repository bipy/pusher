package controllers

import "github.com/gofiber/fiber/v2"

func Pulse(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
