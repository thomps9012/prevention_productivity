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

func FindSchoolReportDebriefDetail(debrief_id string) (*models.SchoolReportDebriefWithNotes, error) {
	return nil, errors.New("method unimplemented")
}
func FindSchoolReportDebriefs(filter bson.D) ([]*models.SchoolReportDebriefOverview, error) {
	return nil, errors.New("method unimplemented")
}
func FindUserSchoolReportDebriefs(user_id string) ([]*models.SchoolReportDebriefOverview, error) {
	return nil, errors.New("method unimplemented")
}

func CreateSchoolReportDebrief(new_debrief models.NewSchoolReportDebrief, debrief_author string) (*models.SchoolReportDebriefRes, error) {
	collection := database.Db.Collection("school_report_debriefs")
	debrief := models.SchoolReportDebrief{
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
	var debrief_author_info models.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{"_id", debrief_author}}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})).Decode(&debrief_author_info)
	if err != nil {
		return nil, err
	}
	var lesson_plan_info models.SchoolReportPlanDescription
	err = database.Db.Collection("lesson_plans").FindOne(context.TODO(), bson.D{{"_id", debrief.LessonPlanID}}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"school", 1}, {"date", 1}})).Decode(&lesson_plan_info)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert debrief")
	}
	return &models.SchoolReportDebriefRes{
		ID:            debrief.ID,
		DebriefAuthor: &debrief_author_info,
		LessonPlan:    &lesson_plan_info,
		Status:        debrief.Status,
		CreatedAt:     debrief.CreatedAt,
	}, nil
}
func UpdateSchoolReportDebrief(update models.UpdateSchoolReportDebrief) (*models.SchoolReportDebrief, error) {
	collection := database.Db.Collection("school_report_debriefs")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: update.ID}}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "student_count", Value: update.StudentCount},
			{Key: "student_list", Value: update.StudentList},
			{Key: "discussion", Value: update.Discussion},
			{Key: "positives", Value: update.Positives},
			{Key: "challenges_improvements", Value: update.ChallengesImprovements},
			{Key: "status", Value: update.Status},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var srd models.SchoolReportDebrief
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&srd)
	if err != nil {
		return nil, err
	}
	return &srd, nil
}
func DeleteDebrief(debrief_id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "_id", Value: debrief_id}}
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