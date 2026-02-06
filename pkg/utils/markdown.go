package utils

import "strings"

// markdownReplacer escapes special characters for Telegram MarkdownV2
var markdownReplacer = strings.NewReplacer(
	"\\", "\\\\",
	"_", "\\_",
	"*", "\\*",
	"[", "\\[",
	"]", "\\]",
	"(", "\\(",
	")", "\\)",
	"~", "\\~",
	"`", "\\`",
	">", "\\>",
	"#", "\\#",
	"+", "\\+",
	"-", "\\-",
	"=", "\\=",
	"|", "\\|",
	"{", "\\{",
	"}", "\\}",
	".", "\\.",
	"!", "\\!",
)

// EscapeMarkdown escapes special characters for Telegram MarkdownV2 format
func EscapeMarkdown(s string) string {
	return markdownReplacer.Replace(s)
}
