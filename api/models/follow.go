package models

type Follow struct {
	UserID       string `bson:"userId" json:"userId"`
	UserFollowID string `bson:"userFollowId" json:"userFollowId"`
}
