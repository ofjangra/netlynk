package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Username string             `json:"username,omitempty" bson:"username" validate:"omitempty,min=3,max=16"`
	PhotoUrl string             `json:"photo_url" bson:"photo_url"`
	Email    string             `json:"email,omitempty" bson:"email" validate:"omitempty,email"`
	Bio      string             `json:"bio" bson:"bio"`
	Password string             `json:"password,omitempty" bson:"password" validate:"omitempty,min=8,max=34"`
}

// regexp=^[A-Za-z0-9._-]{3,16}$"
