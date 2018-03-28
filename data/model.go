package data

import (
	"time"
)

type UserChange struct {
	UserId string    `bson:"userId" json:"userId"`
	Time   time.Time `bson:"time" json:"time"`
	Field  string    `bson:"field" json:"field"`
	Value  string    `bson:"value" json:"value"`
}
