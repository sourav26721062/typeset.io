package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Post - Model description
type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`
	Content   []Paragraph   `json:"content,omitempty" bson:"content,omitempty"`
	CreatedOn time.Time     `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	Link      string        `json:"link,omitempty" bson:"link,omitempty"`
}
