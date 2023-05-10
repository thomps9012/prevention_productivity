package notes

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Note struct {
	ID        string `json:"id" bson:"_id"`
	ItemID    string `json:"item_id" bson:"item_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	Title     string `json:"title" bson:"title"`
	Content   string `json:"content" bson:"content"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

func (n *Note) Create() (*Note, error) {
	collection := database.Db.Collection("notes")
	n.ID = uuid.New().String()
	n.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	n.UpdatedAt = n.CreatedAt
	res, err := collection.InsertOne(context.TODO(), n)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert note")
	}
	return n, nil
}

func (n *Note) Update() (*Note, error) {
	collection := database.Db.Collection("notes")
	n.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: n.ID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: n.Title},
			{Key: "content", Value: n.Content},
			{Key: "updated_at", Value: n.UpdatedAt},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (n *Note) Remove(id string) (*Note, error) {
	collection := database.Db.Collection("notes")
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, errors.New("no note with that id found")
	}
	return n, nil
}
