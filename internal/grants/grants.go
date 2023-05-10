package grants

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Grant struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Goals       []*string `json:"goals" bson:"goals"`
	Objectives  []*string `json:"objectives" bson:"objectives"`
	StartDate   string    `json:"start_date" bson:"start_date"`
	AwardDate   string    `json:"award_date" bson:"award_date"`
	EndDate     string    `json:"end_date" bson:"end_date"`
	AwardNumber string    `json:"award_number" bson:"award_number"`
	Budget      float64   `json:"budget" bson:"budget"`
	CreatedBy   *string   `json:"created_by" bson:"created_by"`
	CreatedAt   string    `json:"created_at" bson:"created_at"`
	UpdatedAt   string    `json:"updated_at" bson:"updated_at"`
	Active      bool
}

func (g *Grant) Create() (*Grant, error) {
	collection := database.Db.Collection("grants")
	g.ID = uuid.New().String()
	if g.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		g.Active = true
	} else {
		g.Active = false
	}
	g.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
	g.UpdatedAt = g.CreatedAt
	res, err := collection.InsertOne(context.TODO(), g)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create grant")
	}
	return g, nil
}

func (g *Grant) Update(id string) (*Grant, error) {
	collection := database.Db.Collection("grants")
	g.UpdatedAt = time.Now().Format("01-02-2006 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	if g.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		g.Active = true
	} else {
		g.Active = false
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: g.Name},
			{Key: "description", Value: g.Description},
			{Key: "goals", Value: g.Goals},
			{Key: "objectives", Value: g.Objectives},
			{Key: "start_date", Value: g.StartDate},
			{Key: "award_date", Value: g.AwardDate},
			{Key: "end_date", Value: g.EndDate},
			{Key: "award_number", Value: g.AwardNumber},
			{Key: "budget", Value: g.Budget},
			{Key: "updated_at", Value: g.UpdatedAt},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Grant) Delete(id string) (*Grant, error) {
	collection := database.Db.Collection("grants")
	filter := bson.D{{Key: "_id", Value: id}}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, errors.New("failed to delete grant")
	}
	return g, nil
}
