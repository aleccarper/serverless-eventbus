package eventbus

import (
	"github.com/jinzhu/gorm"
)

// Subscription model
type Subscription struct {
	gorm.Model
	EventType string `json:"event_type" sql:"index"`
	Endpoint  string `json:"endpoint"`
}

// CreateSubscription persists a subscription
func CreateSubscription(eventType string, endpoint string) Subscription {
	subscription := Subscription{
		EventType: eventType,
		Endpoint:  endpoint,
	}

	db.Create(&subscription)

	return subscription
}

// DeleteSubscription deletes a subscription by ID
func DeleteSubscription(id string) Subscription {
	var subscription Subscription
	db.First(&subscription, id)
	db.Delete(&subscription)
	return subscription
}
