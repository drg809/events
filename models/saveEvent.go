package models

import "time"

type SaveEvent struct {
	UserID        string    `bson:"userId" json:"userId,omitempty"`
	Name          string    `bson:"name" json:"name"`
	Detail        string    `bson:"detail" json:"detail"`
	Date          time.Time `bson:"date" json:"date,omitempty"`
	DateEnd       time.Time `bson:"dateEnd" json:"dateEnd,omitempty"`
	Type          bool      `bson:"type" json:"type,omitempty"`
	Photo         string    `bson:"photo" json:"photo,omitempty"`
	TParticipants int       `bson:"tParticipants" json:"tParticipants,omitempty"`
}
