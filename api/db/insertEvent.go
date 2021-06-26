package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertEvent(t models.Event) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("events")

	entry := bson.M{
		"userId": t.UserID,
		"name":   t.Name,
		"detail": t.Detail,
		"type":   t.Type,
		"photo":  t.Photo,
		"date":   t.Date,
	}

	result, err := col.InsertOne(ctx, entry)
	if err != nil {
		return string(""), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
