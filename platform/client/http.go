package client

import "github.com/gofiber/fiber/v2"

var HttpClient *fiber.Client

func init() {
    HttpClient = &fiber.Client{}
}
