package eventbus

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	EventType string `json:"event_type"`
	Payload   string `json:"payload"`
}

func CreateEvent(event_type string, payload string) Event {
	event := Event{
		EventType: event_type,
		Payload:   payload,
	}

	db.Create(&event)

	startDeliveries(&event)

	return event
}
