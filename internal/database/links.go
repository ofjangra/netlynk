package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	db_config "github.com/ofjangra/netlynk_server/internal/config/db"
	"github.com/ofjangra/netlynk_server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateALink(link models.Links) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	_, err := db_config.LinksCollection.InsertOne(ctx, link)

	if err != nil {
		return errors.New("failed to create link")
	}

	return nil
}

func EditALink(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update the link")
	}

	_, updateErr := db_config.LinksCollection.UpdateByID(ctx, linkId, bson.M{"$set": update})

	if updateErr != nil {
		return errors.New("failed to update the link")
	}

	return nil
}

func DeleteALink(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to delete the link")
	}

	deleteResult, deleteErr := db_config.LinksCollection.DeleteOne(ctx, bson.M{"_id": linkId})

	if deleteErr != nil || deleteResult.DeletedCount < 1 {
		return errors.New("failed to update the link")
	}

	return nil

}

func GetALink(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, errors.New("failed to fetch link")
	}

	result := db_config.LinksCollection.FindOne(ctx, bson.M{"_id": linkId})

	return result, nil
}

// Get All Links of a User

func GetAllLinks(createrId primitive.ObjectID) ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	fmt.Println("creator id", createrId)

	// id, idErr := primitive.o
	// fmt.Println(id)
	// if idErr != nil {
	// 	fmt.Println("id err : ", idErr)
	// 	return nil, errors.New("failed to fetch links")
	// }

	links := []primitive.M{}

	opts := options.Find().SetSort(bson.M{"created_on": -1})

	cur, curErr := db_config.LinksCollection.Find(ctx, bson.M{"created_by": createrId}, opts)

	if curErr != nil {
		fmt.Println("curr error : ", curErr)
		return nil, errors.New("failed to fetch links")
	}

	for cur.Next(context.Background()) {
		link := bson.M{}

		err := cur.Decode(&link)

		if err != nil {
			fmt.Println("link decode err : ", err)
			return nil, errors.New("failed to fetch links")
		}

		links = append(links, link)
	}

	defer cur.Close(context.Background())
	return links, nil
}
