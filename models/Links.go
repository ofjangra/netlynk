package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Links struct {
	Title      string `json:"title"`
	URL        string `json:"url"`
	Created_by primitive.ObjectID
}
