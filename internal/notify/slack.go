package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// SendSlackNotificationInput ...
type SendSlackNotificationInput struct {
	WebhookURL string
	Message    string
	Channel    string
	Username   string
	IconEmoji  string
}

// SlackRequestBody ...
type SlackRequestBody struct {
	Text      string `json:"text"`
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

// SendSlackNotification ...
func SendSlackNotification(input *SendSlackNotificationInput) error {
	slackBody, err := json.Marshal(&SlackRequestBody{
		Text:      input.Message,
		Channel:   input.Channel,
		Username:  input.Username,
		IconEmoji: input.IconEmoji,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, input.WebhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(buf) != "ok" {
		return errors.New("Not OK")
	}

	return nil
}
