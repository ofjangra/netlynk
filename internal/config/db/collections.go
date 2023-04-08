package config_db

import "go.mongodb.org/mongo-driver/mongo"

const DBNAME = "netlynk"

const usersCollName string = "Users"

const linksCollName string = "Links"

var UserCollection *mongo.Collection

var LinksCollection *mongo.Collection

func GetCollections() {

	client := DbInstance()

	UserCollection = client.Database(DBNAME).Collection(usersCollName)

	LinksCollection = client.Database(DBNAME).Collection(linksCollName)

}
