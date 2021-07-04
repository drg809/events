package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"userId" json:"userId,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Detail string             `bson:"detail" json:"detail"`
	Date   time.Time          `bson:"date" json:"date,omitempty"`
	Type   bool               `bson:"type" json:"type,omitempty"`
	Photo  string             `bson:"photo" json:"photo,omitempty"`
}
