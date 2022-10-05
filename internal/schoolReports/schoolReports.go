package schoolReports

import (
	"context"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (e *SchoolReportPlan) Create() {
	collection := database.Db.Collection("school_report_plans")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	e.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		panic(err)
	}
}

func (e *SchoolReportPlan) Update(id string) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"curriculum", e.Curriculum},
			{"school", e.School},
			{"lesson_topics", e.LessonTopics},
			{"cofacilitators", e.Cofacilitators},
			{"status", e.Status},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportPlan) Delete(id string) {
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{"_id", id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if result.DeletedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportPlan) Approve(id string) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "approved"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportPlan) Reject(id string) {
	collection := database.Db.Collection("school_report_plans")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "rejected"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}
func (e *SchoolReportDebrief) Create() {
	collection := database.Db.Collection("school_report_debriefs")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = e.CreatedAt
	e.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		panic(err)
	}
}

func (e *SchoolReportDebrief) Update(id string) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"student_count", e.StudentCount},
			{"student_list", e.StudentList},
			{"discussion", e.Discussion},
			{"positives", e.Positives},
			{"challenges_improvements", e.ChallengesImprovements},
			{"status", e.Status},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportDebrief) Delete(id string) {
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{"_id", id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if result.DeletedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportDebrief) Approve(id string) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "approved"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}

func (e *SchoolReportDebrief) Reject(id string) {
	collection := database.Db.Collection("school_report_debriefs")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", "rejected"},
			{"updated_at", e.UpdatedAt},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("No Report Found")
	}
}
