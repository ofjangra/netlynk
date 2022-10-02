package helpers

import (
	"context"
	"errors"
	"time"

	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	username := user.Username

	userEmail := user.Email

	foundUser := bson.M{}

	usernameExists := userCollection.FindOne(ctx, bson.M{"username": username})

	usernameExists.Decode(&foundUser)

	if foundUser["username"] == username {
		return errors.New("username already taken")
	}

	emailExists := userCollection.FindOne(ctx, bson.M{"email": userEmail})

	emailExists.Decode(&foundUser)

	if foundUser["email"] == userEmail {
		return errors.New("Email already used")
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

func GetUserById(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, idErr
	}

	result := userCollection.FindOne(ctx, bson.M{"_id": userId})

	return result, nil
}

func EditProfile(id string, update *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("Failed to update profile")
	}

	user := new(models.User)

	thisUser := userCollection.FindOne(ctx, bson.M{"_id": userId})

	decodeErr := thisUser.Decode(&user)

	if decodeErr != nil {
		return errors.New("Failed to update profile")
	}

	query := bson.M{}

	if update.Username != user.Username {

		res := userCollection.FindOne(ctx, bson.M{"username": update.Username})
		if res.Err() == nil {
			return errors.New("Username already taken")
		}
		query["username"] = update.Username
	}

	if update.Email != user.Email {
		res := userCollection.FindOne(ctx, bson.M{"email": update.Email})
		if res.Err() == nil {
			return errors.New("Email already taken")
		}
		query["email"] = update.Email
	}

	if update.Bio != user.Bio {
		query["bio"] = update.Bio
	}

	_, updateErr := userCollection.UpdateByID(ctx, userId, bson.M{"$set": query})

	if updateErr != nil {
		return errors.New("Failed to update profile")
	}

	return nil
}

func EditProfilePhoto(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("Failed to update profile")
	}

	_, updateErr := userCollection.UpdateByID(ctx, userId, bson.M{"$set": update})

	if updateErr != nil {
		return errors.New("Failed to update profile")
	}

	return nil
}
