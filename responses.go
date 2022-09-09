package goform

import (
	"encoding/json"
	"io"
)

type WebhookResponse struct {
	EventID      string `json:"event_id"`
	EventType    string `json:"event_type"`
	FormResponse Form   `json:"form_response"`
}

func ParseWebhookResponse(bd io.ReadCloser) (interface{}, error) {
	var out WebhookResponse
	if err := json.NewDecoder(bd).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}
