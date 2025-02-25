package helpers

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// รับ text format และ LINE Token เพื่อส่งแจ้งเตือนผ่าน LINE notification
func SendLineNotification(text, lineToken string) error {
	tokens := strings.Split(lineToken, ",")
	for _, token := range tokens {

		data := url.Values{}
		data.Set("message", text)
		lineUrl := "https://notify-api.line.me/api/notify"
		req, err := http.NewRequest("POST", lineUrl, strings.NewReader(data.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return errors.New(string(body))
		}

	}

	return nil
}
