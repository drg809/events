package db

import (
	"context"
	"log"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListUsers(page int64, search string) ([]*models.ListUsers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("users")

	var results []*models.ListUsers

	config := options.Find()
	config.SetLimit(20)
	config.SetSkip((page - 1) * 20)

	query := bson.M{
		//"name": bson.M{"$regex": `($i)` + search},
	}

	cursor, err := col.Find(ctx, query, config)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var entry models.ListUsers
		err := cursor.Decode(&entry)
		if err != nil {
			return results, false
		}
		results = append(results, &entry)
	}
	return results, true
}
