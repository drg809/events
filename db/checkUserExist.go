package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
