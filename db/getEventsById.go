package db

import (
	"context"
	"fmt"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEventsById(ID string) (*models.GetEvents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("events")

	var event *models.GetEvents
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&event)
	if err != nil {
		fmt.Println("Registro no encotrado " + err.Error())
		return event, err
	}

	return event, nil

}
