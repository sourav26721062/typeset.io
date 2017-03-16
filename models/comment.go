package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Comment - Model description
type Comment struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Text      string        `json:"text,omitempty" bson:"text,omitempty"`
	CreatedOn time.Time     `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
}
