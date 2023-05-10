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

func FindSchoolReportPlanDetail(debrief_id string) (*models.SchoolReportPlanWithNotes, error) {
	return nil, errors.New("method unimplemented")
}
func FindSchoolReportPlans(filter bson.D) ([]*models.SchoolReportPlanOverview, error) {
	return nil, errors.New("method unimplemented")
}
func FindUserSchoolReportPlans(user_id string) ([]*models.SchoolReportPlanOverview, error) {
	return nil, errors.New("method unimplemented")
}

func CreateSchoolReportPlan(new_plan models.NewSchoolReportPlan, plan_creator string) (*models.SchoolReportPlanRes, error) {
	collection := database.Db.Collection("school_report_plans")
	plan := models.SchoolReportPlan{
		ID:              uuid.New().String(),
		UserID:          &plan_creator,
		Date:            new_plan.Date,
		Co_Facilitators: new_plan.CoFacilitators,
		Curriculum:      new_plan.Curriculum,
		School:          new_plan.School,
		LessonTopics:    new_plan.LessonTopics,
		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:       bson.TypeNull.String(),
		Status:          "pending",
	}
	res, err := collection.InsertOne(context.TODO(), plan)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert plan")
	}
	var plan_author models.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{"_id", plan_creator}}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})).Decode(&plan_author)
	if err != nil {
		return nil, err
	}
	return &models.SchoolReportPlanRes{
		ID:         plan.ID,
		PlanAuthor: &plan_author,
		Date:       plan.Date,
		School:     plan.School,
		Status:     plan.Status,
		CreatedAt:  plan.CreatedAt,
	}, nil
}
func UpdateSchoolReportPlan(update models.UpdateSchoolReportPlan) (*models.SchoolReportPlan, error) {
	collection := database.Db.Collection("school_report_plans")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: update.ID}}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "curriculum", Value: update.Curriculum},
			{Key: "school", Value: update.School},
			{Key: "lesson_topics", Value: update.LessonTopics},
			{Key: "co_facilitators", Value: update.CoFacilitators},
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
	var srp models.SchoolReportPlan
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&srp)
	if err != nil {
		return nil, err
	}
	return &srp, nil
}
func DeleteSchoolReportPlan(plan_id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: plan_id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}
func ApproveSchoolReportPlan(plan_id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: plan_id}}
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
func RejectSchoolReportPlan(plan_id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: plan_id}}
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
