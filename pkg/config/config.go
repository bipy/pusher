package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// Telegram API constants
	BaseURL        = "https://api.telegram.org/bot"
	ParseMode      = "MarkdownV2"
	APISendMessage = "/sendMessage"

	// Message constraints
	MaxRuneLength = 1000 // Max runes per message chunk (Telegram limit is 4096 bytes)
)

var (
	// Telegram configuration
	Token  string
	ChatID int
	APIURL string

	// Security key for optional authentication
	Key string
)

func init() {
	// Validate and load Telegram token
	Token = os.Getenv("TG_TOKEN")
	if Token == "" {
		panic("TG_TOKEN environment variable is required")
	}

	// Validate and load chat ID
	chatIDStr := os.Getenv("CHAT_ID")
	if chatIDStr == "" {
		panic("CHAT_ID environment variable is required")
	}

	var err error
	ChatID, err = strconv.Atoi(chatIDStr)
	if err != nil {
		panic(fmt.Sprintf("invalid CHAT_ID: %v", err))
	}

	// Load optional security key
	Key = os.Getenv("SECURE_KEY")

	// Build API URL
	APIURL = BaseURL + Token + APISendMessage
}
