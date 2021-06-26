package db

import (
	"context"
	"time"

	"github.com/drg809/events/models"
)

func InsertFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("follows")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
