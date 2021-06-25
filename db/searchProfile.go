package db

import (
	"context"
	"fmt"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.ListUsers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	dbObject := MongoCN.Database("events")
	col := dbObject.Collection("users")

	var profile models.ListUsers
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	if err != nil {
		fmt.Println("Pegistro no encotrado " + err.Error())
		return profile, err
	}

	return profile, nil
}
