package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func InitDB() {
	ATLAS_URI := "mongodb+srv://spars01:H0YXCAGHoUihHcSZ@cluster0.wuezj.mongodb.net/prevention_productivity?retryWrites=true&w=majority"
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

	Db = client.Database("prevention_productivity")
}

func CloseDB() {
	err := Db.Client().Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
