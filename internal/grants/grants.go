package grants

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
	"github.com/google/uuid"
)

type Grant struct {
	ID	   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	StartDate string `json:"start_date" bson:"start_date"`
	AwardDate string `json:"award_date" bson:"award_date"`
	EndDate string `json:"end_date" bson:"end_date"`
	AwardNumber string `json:"award_number" bson:"award_number"`
	Budget float64 `json:"budget" bson:"budget"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	IsActive bool `json:"is_active"`
}

func (g *Grant) Create() {
	collection := database.Db.Collection("grants")
	g.ID = uuid.New().String()
	if (g.EndDate > time.Now().Format("01-02-2006 15:04:05")) {
		g.IsActive = true
	} else {
		g.IsActive = false
	}
	g.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
	g.UpdatedAt = g.CreatedAt
	_, err := collection.InsertOne(context.TODO(), g)
	if err != nil {
		panic(err)
	}
}

func (g *Grant) Update(id string) {
	collection := database.Db.Collection("grants")
	g.UpdatedAt = time.Now().Format("01-02-2006 15:04:05")
	filter := bson.D{{"_id", id}}
	if (g.EndDate > time.Now().Format("01-02-2006 15:04:05")) {
		g.IsActive = true
	} else {
		g.IsActive = false
	}
	update := bson.D{
		{"$set", bson.D{
			{"name", g.Name},
			{"description", g.Description},
			{"start_date", g.StartDate},
			{"award_date", g.AwardDate},
			{"end_date", g.EndDate},
			{"award_number", g.AwardNumber},
			{"budget", g.Budget},
			{"updated_at", g.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No grant found with that ID")
	}
}

func (g *Grant) Delete(id string) {
	collection := database.Db.Collection("grants")
	filter := bson.D{{"_id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}