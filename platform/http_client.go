package platform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"pusher/app/models"
	"pusher/pkg/config"
	"pusher/pkg/utils"
	"strings"
)

// max length 4096 byte
const runeLength int = 1000

func Push(message []rune, disableLinkPreview bool, ip string) error {
	for i := 0; i*runeLength < len(message); i++ {
		s := message[i*runeLength : utils.Min((i+1)*runeLength, len(message))]
		msg := &models.TgMessage{
			ChatId:             config.ChatId,
			Text:               fmt.Sprintf("%s\n\n%s\n*From: %s*", string(s), strings.Repeat("\\-", 10), ip),
			ParseMode:          config.ParseMode,
			DisableLinkPreview: disableLinkPreview,
		}
		err := send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func send(msg *models.TgMessage) error {
	buf, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Post(config.ApiURL, echo.MIMEApplicationJSONCharsetUTF8, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errMsg, _ := io.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("status code: %d, error: %s", resp.StatusCode, errMsg))
	}
	return nil
}
