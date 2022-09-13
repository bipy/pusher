package config

import (
	"os"
	"strconv"
)

const (
	BaseURL        = "https://api.telegram.org/bot"
	ParseMode      = "MarkdownV2"
	ApiSendMessage = "/sendMessage"
)

var (
	Token  string
	ChatId int
	ApiURL string
	Key    string
)

func init() {
	var err error
	ChatId, err = strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		panic("invalid chat id")
	}
	Token = os.Getenv("TG_TOKEN")
	Key = os.Getenv("SECURE_KEY")

	ApiURL = BaseURL + Token + ApiSendMessage
}
