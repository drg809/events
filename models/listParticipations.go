package models

import "time"

type ListParticipations struct {
	UserID  string    `bson:"userId" json:"userId"`
	EventID string    `bson:"eventId" json:"eventId"`
	Details string    `bson:"details" json:"details,omitempty"`
	Name    string    `bson:"name" json:"name,omitempty"`
	Date    time.Time `bson:"date" json:"date"`
}
