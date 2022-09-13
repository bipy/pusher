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
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	data, err = utils.Gzip(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, config.ApiURL, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)
	req.Header.Set("Content-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errMsg, _ := io.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("status code: %d, error: %s", resp.StatusCode, errMsg))
	}
	return nil
}
