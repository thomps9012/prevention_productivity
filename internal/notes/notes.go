package notes

import (
	"time"
	"context"
	"strings"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	database "thomps9012/prevention_productivity/internal/db"
)

type Note struct {
	ID	   string `json:"id" bson:"_id"`
	ItemID string `json:"item_id" bson:"item_id"`
	UserID string `json:"user_id" bson:"user_id"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

func (n *Note) Create() {
	collection := database.Db.Collection("notes")
	n.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	n.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := collection.InsertOne(context.TODO(), n)
	if err != nil {
		panic(err)
	}
}

func (n *Note) Update(id string) {
	collection := database.Db.Collection("notes")
	n.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"title", n.Title},
		{"content", n.Content},
		{"updated_at", n.UpdatedAt},
	}}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Note not found")
	}
}