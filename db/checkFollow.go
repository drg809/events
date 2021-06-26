package db

import (
	"context"
	"fmt"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("follows")

	condition := bson.M{
		"userFollowId": t.UserFollowID,
		"userId":       t.UserID,
	}

	var result models.Follow
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
