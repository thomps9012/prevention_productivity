package logs

import (
	database "prevention_productivity/base/internal/db"
	"context"
	"strings"
	"github.com/google/uuid"
	"time"
)

type Log struct {
	ID	   string `json:"id" bson:"_id"`
	UserID string `json:"user_id"`
	Action string `json:"action"`
	CreatedAt string `json:"created_at"`
}


func (l *Log) Create() {
	collection := database.Db.Collection("logs")
	l.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	l.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := collection.InsertOne(context.TODO(), l)
	if err != nil {
		panic(err)
	}
}