package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyEntry(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("users")

	entry := make(map[string]interface{})

	if len(u.Name) > 0 {
		entry["name"] = u.Name
	}

	if len(u.Surname) > 0 {
		entry["surname"] = u.Surname
	}

	if len(u.Email) > 0 {
		entry["email"] = u.Email
	}

	if len(u.Avatar) > 0 {
		entry["avatar"] = u.Avatar
	}

	if len(u.Location) > 0 {
		entry["location"] = u.Location
	}

	if len(u.Web) > 0 {
		entry["web"] = u.Web
	}

	if len(u.Bio) > 0 {
		entry["bio"] = u.Bio
	}

	if len(u.Banner) > 0 {
		entry["banner"] = u.Banner
	}

	entry["date"] = u.Date

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
