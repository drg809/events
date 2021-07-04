package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListEventsFollowers struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId       string             `bson:"userId" json:"userId,omitempty"`
	UserFollowId string             `bson:"userFollowId" json:"userFollowId,omitempty"`
	Event        struct {
		ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		Name   string             `bson:"name" json:"name,omitempty"`
		Detail string             `bson:"detail" json:"detail,omitempty"`
		Date   time.Time          `bson:"date" json:"date"`
		Type   bool               `bson:"type" json:"type"`
		Photo  string             `bson:"photo" json:"photo"`
	}
}
