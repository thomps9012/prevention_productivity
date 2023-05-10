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

func EventSummaryDetail(summary_id string) (*model.EventSummaryWithNotes, error) {
	return nil, errors.New("method unimplemented")
}
func FindEventSummaries(filter bson.D) ([]*model.EventSummaryOverview, error) {
	return nil, errors.New("method unimplemented")
}
func FindUserEventSummaries(user_id string) ([]*model.EventSummaryOverview, error) {
	return nil, errors.New("method unimplemented")
}

func CreateEventSummary(new_summary model.NewEventSummary, summary_creator string) (*model.EventSummaryRes, error) {
	collection := database.Db.Collection("event_summaries")
	summary := model.EventSummary{
		ID:            uuid.New().String(),
		UserID:        summary_creator,
		EventID:       new_summary.EventID,
		CoPlanners:    new_summary.CoPlanners,
		AttendeeCount: new_summary.AttendeeCount,
		Challenges:    new_summary.Challenges,
		Successes:     new_summary.Successes,
		Improvements:  new_summary.Improvements,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     bson.TypeNull.String(),
		Status:        "pending",
	}
	var event_description model.EventDescription
	var summary_author model.UserOverview
	event_projection := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {"title", 1}, {"start_date", 1}})
	author_projection := options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})
	res, err := collection.InsertOne(context.TODO(), summary)
	if err != nil {
		return nil, err
	}
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.M{"_id": summary_creator}, author_projection).Decode(&summary_author)
	if err != nil {
		return nil, err
	}
	err = database.Db.Collection("events").FindOne(context.TODO(), bson.M{"_id": new_summary.EventID}, event_projection).Decode(&event_description)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create event summary")
	}
	return &model.EventSummaryRes{
		ID:            summary.ID,
		Event:         &event_description,
		SummaryAuthor: &summary_author,
		Status:        summary.Status,
		CreatedAt:     summary.CreatedAt,
	}, nil
}
func UpdateEventSummary(update model.UpdateEventSummary, filter bson.D) (*model.EventSummary, error) {
	collection := database.Db.Collection("event_summaries")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "event_id", Value: update.EventID},
			{Key: "attendee_count", Value: update.AttendeeCount},
			{Key: "challenges", Value: update.Challenges},
			{Key: "successes", Value: update.Successes},
			{Key: "improvements", Value: update.Improvements},
			{Key: "status", Value: update.Status},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var e model.EventSummary
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func DeleteEventSummary(filter bson.D) (bool, error) {
	collection := database.Db.Collection("event_summaries")
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}
func ApproveEventSummary(summary_id string) (bool, error) {
	collection := database.Db.Collection("event_summaries")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: summary_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return res.ModifiedCount == 1, nil
}
func RejectEventSummary(summary_id string) (bool, error) {
	collection := database.Db.Collection("event_summaries")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: summary_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return res.ModifiedCount == 1, nil
}
