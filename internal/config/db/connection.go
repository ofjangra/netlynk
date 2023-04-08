package config_db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	// if envErr := godotenv.Load(".env"); envErr != nil {
	// 	log.Fatal("Failed to load db enviorenment variables")
	// }

	DBURI := os.Getenv("DBURI")

	clientOptions := options.Client().ApplyURI(DBURI)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection is ready")

	return client
}
