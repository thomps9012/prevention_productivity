package logs

import (
	database "prevention_productivity/base/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"strings"
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID	   string `json:"id" bson:"_id"`
	UserID string `json:"user_id"`
	FocusArea string `json:"focus_area"`
	Actions string `json:"actions"`
	Successes string `json:"successes"`
	Improvements string `json:"improvements"`
	NextSteps string `json:"next_steps"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


func (l *Log) Create() {
	collection := database.Db.Collection("logs")
	l.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	l.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	l.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), l)
	if err != nil {
		panic(err)
	}
}

func (l *Log) Update(id string) {
	collection := database.Db.Collection("logs")
	l.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"focus_area", l.FocusArea},
		{"actions", l.Actions},
		{"successes", l.Successes},
		{"improvements", l.Improvements},
		{"next_steps", l.NextSteps},
		{"status", l.Status},
		{"updated_at", l.UpdatedAt},
	}}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Log not found")
	}
}