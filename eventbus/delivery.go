package eventbus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func startDeliveries(event *Event) {
	var subscriptions []Subscription
	db.Where("event_type = ?", event.EventType).Find(&subscriptions)

	for _, subscription := range subscriptions {
		deliverEvent(event, subscription.Endpoint)
	}
}

type Payload struct {
	Payload string `json:"payload"`
}

func deliverEvent(event *Event, path string) {
	payload := Payload{
		Payload: event.Payload,
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, _ := client.Do(req)

	fmt.Println(resp)
}
