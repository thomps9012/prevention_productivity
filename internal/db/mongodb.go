package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func InitDB() {
	// change on deployment
	ATLAS_URI := os.Getenv("ATLAS_URI")
	// Set client options
	if ATLAS_URI == "" {
		ATLAS_URI = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(ATLAS_URI)

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

	// change on deployment
	dbName := os.Getenv("DB_NAME")
	Db = client.Database(dbName)
}

func CloseDB() {
	err := Db.Client().Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
