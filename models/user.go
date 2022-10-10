package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	PhotoUrl string             `json:"photo_url" bson:"photo_url"`
	Email    string             `json:"email" bson:"email"`
	Bio      string             `json:"bio" bson:"bio"`
	Password string             `json:"-" bson:"password"`
}
