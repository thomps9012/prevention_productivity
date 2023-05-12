package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// queries
func FindGrantDetail(grant_id string) (*model.GrantDetail, error) {
	collection := database.Db.Collection("grants")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "created_by"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "created_by"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "created_by"}}
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: grant_id}}}}, user_stage, unwind}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	res := make([]*model.GrantDetail, 0)
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil, err
	}
	return res[0], nil
}
func FindAllGrants() ([]*model.GrantOverview, error) {
	projection := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}, {Key: "start_date", Value: 1}, {Key: "end_date", Value: 1}, {Key: "award_date", Value: 1}, {Key: "award_number", Value: 1}, {Key: "budget", Value: 1}, {Key: "active", Value: 1}})
	collection := database.Db.Collection("grants")
	cursor, err := collection.Find(context.TODO(), bson.D{{}}, projection)
	if err != nil {
		return nil, err
	}
	var res []*model.GrantOverview
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// mutations
func CreateGrant(new_grant model.NewGrant, grant_creator string) (*model.GrantDetail, error) {
	collection := database.Db.Collection("grants")
	var active bool
	if new_grant.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		active = true
	} else {
		active = false
	}
	grant := model.Grant{
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
	var grant_creator_info model.UserOverview
	user_projection := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: grant_creator}}, user_projection).Decode(&grant_creator_info)
	if err != nil {
		return nil, err
	}
	return &model.GrantDetail{
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
		CreatedBy:   &grant_creator_info,
		CreatedAt:   grant.CreatedAt,
		UpdatedAt:   grant.UpdatedAt,
	}, nil
}
func UpdateGrant(update model.UpdateGrant) (*model.Grant, error) {
	collection := database.Db.Collection("grants")
	updated_at := time.Now().Format("01-02-2006 15:04:05")
	filter := bson.D{{Key: "_id", Value: update.ID}}
	var active bool
	if update.EndDate > time.Now().Format("01-02-2006 15:04:05") {
		active = true
	} else {
		active = false
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
			{Key: "active", Value: active},
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
	var g model.Grant
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
