package db

import (
	"context"
	"log"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListEvents(ID string, page int64) ([]*models.GetEvents, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("events")

	var results []*models.GetEvents

	condition := bson.M{
		"userId": bson.M{"$ne": ID},
	}

	config := options.Find()
	config.SetLimit(20)
	config.SetSort(bson.D{{Key: "date", Value: -1}})
	config.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, config)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var entry models.GetEvents
		err := cursor.Decode(&entry)
		if err != nil {
			return results, false
		}
		results = append(results, &entry)
	}
	return results, true

}
