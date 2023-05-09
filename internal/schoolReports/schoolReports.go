package schoolReports

import (
	"context"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchoolReportPlan struct {
	ID             string    `json:"id" bson:"_id"`
	UserID         *string   `json:"user_id" bson:"user_id"`
	Cofacilitators []*string `json:"cofacilitators" bson:"cofacilitators"`
	Curriculum     string    `json:"curriculum"`
	School         string    `json:"school"`
	LessonTopics   string    `json:"lesson_topics" bson:"lesson_topics"`
	Status         string    `json:"status"`
	CreatedAt      string    `json:"created_at" bson:"created_at"`
	UpdatedAt      string    `json:"updated_at" bson:"updated_at"`
}
type SchoolReportDebrief struct {
	ID                     string   `json:"id" bson:"_id"`
	UserID                 *string  `json:"user_id" bson:"user_id"`
	LessonPlanID           string   `json:"lesson_plan_id" bson:"lesson_plan_id"`
	StudentCount           int      `json:"student_count" bson:"student_count"`
	StudentList            []string `json:"student_list" bson:"student_list"`
	ChallengesImprovements string   `json:"challenges_improvements" bson:"challenges_improvements"`
	Positives              string   `json:"positives"`
	Discussion             string   `json:"discussion"`
	Status                 string   `json:"status"`
	CreatedAt              string   `json:"created_at" bson:"created_at"`
	UpdatedAt              string   `json:"updated_at" bson:"updated_at"`
}

func (e *SchoolReportPlan) Create() (*SchoolReportPlan, error) {
	collection := database.Db.Collection("school_report_plans")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	e.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *SchoolReportPlan) Update(id string) (*SchoolReportPlan, error) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "curriculum", Value: e.Curriculum},
			{Key: "school", Value: e.School},
			{Key: "lesson_topics", Value: e.LessonTopics},
			{Key: "cofacilitators", Value: e.Cofacilitators},
			{Key: "status", Value: e.Status},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *SchoolReportPlan) Delete(id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}

func (e *SchoolReportPlan) Approve(id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

func (e *SchoolReportPlan) Reject(id string) (bool, error) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
func (e *SchoolReportDebrief) Create() (*SchoolReportDebrief, error) {
	collection := database.Db.Collection("school_report_debriefs")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	e.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *SchoolReportDebrief) Update(id string) (*SchoolReportDebrief, error) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "student_count", Value: e.StudentCount},
			{Key: "student_list", Value: e.StudentList},
			{Key: "discussion", Value: e.Discussion},
			{Key: "positives", Value: e.Positives},
			{Key: "challenges_improvements", Value: e.ChallengesImprovements},
			{Key: "status", Value: e.Status},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *SchoolReportDebrief) Delete(id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}

func (e *SchoolReportDebrief) Approve(id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

func (e *SchoolReportDebrief) Reject(id string) (bool, error) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
