package eventSummaries

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventSummary struct {
	ID            string    `json:"id" bson:"_id"`
	EventID       string    `json:"event_id" bson:"event_id"`
	UserID        string    `json:"user_id" bson:"user_id"`
	Co_Planners   []*string `json:"co_planners" bson:"co_planners"`
	AttendeeCount *int      `json:"attendee_count" bson:"attendee_count"`
	Challenges    string    `json:"challenges" bson:"challenges"`
	Successes     string    `json:"successes" bson:"successes"`
	Improvements  string    `json:"improvements" bson:"improvements"`
	Status        string    `json:"status" bson:"status"`
	CreatedAt     string    `json:"created_at" bson:"created_at"`
	UpdatedAt     string    `json:"updated_at" bson:"updated_at"`
}

func (e *EventSummary) Create() (*EventSummary, error) {
	collection := database.Db.Collection("event_summaries")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	e.Status = "pending"
	res, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create event summary")
	}
	return e, nil
}

func (e *EventSummary) Update(id string) (*EventSummary, error) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "event_id", Value: e.EventID},
			{Key: "user_id", Value: e.UserID},
			{Key: "attendee_count", Value: e.AttendeeCount},
			{Key: "challenges", Value: e.Challenges},
			{Key: "successes", Value: e.Successes},
			{Key: "improvements", Value: e.Improvements},
			{Key: "status", Value: e.Status},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *EventSummary) Delete(id string) (*EventSummary, error) {
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{Key: "_id", Value: id}}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, errors.New("failed to delete event summary")
	}
	return e, nil
}

func (e *EventSummary) Approve(id string) (bool, error) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return res.ModifiedCount == 1, nil
}

func (e *EventSummary) Reject(id string) (bool, error) {
	collection := database.Db.Collection("event_summaries")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return res.ModifiedCount == 1, nil
}
