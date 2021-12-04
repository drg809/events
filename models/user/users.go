package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Surname  string             `bson:"surname" json:"surname,omitempty"`
	Date     time.Time          `bson:"date" json:"date,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Bio      string             `bson:"bio" json:"bio,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	Web      string             `bson:"web" json:"web,omitempty"`
}
