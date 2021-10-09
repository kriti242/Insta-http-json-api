package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//user structure
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

//post structure
type Post struct {
	ID            bson.ObjectId `json:"_id" bson:"_id"`
	Caption       string        `json:"caption" bson:"caption"`
	ImageUrl      string        `json:"image_url" bson:"image_url"`
	UserId        string        `json:"user_id" bson:"user_id"`
	PostTimestamp time.Time     `json:"post_timestamp" bson:"post_timestamp"`
}
