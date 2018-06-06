package eventbus

import (
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	EventType string `json:"event_type" sql:"index"`
	Endpoint  string `json:"endpoint"`
}

func CreateSubscription(eventType string, endpoint string) Subscription {
	subscription := Subscription{
		EventType: eventType,
		Endpoint:  endpoint,
	}

	db.Create(&subscription)

	return subscription
}

func DeleteSubscription(id string) Subscription {
	var subscription Subscription
	db.First(&subscription, id)
	db.Delete(&subscription)
	return subscription
}
