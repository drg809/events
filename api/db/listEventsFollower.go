package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListEventsFollowers(ID string, pagina int) ([]models.ListEventsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("follows")

	skip := (pagina - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "events",
			"localField":   "UserFollowID",
			"foreingField": "userId",
			"as":           "event",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$event"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"event.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	var result []models.ListEventsFollowers
	cursor, err := col.Aggregate(ctx, conditions)
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
