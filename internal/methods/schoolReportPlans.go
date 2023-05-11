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

func FindSchoolReportPlanDetail(filter bson.D) (*model.SchoolReportPlanWithNotes, error) {
	plan_detail := make([]*model.SchoolReportPlanWithNotes, 0)
	collection := database.Db.Collection("school_report_plans")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "plan_author"}}}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, unwind}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &plan_detail)
	if err != nil {
		return nil, err
	}
	return plan_detail[0], nil
}
func FindSchoolReportPlans(filter bson.D) ([]*model.SchoolReportPlanOverview, error) {
	plans := make([]*model.SchoolReportPlanOverview, 0)
	collection := database.Db.Collection("school_report_plans")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "plan_author"}}}}
	notes_stage := bson.D{{Key: "$count", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "note_count"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, unwind}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &plans)
	if err != nil {
		return nil, err
	}
	return plans, nil
}
func FindUserSchoolReportPlans(user_id string) ([]*model.SchoolReportPlanOverview, error) {
	plans := make([]*model.SchoolReportPlanOverview, 0)
	collection := database.Db.Collection("school_report_plans")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "plan_author"}}}}
	notes_stage := bson.D{{Key: "$count", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "note_count"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}, notes_stage, user_stage, unwind}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &plans)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func CreateSchoolReportPlan(new_plan model.NewSchoolReportPlan, plan_creator string) (*model.SchoolReportPlanRes, error) {
	collection := database.Db.Collection("school_report_plans")
	plan := model.SchoolReportPlan{
		ID:             uuid.New().String(),
		UserID:         plan_creator,
		Date:           new_plan.Date,
		CoFacilitators: new_plan.CoFacilitators,
		Curriculum:     new_plan.Curriculum,
		School:         new_plan.School,
		LessonTopics:   new_plan.LessonTopics,
		CreatedAt:      time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:      bson.TypeNull.String(),
		Status:         "pending",
	}
	res, err := collection.InsertOne(context.TODO(), plan)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert plan")
	}
	var plan_author model.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: plan_creator}}, options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})).Decode(&plan_author)
	if err != nil {
		return nil, err
	}
	return &model.SchoolReportPlanRes{
		ID:         plan.ID,
		PlanAuthor: &plan_author,
		Date:       plan.Date,
		School:     plan.School,
		Status:     plan.Status,
		CreatedAt:  plan.CreatedAt,
	}, nil
}
func UpdateSchoolReportPlan(update model.UpdateSchoolReportPlan, filter bson.D) (*model.SchoolReportPlan, error) {
	collection := database.Db.Collection("school_report_plans")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
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
	var srp model.SchoolReportPlan
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&srp)
	if err != nil {
		return nil, err
	}
	return &srp, nil
}
func DeleteSchoolReportPlan(filter bson.D) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
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
