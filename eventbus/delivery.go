package eventbus

import (
	"bytes"
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

func deliverEvent(event *Event, path string) {
	var json = bytes.NewBuffer([]byte(event.Payload))

	req, _ := http.NewRequest("POST", path, json)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, _ := client.Do(req)

	fmt.Println(resp)
}
