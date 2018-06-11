package main

import "gopkg.in/mgo.v2/bson"

type Book struct {
	ID     bson.ObjectId       `bson:"_id" json:"id"`
	Title  string              `json:"title" binding:"required"`
	Author string              `json:"author"`
	Genre  string              `json:"genre"`
	Read   bool                `json:"read"`
	Links  []map[string]string `json:"links,omitempty"`
}
