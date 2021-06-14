package db

import (
	"context"
	"fmt"
	"time"

	"github.com/drg809/events/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListParticipations(ID string, page int64, search string, tipo string) ([]*models.ListParticipations, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dbObj := MongoCN.Database("events")
	col := dbObj.Collection("participations")

	var results []*models.ListParticipations

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		//"name": bson.M{"$regex": `($i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var find, include bool

	for cur.Next(ctx) {
		var s models.GetEvents
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Participation
		r.UserID = ID
		r.EventID = s.ID.Hex()

		include = false
		/* si queremos discriminar por participados y no participados tenemos que hacer el listado */

		find, _ = CheckParticipation(r)
		if tipo == "new" && !find {
			include = true
		}
		if tipo == "participate" && find {
			include = true
		}

		if r.UserID == ID {
			include = false
		}

		if include {
			var p models.ListParticipations
			p.Date = s.Date
			p.Name = s.Name
			p.UserID = r.EventID
			p.EventID = r.EventID
			results = append(results, &p)
		}

	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
