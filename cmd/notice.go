package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func Notice(cfg *Config, content []byte) error {
	body, err := json.Marshal(map[string]string{
		"text": string(content),
	})
	if err != nil {
		return err
	}
	res, err := http.Post(cfg.Slack.WebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("err: " + res.Status)
	}
	return nil
}
