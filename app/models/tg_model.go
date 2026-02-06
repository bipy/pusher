package models

// TgMessage represents a message to be sent via Telegram API
type TgMessage struct {
	ChatID             int    `json:"chat_id"`
	Text               string `json:"text"`
	ParseMode          string `json:"parse_mode"`
	DisableLinkPreview bool   `json:"disable_web_page_preview"`
}

// ReqMessage represents an incoming message request
type ReqMessage struct {
	Text     string `json:"text"`
	Preview  bool   `json:"preview"`
	Msg      string `json:"msg"`      // Alternative field for text
	Markdown bool   `json:"markdown"` // Enable markdown parsing
}
