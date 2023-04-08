package database

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	db_config "github.com/ofjangra/netlynk_server/internal/config/db"
	"github.com/ofjangra/netlynk_server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(user *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	usernameExists := db_config.UserCollection.FindOne(ctx, bson.M{"username": user.Username})

	if usernameExists.Err() == nil {
		return errors.New("username already taken")
	}

	emailExists := db_config.UserCollection.FindOne(ctx, bson.M{"email": user.Email})

	if emailExists.Err() == nil {
		return errors.New("email already taken")
	}

	_, err := db_config.UserCollection.InsertOne(ctx, user)

	if err != nil {
		return errors.New("failed to create account")
	}

	return nil
}

func GetuserByUsername(username string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	username = strings.ToLower(username)

	result := db_config.UserCollection.FindOne(ctx, bson.M{"username": username})

	return result

}

func GetUserById(id string) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, errors.New("failed to get user")
	}

	result := db_config.UserCollection.FindOne(ctx, bson.M{"_id": userId})

	return result, nil
}

func EditProfile(id string, body bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		fmt.Println(idErr)
		return errors.New("failed to update profile")
	}

	fmt.Println("body : ", body)

	delete(body, "_id")

	username, isUsername := body["username"]

	if isUsername {
		userameIsValid := regexp.MustCompile(`^[A-Za-z0-9._-]{3,16}$`).MatchString(username.(string))

		if !userameIsValid {

			return errors.New("username can only contain alphabets, numbers, '-', '.', and '_'")

		} else if len(username.(string)) < 3 || len(username.(string)) > 16 {

			return errors.New("username length must be 3 to 16")

		}

		usernameExists := db_config.UserCollection.FindOne(ctx, bson.M{"username": username.(string)})

		if usernameExists.Err() == nil {
			return errors.New("username already taken")
		}

	}

	email, isEmail := body["email"]

	if isEmail {
		emailIsValid := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(email.(string))

		if !emailIsValid {

			return errors.New("invalid Email")

		}

		emailExists := db_config.UserCollection.FindOne(ctx, bson.M{"email": email})

		if emailExists.Err() == nil {
			return errors.New("email already taken")
		}

	}

	_, err := db_config.UserCollection.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": body}, options.Update().SetUpsert(true))

	if err != nil {
		fmt.Println("update Err : ", err)
		return errors.New("failed to update profile")
	}

	return nil
}
