package eventSummaries

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"strings"
	"time"
	"github.com/google/uuid"
)

type EventSummary struct {
	ID            string `json:"id" bson:"_id"`
	EventID       string `json:"event_id" bson:"event_id"`
	UserID        string `json:"user_id" bson:"user_id"`
	AttendeeCount int    `json:"attendee_count" bson:"attendee_count"`
	Challenges    string `json:"challenges"`
	Successes     string `json:"successes"`
	Improvements  string `json:"improvements"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
}

func (e *EventSummary) Create() {
	collection := database.Db.Collection("event_summaries")
	e.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		panic(err)
	}
}

func (e *EventSummary) Update(id string) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"event_id", e.EventID},
			{"user_id", e.UserID},
			{"attendee_count", e.AttendeeCount},
			{"challenges", e.Challenges},
			{"successes", e.Successes},
			{"improvements", e.Improvements},
			{"status", e.Status},
			{"updated_at", e.UpdatedAt},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func (e *EventSummary) Delete(id string) {
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{"_id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}

func (e *EventSummary) Approve(id string) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "approved"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func (e *EventSummary) Reject(id string) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "rejected"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}