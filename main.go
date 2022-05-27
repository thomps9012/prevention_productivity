package main
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("prevention_db")
	usersCollection := database.Collection("users")
	notesCollection := database.Collection("notes")
	logsCollection := database.Collection("productivity_logs")
	grantsCollection := database.Collection("grant_programs")
	schoolreportsCollection := database.Collection("school_reports")
	eventsCollection := database.Collection("events")
	eventSummariesCollection := database.Collection("event_summaries")
	contactsCollection := database.Collection("contacts")
}