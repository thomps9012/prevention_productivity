package logs

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"strings"
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID	   string `json:"id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	FocusArea string `json:"focus_area" bson:"focus_area"`
	Actions string `json:"actions" bson:"actions"`
	Successes string `json:"successes" bson:"successes"`
	Improvements string `json:"improvements" bson:"improvements"`
	NextSteps string `json:"next_steps bson:"next_steps"`
	Status string `json:"status" bson:"status"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}


func (l *Log) Create() {
	collection := database.Db.Collection("logs")
	l.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	l.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	l.UpdatedAt = l.CreatedAt
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
	println(l.UpdatedAt)
	update := bson.D{
		{"$set", bson.D{
			{"focus_area", l.FocusArea},
			{"actions", l.Actions},
			{"successes", l.Successes},
			{"improvements", l.Improvements},
			{"next_steps", l.NextSteps},
			{"status", l.Status},
			{"updated_at", l.UpdatedAt},
		}},
	}
	println(update)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	println(result)
	if result.ModifiedCount == 0 {
		panic("Log not found")
	}
}

func (l *Log) Approve(id string) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "approved"},
			{"updated_at", time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Log not found")
	}
}

func (l *Log) Reject(id string) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "rejected"},
			{"updated_at", time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Log not found")
	}
}