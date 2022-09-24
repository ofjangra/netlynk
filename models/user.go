package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Socialmedia struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Youtube   string `json:"youtube"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
	Github    string `json:"github"`
}
type User struct {
	ID          primitive.ObjectID `json:"_id"`
	UserId      string             `json:"user_id" bson:"user_id"`
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Phone       string             `json:"phone"`
	Bio         string             `json:"bio"`
	Password    string             `json:"password"`
	Socialmedia Socialmedia        `json:"socialmedia"`
}
