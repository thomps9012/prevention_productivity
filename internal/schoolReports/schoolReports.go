package schoolReports

import (
		database "thomps9012/prevention_productivity/internal/db"
		"go.mongodb.org/mongo-driver/bson"
		"context"
		"strings"
		"time"
		"github.com/google/uuid"
)

type SchoolReport struct {
		ID           string  `json:"id" bson:"id"`
		UserID       *string  `json:"user_id" bson:"user_id"`
		Curriculum   string   `json:"curriculum"`
		LessonPlan   string   `json:"lesson_plan" bson:"lesson_plan"`
		School       string   `json:"school"`
		Topics       string   `json:"topics"`
		StudentCount int      `json:"student_count" bson:"student_count"`
		StudentList  []string `json:"student_list" bson:"student_list"`
		Challenges   string   `json:"challenges"`
		Successes    string   `json:"successes"`
		Improvements string   `json:"improvements"`
		Status       string   `json:"status"`
		CreatedAt    string   `json:"created_at" bson:"created_at"`
		UpdatedAt    string   `json:"updated_at" bson:"updated_at"`
}

func (e *SchoolReport) Create() {
		collection := database.Db.Collection("school_reports")
		e.ID = strings.Replace(uuid.New().String(), "-", "", -1)
		e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		e.UpdatedAt = e.CreatedAt
		e.Status = "pending"
		_, err := collection.InsertOne(context.TODO(), e)
		if err != nil {
				panic(err)
		}
}

func (e *SchoolReport) Update(id string) {
		collection := database.Db.Collection("school_reports")
		e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		filter := bson.D{{"id", id}}
		update := bson.D{
			{"$set", bson.D{
				{"curriculum", e.Curriculum},
				{"lesson_plan", e.LessonPlan},
				{"school", e.School},
				{"topics", e.Topics},
				{"student_count", e.StudentCount},
				{"student_list", e.StudentList},
				{"challenges", e.Challenges},
				{"successes", e.Successes},
				{"improvements", e.Improvements},
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

func (e *SchoolReport) Delete(id string) {
		collection := database.Db.Collection("school_reports")
		filter := bson.D{{"id", id}}
		result, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
				panic(err)
		}
		if result.DeletedCount == 0 {
			panic("No Report Found")
		}
}

func (e *SchoolReport) Approve(id string) {
		collection := database.Db.Collection("school_reports")
		e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		filter := bson.D{{"id", id}}
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

func (e *SchoolReport) Reject(id string) {
		collection := database.Db.Collection("school_reports")
		e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		filter := bson.D{{"id", id}}
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

