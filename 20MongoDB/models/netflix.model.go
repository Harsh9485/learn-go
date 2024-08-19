package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID 		  primitive.ObjectID `json: "_id, omitempty" bson: "_id, omitempty"`
	MovieNeme string 			 `json: "movie_name, omitempty"`
	IsWatched bool 				 `json: "is_watched, omitempty"`
}