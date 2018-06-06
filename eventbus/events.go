package eventbus

import (
	"github.com/jinzhu/gorm"
)

// Event model
type Event struct {
	gorm.Model
	EventType string `json:"event_type"`
	Payload   string `json:"payload"`
}

// CreateEvent will persist an event and deliver it to applicable subscriptions
func CreateEvent(eventType string, payload string) Event {
	event := Event{
		EventType: eventType,
		Payload:   payload,
	}

	db.Create(&event)

	startDeliveries(&event)

	return event
}
