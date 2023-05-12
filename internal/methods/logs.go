package methods

import (
	"context"
	"errors"
	"fmt"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// queries
func FindLogDetail(filter bson.D) (*model.LogWithNotes, error) {
	LogWithNotes := make([]*model.LogWithNotes, 0)
	logCollection := database.Db.Collection("logs")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "log_author"}}}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$log_author"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, unwind}
	cursor, err := logCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &LogWithNotes)
	if err != nil {
		return nil, err
	}
	if len(LogWithNotes) == 0 {
		return nil, errors.New("you're attempting to view a log that either doesn't exist, or you didn't create")
	}
	return LogWithNotes[0], nil
}
func FindAllLogs(filter bson.D) ([]*model.LogOverview, error) {
	logsCollection := database.Db.Collection("logs")
	sort_stage := bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "log_author"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$log_author"}}
	projection := bson.D{{Key: "$project", Value: bson.D{{Key: "notes", Value: 0}}}}
	log_pipeline := mongo.Pipeline{filter, sort_stage, user_stage, note_stage, unwind, note_count, projection}
	cursor, err := logsCollection.Aggregate(context.TODO(), log_pipeline)
	if err != nil {
		return nil, err
	}
	var allLogs []*model.LogOverview
	err = cursor.All(context.TODO(), &allLogs)
	if err != nil {
		return nil, err
	}
	return allLogs, nil
}
func FindUserLogs(user_id string) ([]*model.LogOverview, error) {
	logsCollection := database.Db.Collection("logs")
	sort_stage := bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "log_author"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$log_author"}}
	log_pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}, sort_stage, user_stage, note_stage, note_count, unwind}
	cursor, err := logsCollection.Aggregate(context.TODO(), log_pipeline)
	if err != nil {
		return nil, err
	}
	var userLogs []*model.LogOverview
	err = cursor.All(context.TODO(), &userLogs)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(userLogs))
	return userLogs, nil
}

// mutations
func CreateNewLog(new_log model.NewLog, log_author string) (*model.LogRes, error) {
	collection := database.Db.Collection("logs")
	log := model.Log{
		ID:            uuid.New().String(),
		UserID:        log_author,
		DailyActivity: new_log.DailyActivity,
		Positives:     new_log.Positives,
		Improvements:  new_log.Improvements,
		NextSteps:     new_log.NextSteps,
		Status:        "pending",
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     bson.TypeNull.String(),
	}
	var author_info model.UserOverview
	err := database.Db.Collection("users").FindOne(context.Background(), bson.M{"_id": log_author}, options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})).Decode(&author_info)
	if err != nil {
		return nil, err
	}
	res, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create productivity log")
	}
	return &model.LogRes{
		ID:        log.ID,
		LogAuthor: &author_info,
		Status:    log.Status,
		CreatedAt: log.CreatedAt,
	}, nil
}
func UpdateLog(update model.UpdateLog, filter bson.D) (*model.Log, error) {
	collection := database.Db.Collection("logs")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	var log model.Log
	err := collection.FindOne(context.TODO(), filter).Decode(&log)
	if err != nil {
		return nil, errors.New("you're attempting to update a log that either doesn't exist or you didn't create")
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.M{
			"daily_activity": update.DailyActivity,
			"positives":      update.Positives,
			"improvements":   update.Improvements,
			"next_steps":     update.NextSteps,
			"updated_at":     updated_at,
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var l model.Log
	err = collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
func DeleteLog(filter bson.D) (bool, error) {
	collection := database.Db.Collection("logs")
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}
func ApproveLog(log_id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: log_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
func RejectLog(log_id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: log_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
