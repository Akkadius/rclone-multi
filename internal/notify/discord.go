package notify

import (
	"bytes"
	"encoding/json"
	"github.com/pterm/pterm"
	"net/http"
	"time"
)

type DiscordRequestBody struct {
	Content string `json:"content"`
}

func sendDiscordWebhook(msg string, webhook string) error {
	body, _ := json.Marshal(DiscordRequestBody{Content: msg})
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		pterm.Error.Println(err.Error())
		return err
	}
	return nil
}
