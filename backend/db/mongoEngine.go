package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collections *mongo.Collection
)

func InitMongoDB() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:qwertyuiop@localhost:27017/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collections = client.Database("mirror").Collection("catalog")
}
