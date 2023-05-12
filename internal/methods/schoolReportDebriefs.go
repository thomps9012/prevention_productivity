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
func FindSchoolReportDebriefDetail(filter bson.D) (*model.SchoolReportDebriefWithNotes, error) {
	debrief_detail := make([]*model.SchoolReportDebriefWithNotes, 0)
	collection := database.Db.Collection("school_report_debriefs")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "debrief_author"}}}}
	plan_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "school_report_plans"}, {Key: "localField", Value: "lesson_plan_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "lesson_plan"}}}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$debrief_author"}}
	unwind_lesson_plan := bson.D{{Key: "$unwind", Value: "$lesson_plan"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, plan_stage, unwind_author, unwind_lesson_plan}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &debrief_detail)
	if err != nil {
		return nil, err
	}
	if len(debrief_detail) == 0 {
		return nil, errors.New("you're attempting to view a debrief that either doesn't exist of you didn't create")
	}
	return debrief_detail[0], nil
}
func FindSchoolReportDebriefs(filter bson.D) ([]*model.SchoolReportDebriefOverview, error) {
	debriefs := make([]*model.SchoolReportDebriefOverview, 0)
	collection := database.Db.Collection("school_report_debriefs")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "debrief_author"}}}}
	plan_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "school_report_plans"}, {Key: "localField", Value: "lesson_plan_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "lesson_plan"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$debrief_author"}}
	unwind_lesson_plan := bson.D{{Key: "$unwind", Value: "$lesson_plan"}}
	pipeline := mongo.Pipeline{filter, note_stage, note_count, user_stage, plan_stage, unwind_author, unwind_lesson_plan}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &debriefs)
	if err != nil {
		return nil, err
	}
	return debriefs, nil
}
func FindUserSchoolReportDebriefs(user_id string) ([]*model.SchoolReportDebriefOverview, error) {
	debriefs := make([]*model.SchoolReportDebriefOverview, 0)
	collection := database.Db.Collection("school_report_debriefs")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "debrief_author"}}}}
	plan_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "school_report_plans"}, {Key: "localField", Value: "lesson_plan_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "lesson_plan"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind_author := bson.D{{Key: "$unwind", Value: "$debrief_author"}}
	unwind_lesson_plan := bson.D{{Key: "$unwind", Value: "$lesson_plan"}}
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}, note_stage, note_count, user_stage, plan_stage, unwind_author, unwind_lesson_plan}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &debriefs)
	if err != nil {
		return nil, err
	}
	return debriefs, nil
}

// mutations
func CreateSchoolReportDebrief(new_debrief model.NewSchoolReportDebrief, debrief_author string) (*model.SchoolReportDebriefRes, error) {
	collection := database.Db.Collection("school_report_debriefs")
	debrief := model.SchoolReportDebrief{
		ID:                     uuid.New().String(),
		UserID:                 debrief_author,
		LessonPlanID:           new_debrief.LessonPlanID,
		StudentCount:           new_debrief.StudentCount,
		StudentList:            new_debrief.StudentList,
		ChallengesImprovements: new_debrief.ChallengesImprovements,
		Positives:              new_debrief.Positives,
		Discussion:             new_debrief.Discussion,
		CreatedAt:              time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:              bson.TypeNull.String(),
		Status:                 "pending",
	}
	res, err := collection.InsertOne(context.TODO(), debrief)
	if err != nil {
		return nil, err
	}
	var debrief_author_info model.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: debrief_author}}, options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})).Decode(&debrief_author_info)
	if err != nil {
		return nil, err
	}
	var lesson_plan_info model.SchoolReportPlanDescription
	err = database.Db.Collection("school_report_plans").FindOne(context.TODO(), bson.D{{Key: "_id", Value: debrief.LessonPlanID}}, options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "school", Value: 1}, {Key: "date", Value: 1}})).Decode(&lesson_plan_info)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert debrief")
	}
	return &model.SchoolReportDebriefRes{
		ID:            debrief.ID,
		DebriefAuthor: &debrief_author_info,
		LessonPlan:    &lesson_plan_info,
		Status:        debrief.Status,
		CreatedAt:     debrief.CreatedAt,
	}, nil
}
func UpdateSchoolReportDebrief(update model.UpdateSchoolReportDebrief, filter bson.D) (*model.SchoolReportDebrief, error) {
	collection := database.Db.Collection("school_report_debriefs")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	var debrief model.SchoolReportDebrief
	err := collection.FindOne(context.TODO(), filter).Decode(&debrief)
	if err != nil {
		return nil, errors.New("you're attempting to update a debrief that either doesn't exist, or you didn't create")
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "student_count", Value: update.StudentCount},
			{Key: "student_list", Value: update.StudentList},
			{Key: "discussion", Value: update.Discussion},
			{Key: "positives", Value: update.Positives},
			{Key: "challenges_improvements", Value: update.ChallengesImprovements},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var srd model.SchoolReportDebrief
	err = collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&srd)
	if err != nil {
		return nil, err
	}
	return &srd, nil
}
func DeleteDebrief(filter bson.D) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}
func ApproveDebrief(debrief_id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: debrief_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
func RejectDebrief(debrief_id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: debrief_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
