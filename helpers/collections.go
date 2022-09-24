package helpers

import (
	"github.com/ofjangra/netlynk_server/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME = "netlynk"

const usersCollName string = "Users"

const linksCollName string = "Links"

var userCollection *mongo.Collection

var linksCollection *mongo.Collection

func init() {

	client := db.DbInstance()

	userCollection = client.Database(DBNAME).Collection(usersCollName)

	linksCollection = client.Database(DBNAME).Collection(linksCollName)

}

// create user
