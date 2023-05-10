package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindContacts(filter bson.D) ([]*model.ContactOverview, error) {
	return nil, errors.New("not implemented")
}
func FindContactDetail(contact_id string) (*model.ContactDetail, error) {
	return nil, errors.New("not implemented")
}
func FindUserContacts(user_id string) ([]*model.ContactOverview, error) {
	return nil, errors.New("not implemented")
}

func CreateContact(new_contact model.NewContact, contact_creator string) (*model.ContactDetail, error) {
	collection := database.Db.Collection("contacts")
	filter := bson.D{{Key: "name", Value: new_contact.Name}, {Key: "email", Value: new_contact.Email}, {Key: "phone", Value: new_contact.Phone}}
	exists, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("contact already exists")
	}
	opts := options.FindOne().SetProjection(bson.D{{Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}, {Key: "_id", Value: 1}})
	var created_by model.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: contact_creator}}, opts).Decode(&created_by)
	if err != nil {
		return nil, err
	}
	contact := model.Contact{
		ID:        uuid.New().String(),
		Name:      new_contact.Name,
		Email:     new_contact.Email,
		Type:      new_contact.Type,
		Phone:     new_contact.Phone,
		CreatedBy: created_by.ID,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Active:    true,
		DeletedAt: bson.TypeNull.String(),
		UpdatedAt: bson.TypeNull.String(),
	}
	res, err := collection.InsertOne(context.TODO(), contact)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create contact")
	}
	return &model.ContactDetail{
		ID:     contact.ID,
		Name:   contact.Name,
		Type:   contact.Type,
		Email:  contact.Email,
		Phone:  contact.Phone,
		Notes:  contact.Notes,
		Active: true,
		CreatedBy: []*model.UserOverview{{
			ID:        contact.CreatedBy,
			FirstName: created_by.FirstName,
			LastName:  created_by.LastName,
		}},
		CreatedAt: contact.CreatedAt,
		UpdatedAt: bson.TypeNull.String(),
		DeletedAt: bson.TypeNull.String(),
	}, nil
}
func UpdateContact(update model.UpdateContact, filter bson.D) (*model.Contact, error) {
	collection := database.Db.Collection("contacts")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	update_fields := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: update.Name},
			{Key: "email", Value: update.Email},
			{Key: "phone", Value: update.Phone},
			{Key: "notes", Value: update.Notes},
			{Key: "type", Value: update.Type},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var c model.Contact
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_fields, &opts).Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func DeleteContact(filter bson.D) (bool, error) {
	collection := database.Db.Collection("contacts")
	deleted_at := time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "active", Value: false},
			{Key: "deleted_at", Value: deleted_at},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var c model.Contact
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&c)
	if err != nil {
		return false, err
	}
	return !c.Active, nil
}
