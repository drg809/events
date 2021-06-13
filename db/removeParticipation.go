package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RemoveParticipation(t models.Participation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("participations")

	condition := bson.M{
		"eventId": t.EventID,
		"userId":  t.UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil
}
