package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Statement struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Tags        []string      `bson:"tags" json:"tags"`
	Description string        `bson:"description" json:"description"`
	Date        time.Time     `bson:"date" json:"date"`
	Value       float64       `bson:"value" json:"value"`
	Type        int           `bson:"type" json:"type"`
}
