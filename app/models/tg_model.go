package models

type Message struct {
    ChatId             string `json:"chat_id"`
    Text               string `json:"text"`
    ParseMode          string `json:"parse_mode"`
    DisableLinkPreview bool   `json:"disable_web_page_preview"`
}
