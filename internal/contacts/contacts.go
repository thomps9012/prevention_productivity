package contacts

import (
	"context"
	"fmt"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Contact struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Type      string `json:"type" bson:"type"`
	Notes     string `json:"notes" bson:"notes"`
	Active    bool
	CreatedBy string `json:"created_by" bson:"created_by"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}

func (c *Contact) Create() {
	collection := database.Db.Collection("contacts")
	var contact Contact
	filter := bson.D{{"name", c.Name}, {"email", c.Email}, {"phone", c.Phone}}
	err := collection.FindOne(context.TODO(), filter).Decode(&contact)
	if err != nil{
		c.ID = uuid.New().String()
		c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		c.UpdatedAt = c.CreatedAt
		c.Active = true
		_, err := collection.InsertOne(context.TODO(), c)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Errorf("contact already exists")
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
			{"type", c.Type},
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
	c.Active = false
	c.DeletedAt = time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{
		{"$set", bson.D{
			{"active", c.Active},
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
