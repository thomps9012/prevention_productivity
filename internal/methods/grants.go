package methods

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"thomps9012/prevention_productivity/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindGrantDetail(grant_id string) (*models.GrantDetail, error) {
	return nil, errors.New("method unimplemented")
}
func FindAllGrants() ([]*models.GrantOverview, error) {
	return nil, errors.New("method unimplemented")
}

func CreateGrant(new_grant models.NewGrant, grant_creator string) (*models.GrantDetail, error) {
	collection := database.Db.Collection("grants")
	var active bool
	if new_grant.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		active = true
	} else {
		active = false
	}
	grant := models.Grant{
		ID:          uuid.New().String(),
		Name:        new_grant.Name,
		Description: new_grant.Description,
		Goals:       new_grant.Goals,
		Objectives:  new_grant.Objectives,
		StartDate:   new_grant.StartDate,
		EndDate:     new_grant.EndDate,
		AwardNumber: new_grant.AwardNumber,
		AwardDate:   new_grant.AwardDate,
		Budget:      new_grant.Budget,
		CreatedBy:   grant_creator,
		CreatedAt:   time.Now().Format("01-02-2006 15:04:05"),
		UpdatedAt:   bson.TypeNull.String(),
		Active:      active,
	}
	res, err := collection.InsertOne(context.TODO(), grant)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create grant")
	}
	var grant_creator_info models.UserOverview
	user_projection := options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{"_id", grant_creator}}, user_projection).Decode(&grant_creator_info)
	if err != nil {
		return nil, err
	}
	return &models.GrantDetail{
		ID:          grant.ID,
		Name:        grant.Name,
		Description: grant.Description,
		Goals:       grant.Goals,
		Objectives:  grant.Objectives,
		StartDate:   grant.StartDate,
		AwardDate:   grant.AwardDate,
		EndDate:     grant.EndDate,
		AwardNumber: grant.AwardNumber,
		Budget:      grant.Budget,
		Active:      grant.Active,
		CreatedBy:   []*models.UserOverview{&grant_creator_info},
		CreatedAt:   grant.CreatedAt,
		UpdatedAt:   grant.UpdatedAt,
	}, nil
}
func UpdateGrant(update models.UpdateGrant) (*models.Grant, error) {
	collection := database.Db.Collection("grants")
	updated_at := time.Now().Format("01-02-2006 15:04:05")
	filter := bson.D{{Key: "_id", Value: update.ID}}
	if update.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		update.Active = true
	} else {
		update.Active = false
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: update.Name},
			{Key: "description", Value: update.Description},
			{Key: "goals", Value: update.Goals},
			{Key: "objectives", Value: update.Objectives},
			{Key: "start_date", Value: update.StartDate},
			{Key: "award_date", Value: update.AwardDate},
			{Key: "end_date", Value: update.EndDate},
			{Key: "award_number", Value: update.AwardNumber},
			{Key: "active", Value: update.Active},
			{Key: "budget", Value: update.Budget},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	var g models.Grant
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
func DeleteGrant(grant_id string) (bool, error) {
	collection := database.Db.Collection("grants")
	filter := bson.D{{Key: "_id", Value: grant_id}}
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}
