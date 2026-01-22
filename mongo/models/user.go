package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `bson:"name"`
	Gender string        `bson:"gender"`
	Age    int           `bson:"age"`
}

type UserResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
