package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteEvent(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("events")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userId": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
