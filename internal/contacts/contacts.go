package contacts

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"strings"
	"time"
	"github.com/google/uuid"
	"fmt"
)

type Contact struct {
	ID	   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
	Notes string `json:"notes" bson:"notes"`
	IsActive bool `json:"is_active"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}

func (c *Contact) Create() {
	collection := database.Db.Collection("contacts")
	c.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdatedAt = c.CreatedAt
	c.IsActive = true
	_, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		panic(err)
	}
}

func (c *Contact) Update(id string) {
	collection := database.Db.Collection("contacts")
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"name", c.Name},
			{"email", c.Email},
			{"phone", c.Phone},
			{"notes", c.Notes},
			{"updated_at", c.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Matched Count: ", result.MatchedCount)
	fmt.Println("Modified Count: ", result.ModifiedCount)
	if result.ModifiedCount == 0 {
		fmt.Println("No document found")
	}
}

func (c *Contact) Delete(id string) {
	collection := database.Db.Collection("contacts")
	filter := bson.D{{"_id", id}}
	c.IsActive = false
	c.DeletedAt = time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{
		{"$set", bson.D{
			{"is_active", c.IsActive},
			{"deleted_at", c.DeletedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Matched Count: ", result.MatchedCount)
	fmt.Println("Modified Count: ", result.ModifiedCount)
	if result.ModifiedCount == 0 {
		fmt.Println("No document found")
	}
}

	