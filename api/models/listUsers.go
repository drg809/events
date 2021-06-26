package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListUsers struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name    string             `bson:"name" json:"name,omitempty"`
	Surname string             `bson:"surname" json:"surname,omitempty"`
	Avatar  string             `bson:"avatar" json:"avatar,omitempty"`
}
