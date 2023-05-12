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
func EventSummaryDetail(filter bson.D) (*model.EventSummaryWithNotes, error) {
	event_summary := make([]*model.EventSummaryWithNotes, 0)
	collection := database.Db.Collection("event_summaries")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "summary_author"}}}}
	event_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "events"}, {Key: "localField", Value: "event_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "event_description"}}}}
	co_planner_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "co_planners"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "co_planners"}}}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "author"}}}},
		bson.D{{Key: "$unwind", Value: "$author"}},
	}}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$summary_author"}}
	unwind_event := bson.D{{Key: "$unwind", Value: "$event_description"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, co_planner_stage, event_stage, unwind_author, unwind_event}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &event_summary)
	if err != nil {
		return nil, err
	}
	if len(event_summary) == 0 {
		return nil, errors.New("you're attempting to view an event summary that either doesn't exist, or you didn't create")
	}
	return event_summary[0], nil
}
func FindEventSummaries(filter bson.D) ([]*model.EventSummaryOverview, error) {
	event_summaries := make([]*model.EventSummaryOverview, 0)
	collection := database.Db.Collection("event_summaries")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "summary_author"}}}}
	event_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "events"}, {Key: "localField", Value: "event_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "event_description"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$summary_author"}}
	unwind_event := bson.D{{Key: "$unwind", Value: "$event_description"}}
	pipeline := mongo.Pipeline{filter, note_stage, note_count, user_stage, event_stage, unwind_author, unwind_event}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &event_summaries)
	if err != nil {
		return nil, err
	}
	return event_summaries, nil
}
func FindUserEventSummaries(user_id string) ([]*model.EventSummaryOverview, error) {
	event_summaries := make([]*model.EventSummaryOverview, 0)
	collection := database.Db.Collection("event_summaries")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "summary_author"}}}}
	event_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "events"}, {Key: "localField", Value: "event_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "event_description"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$summary_author"}}
	unwind_event := bson.D{{Key: "$unwind", Value: "$event_description"}}
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}, note_stage, note_count, user_stage, event_stage, unwind_author, unwind_event}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &event_summaries)
	if err != nil {
		return nil, err
	}
	return event_summaries, nil
}

// mutations
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
	event_projection := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "title", Value: 1}, {Key: "start_date", Value: 1}})
	author_projection := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})
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
	var summary model.EventSummary
	err := collection.FindOne(context.TODO(), filter).Decode(&summary)
	if err != nil {
		return nil, errors.New("you're attempting to update an event summary that either doesn't exist, or you didn't create")
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "event_id", Value: update.EventID},
			{Key: "attendee_count", Value: update.AttendeeCount},
			{Key: "challenges", Value: update.Challenges},
			{Key: "successes", Value: update.Successes},
			{Key: "improvements", Value: update.Improvements},
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
	err = collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&e)
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
