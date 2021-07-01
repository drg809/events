package db

import (
	"context"
	"log"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListUsers(ID string, page int64, search string, tipo string) ([]*models.ListUsers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("users")

	var results []*models.ListUsers

	config := options.Find()
	config.SetLimit(10)
	config.SetSkip((page - 1) * 10)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, config)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.ListUsers
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Follow
		r.UserID = ID
		r.UserFollowID = s.ID.Hex()

		incluir = false

		encontrado, err = CheckFollow(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UserFollowID == ID {
			incluir = false
		}

		if incluir {
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
