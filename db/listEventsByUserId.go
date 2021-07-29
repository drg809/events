package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListEventsByUserId(ID string, page int) ([]models.ListEventsUser, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("users")

	skip := (page - 1) * 10
	objID, _ := primitive.ObjectIDFromHex(ID)

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"_id": objID}})
	conditions = append(conditions, bson.M{"$addFields": bson.M{"objId": bson.M{"$toString": "$_id"}}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "events",
			"localField":   "objId",
			"foreignField": "userId",
			"as":           "event",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$event"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"event.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 10})

	cursor, _ := col.Aggregate(ctx, conditions)
	var result []models.ListEventsUser
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
