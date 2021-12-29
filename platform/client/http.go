package client

import (
	"github.com/gofiber/fiber/v2"
	"pusher/app/models"
	"pusher/pkg/config"
	"pusher/pkg/utils"
)

var HttpClient *fiber.Client

const msgLength int = 1024

func init() {
	HttpClient = &fiber.Client{}
}

func DoPush(msg *models.Message) (int, []error) {
	for i := 0; i*msgLength < len(msg.Text); i++ {
		cur := &models.Message{
			ChatId:             msg.ChatId,
			Text:               msg.Text[utils.Max(i*msgLength-16, 0):utils.Min((i+1)*msgLength, len(msg.Text))],
			ParseMode:          msg.ParseMode,
			DisableLinkPreview: msg.DisableLinkPreview,
			Msg:                "",
		}
		code, _, err := HttpClient.Post(config.ApiUrl).JSON(cur).Bytes()
		if err != nil || code != fiber.StatusOK {
			return code, err
		}
	}
	return fiber.StatusOK, nil
}
