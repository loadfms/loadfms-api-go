package models

import "gopkg.in/mgo.v2/bson"

type Device struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `bson:"name" json:"name"`
	State int           `bson:"state" json:"state"`
	Port  int           `bson:"port" json:"port"`
}
