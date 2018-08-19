package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNoticeToSlack(t *testing.T) {
	var requestBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		requestBody = string(body)
		return
	}))
	defer server.Close()

	cfg := &slackConfig{
		WebhookURL: server.URL,
	}

	type args struct {
		cfg     *slackConfig
		content []byte
	}
	tests := []struct {
		name     string
		args     args
		expected string
		wantErr  bool
	}{
		{
			name: "basic",
			args: args{
				cfg:     cfg,
				content: bytes.NewBufferString("basic").Bytes(),
			},
			expected: "{\"text\":\"basic\"}",
			wantErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := noticeToSlack(test.args.cfg, test.args.content); (err != nil) != test.wantErr {
				t.Errorf("Notice() error = %#v, wantErr %v", err, test.wantErr)
			}
			if requestBody != test.expected {
				t.Errorf("\nexpected:\n%s\n-----\nrequested body:\n%s", test.expected, requestBody)
			}
		})
	}
}
