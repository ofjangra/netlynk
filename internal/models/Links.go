package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	URL        string             `json:"url" bson:"url"`
	Created_by primitive.ObjectID `json:"created_by" bson:"created_by"`
	CreatedOn  primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn  primitive.DateTime `json:"updated_on" bson:"updated_on"`
}
