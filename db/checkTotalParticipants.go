package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckTotalParticipants(t models.Participation) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("participations")

	query := bson.M{
		"eventId": t.EventID,
	}

	var event models.Event
	col1 := dbObj.Collection("participations")

	objID, _ := primitive.ObjectIDFromHex(t.EventID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	err := col1.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		return false, err.Error()
	}

	// cur, err := col.Find(ctx, query, config)
	cur, err := col.CountDocuments(ctx, query)
	if err != nil {
		return false, err.Error()
	}
	if int64(event.TParticipants) == cur {
		return false, "Evento sin cupo."
	}

	return true, ""
}
