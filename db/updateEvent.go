package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateEvent(e models.Event, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("events")

	entry := make(map[string]interface{})

	if len(e.Name) > 0 {
		entry["name"] = e.Name
	}

	if len(e.Detail) > 0 {
		entry["detail"] = e.Detail
	}

	if !e.Date.IsZero() {
		entry["email"] = e.Date
	}

	if e.Type || !e.Type {
		entry["type"] = e.Type
	}

	if len(e.Photo) > 0 {
		entry["location"] = e.Photo
	}

	updtString := bson.M{
		"$set": entry,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
