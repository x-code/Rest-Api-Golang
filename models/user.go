package models

//developed by R.Firandika | x-c0de

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name   	string        `json:"name" bson:"name"`
		Age    	int           `json:"age" bson:"age"`
		Photo 	string        `json:"photo" bson:"photo"`
	}
)
