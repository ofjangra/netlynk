package helpers

import (
	"context"
	"time"

	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateALink(link models.Links) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	_, err := linksCollection.InsertOne(ctx, link)

	if err != nil {
		return err
	}

	return nil
}

func EditALink(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return idErr
	}

	updateResult, updateErr := linksCollection.UpdateByID(ctx, linkId, bson.M{"$set": update})

	if updateErr != nil || updateResult.UpsertedCount < 1 {
		return updateErr
	}

	return nil
}

func DeleteALink(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return idErr
	}

	deleteResult, deleteErr := linksCollection.DeleteOne(ctx, bson.M{"_id": linkId})

	if deleteErr != nil || deleteResult.DeletedCount < 1 {
		return deleteErr
	}

	return nil

}

func GetALink(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	linkId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, idErr
	}

	result := linksCollection.FindOne(ctx, bson.M{"_id": linkId})

	return result, nil
}

// Get All Links of a User

func GetAllLinks(createrId string) ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	id, idErr := primitive.ObjectIDFromHex(createrId)

	if idErr != nil {
		return nil, idErr
	}

	links := []primitive.M{}

	cur, curErr := linksCollection.Find(ctx, bson.M{"created_by": id})

	if curErr != nil {
		return nil, curErr
	}

	for cur.Next(context.Background()) {
		link := bson.M{}

		err := cur.Decode(&link)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	defer cur.Close(context.Background())

	return links, nil
}
