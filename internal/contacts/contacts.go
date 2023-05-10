package contacts

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contact struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
	Type      string `json:"type" bson:"type"`
	Notes     string `json:"notes" bson:"notes"`
	Active    bool   `json:"active" bson:"active"`
	CreatedBy string `json:"created_by" bson:"created_by"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}

func (c *Contact) Create() (*Contact, error) {
	collection := database.Db.Collection("contacts")
	filter := bson.D{{Key: "name", Value: c.Name}, {Key: "email", Value: c.Email}, {Key: "phone", Value: c.Phone}}
	exists, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("contact already exists")
	}
	c.ID = uuid.New().String()
	c.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	c.UpdatedAt = c.CreatedAt
	c.Active = true
	res, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create contact")
	}
	return c, nil
}

func (c *Contact) Update(id string) (*Contact, error) {
	collection := database.Db.Collection("contacts")
	c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: c.Name},
			{Key: "email", Value: c.Email},
			{Key: "phone", Value: c.Phone},
			{Key: "notes", Value: c.Notes},
			{Key: "type", Value: c.Type},
			{Key: "updated_at", Value: c.UpdatedAt},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Contact) Delete(id string) (*Contact, error) {
	collection := database.Db.Collection("contacts")
	filter := bson.D{{Key: "_id", Value: id}}
	c.Active = false
	c.DeletedAt = time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "active", Value: c.Active},
			{Key: "deleted_at", Value: c.DeletedAt},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
