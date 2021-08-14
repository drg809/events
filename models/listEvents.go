package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetEvents struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Name    string             `bson:"name" json:"name,omitempty"`
	Detail  string             `bson:"detail" json:"detail,omitempty"`
	Date    time.Time          `bson:"date" json:"date"`
	DateEnd time.Time          `bson:"dateEnd" json:"dateEnd,omitempty"`
	Type    bool               `bson:"type" json:"type"`
	Photo   string             `bson:"photo" json:"photo"`
}
