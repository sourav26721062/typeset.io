package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Paragraph - Model description
type Paragraph struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Text      string        `json:"text,omitempty" bson:"text,omitempty"`
	Comments  []Comment     `json:"comments,omitempty" bson:"comments,omitempty"`
	CreatedOn time.Time     `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
}
