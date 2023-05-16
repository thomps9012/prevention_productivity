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
func FindSchoolReportPlanDetail(filter bson.D) (*model.SchoolReportPlanWithNotes, error) {
	plan_detail := make([]*model.SchoolReportPlanWithNotes, 0)
	collection := database.Db.Collection("school_report_plans")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "plan_author"}}}}
	educator_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "contacts"}, {Key: "localField", Value: "educator"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "educator"}}}}
	co_facilitators := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "co_facilitators"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "co_facilitators"}}}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "author"}}}},
		bson.D{{Key: "$unwind", Value: "$author"}},
	}}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	unwind_educator := bson.D{{Key: "$unwind", Value: "$educator"}}
	pipeline := mongo.Pipeline{filter, notes_stage, user_stage, co_facilitators, educator_stage, unwind, unwind_educator}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &plan_detail)
	if err != nil {
		return nil, err
	}
	if len(plan_detail) == 0 {
		return nil, errors.New("you're attempting to view a plan that either doesn't exist, or you didn't create")
	}
	return plan_detail[0], nil
}
func FindSchoolReportPlans(filter bson.D) ([]*model.SchoolReportPlanOverview, error) {
	plans := make([]*model.SchoolReportPlanOverview, 0)
	collection := database.Db.Collection("school_report_plans")
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "plan_author"}}}}
	educator_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "contacts"}, {Key: "localField", Value: "educator"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "educator"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	unwind_educator := bson.D{{Key: "$unwind", Value: "$educator"}}
	pipeline := mongo.Pipeline{filter, note_stage, note_count, user_stage, educator_stage, unwind, unwind_educator}
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
	educator_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "contacts"}, {Key: "localField", Value: "educator"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "educator"}}}}
	note_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	note_count := bson.D{{Key: "$addFields", Value: bson.M{"note_count": bson.M{"$size": "$notes"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$plan_author"}}
	unwind_educator := bson.D{{Key: "$unwind", Value: "$educator"}}
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: user_id}}}}, note_stage, note_count, user_stage, educator_stage, unwind, unwind_educator}
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

// mutations
func CreateSchoolReportPlan(new_plan model.NewSchoolReportPlan, plan_creator string) (*model.SchoolReportPlanRes, error) {
	collection := database.Db.Collection("school_report_plans")
	plan := model.SchoolReportPlan{
		ID:             uuid.New().String(),
		UserID:         plan_creator,
		Date:           new_plan.Date,
		Educator:       new_plan.Educator,
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
	var plan model.SchoolReportPlan
	err := collection.FindOne(context.TODO(), filter).Decode(&plan)
	if err != nil {
		return nil, errors.New("you're attempting to update a plan that either doesn't exist or you didn't create")
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "curriculum", Value: update.Curriculum},
			{Key: "co_facilitators", Value: update.CoFacilitators},
			{Key: "educator", Value: update.Educator},
			{Key: "school", Value: update.School},
			{Key: "date", Value: update.Date},
			{Key: "lesson_topics", Value: update.LessonTopics},
			{Key: "co_facilitators", Value: update.CoFacilitators},
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
	err = collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&srp)
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
