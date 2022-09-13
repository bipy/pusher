package models

type TgMessage struct {
	ChatId             int    `json:"chat_id"`
	Text               string `json:"text"`
	ParseMode          string `json:"parse_mode"`
	DisableLinkPreview bool   `json:"disable_web_page_preview"`
}

type ReqMessage struct {
	Text     string `json:"text"`
	Preview  bool   `json:"preview"`
	Msg      string `json:"msg"`
	Markdown bool   `json:"markdown"`
}
