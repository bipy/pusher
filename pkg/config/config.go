package config

import (
    "os"
)

const (
    BaseUrl        = "https://api.telegram.org/bot"
    ParseMode      = "Markdown"
    ApiSendMessage = "/sendMessage"
)

var (
    Token  string
    ChatId string
    ApiUrl string
    Key    string
)

func init() {
    Token = os.Getenv("TG_TOKEN")
    ChatId = os.Getenv("CHAT_ID")
    Key = os.Getenv("SECURE_KEY")

    ApiUrl = BaseUrl + Token + ApiSendMessage
}
