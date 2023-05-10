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

func FindLogDetail(log_id string) (*model.LogWithNotes, error) {
	// IsAdmin := auth.ForAdmin(ctx)
	// UserID := auth.ForUserID(ctx)
	// var LogWithNotes *model.LogWithNotes
	// logCollection := database.Db.Collection("logs")
	// logFilter := bson.D{{Key: "_id", Value: id}}
	// user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "log_author"}, {
	// 	// add unwinding
	// 	Key: "pipeline", Value: bson.A{
	// 		bson.D{{Key: "$project", Value: bson.D{{Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}, {Key: "_id", Value: 1}}}},
	// 	},
	// }}}}
	// notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	// // add projection
	// pipeline := mongo.Pipeline{logFilter, notes_stage, user_stage}
	// cursor, err := logCollection.Aggregate(context.TODO(), pipeline)
	// if err != nil {
	// 	return nil, err
	// }
	// err = cursor.All(context.TODO(), &LogWithNotes)
	// if err != nil {
	// 	return nil, err
	// }
	// if *LogWithNotes.Log.UserID != UserID && !IsAdmin {
	// 	return nil, fmt.Errorf("Unauthorized")
	// }
	// return LogWithNotes, nil
	return nil, errors.New("method unimplemented")
}
func FindAllLogs(filter bson.D) ([]*model.LogOverview, error) {
	// IsAdmin := auth.ForAdmin(ctx)
	// UserID := auth.ForUserID(ctx)
	// if UserID == "" {
	// 	return nil, fmt.Errorf("Unauthorized")
	// }
	// var filter bson.D
	// if IsAdmin {
	// 	filter = bson.D{{}}
	// } else {
	// 	filter = bson.D{{Key: "user_id", Value: UserID}}
	// }
	// logsCollection := database.Db.Collection("logs")
	// sort_stage := bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}
	// user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "log_author"}, {
	// 	Key: "pipeline", Value: bson.A{
	// 		bson.D{{Key: "$project", Value: bson.D{{Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}, {Key: "_id", Value: 1}}}},
	// 	},
	// }}}}
	// note_stage := bson.D{{Key: "$count", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "note_count"}}}}
	// // add in projection
	// log_pipeline := mongo.Pipeline{filter, sort_stage, user_stage, note_stage}
	// cursor, err := logsCollection.Aggregate(context.TODO(), log_pipeline)
	// if err != nil {
	// 	return nil, err
	// }
	// var allLogs []*model.AllLogs
	// err = cursor.All(context.TODO(), &allLogs)
	// if err != nil {
	// 	return nil, err
	// }
	// return allLogs, nil
	return nil, errors.New("method unimplemented")
}
func FindUserLogs(user_id string) ([]*model.LogOverview, error) {
	return nil, errors.New("method unimplemented")
}

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
	err := collection.FindOne(context.Background(), bson.M{"_id": log_author}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})).Decode(&author_info)
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
	update_args := bson.D{
		{Key: "$set", Value: bson.M{
			"daily_activity": update.DailyActivity,
			"positives":      update.Positives,
			"improvements":   update.Improvements,
			"next_steps":     update.NextSteps,
			"status":         update.Status,
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
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&l)
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
