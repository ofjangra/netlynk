package helpers

import (
	"context"
	"errors"
	"time"

	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	username := user.Username

	userPhone := user.Phone

	foundUser := bson.M{}

	usernameExists := userCollection.FindOne(ctx, bson.M{"username": username})

	usernameExists.Decode(&foundUser)

	if foundUser["username"] == username {
		return errors.New("username already taken")
	}

	phoneExists := userCollection.FindOne(ctx, bson.M{"phone": userPhone})

	phoneExists.Decode(&foundUser)

	if foundUser["phone"] == userPhone {
		return errors.New("Phone number already used")
	}

	_, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		return errors.New("Failed to create account")
	}

	return nil
}

func GetuserByUsername(username string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	result := userCollection.FindOne(ctx, bson.M{"username": username})

	return result

}
