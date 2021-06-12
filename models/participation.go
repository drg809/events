package models

type Participation struct {
	UserID  string `bson:"userId" json:"userId"`
	EventID string `bson:"eventId" json:"eventId"`
	Details string `bson:"details" json:"details"`
}
