package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func Notice(cfg *Config, content []byte) (err error) {
	switch strings.ToLower(cfg.Type) {
	case "slack":
		err = noticeToSlack(&cfg.Slack, content)
	case "stdout":
		fallthrough
	default:
		fmt.Print(string(content))
	}
	return err
}

func noticeToSlack(cfg *slackConfig, content []byte) error {
	body, err := json.Marshal(map[string]string{
		"text": string(content),
	})
	if err != nil {
		return err
	}
	res, err := http.Post(cfg.WebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("err: " + res.Status)
	}
	return nil
}
