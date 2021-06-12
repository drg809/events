package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEventsById(ID string) (*models.GetEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("events")

	var result *models.GetEvents

	condition := bson.M{
		"_id": ID,
	}

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil

}
