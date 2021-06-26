package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertParticipation(t models.Participation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("participations")

	entry := bson.M{
		"userId":  t.UserID,
		"eventId": t.EventID,
		"details": t.Details,
	}

	_, err := col.InsertOne(ctx, entry)
	if err != nil {
		return false, err
	}

	return true, nil
}
