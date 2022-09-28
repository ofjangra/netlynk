package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Title      string             `json:"title" bson:"title"`
	URL        string             `json:"url" bson:"url"`
	Created_by primitive.ObjectID `json:"created_by" bson:"created_by"`
}
