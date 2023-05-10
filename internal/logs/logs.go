package logs

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Log struct {
	ID            string `json:"id" bson:"_id"`
	UserID        string `json:"user_id" bson:"user_id"`
	DailyActivity string `json:"daily_activity" bson:"daily_activity"`
	Positives     string `json:"positives" bson:"positives"`
	Improvements  string `json:"improvements" bson:"improvements"`
	NextSteps     string `json:"next_steps" bson:"next_steps"`
	Status        string `json:"status" bson:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
}

func (l *Log) Create() (*Log, error) {
	collection := database.Db.Collection("logs")
	l.ID = uuid.New().String()
	l.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	l.UpdatedAt = l.CreatedAt
	l.Status = "pending"
	res, err := collection.InsertOne(context.TODO(), l)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create productivity log")
	}
	return l, nil
}

func (l *Log) Update(id string) (*Log, error) {
	collection := database.Db.Collection("logs")
	l.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.M{
			"daily_activity": l.DailyActivity,
			"positives":      l.Positives,
			"improvements":   l.Improvements,
			"next_steps":     l.NextSteps,
			"status":         l.Status,
			"updated_at":     l.UpdatedAt,
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Log) Approve(id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

func (l *Log) Reject(id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
