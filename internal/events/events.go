package events

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"strings"
	"time"
	"github.com/google/uuid"
)


type Event struct {
	ID	   string `json:"id" bson:"_id"`
	EventLead string `json:"event_lead" bson:"event_lead"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	StartDate              string        `json:"start_date" bson:"start_date"`
	SetUp                  string        `json:"set_up" bson:"set_up"`
	CleanUp                string        `json:"clean_up" bson:"clean_up"`
	EndDate                string        `json:"end_date" bson:"end_date"`
	GrantID                string        `json:"grant_id" bson:"grant_id"`
	Public                 bool          `json:"public"`
	Rsvp                   bool          `json:"rsvp"`
	AnnualEvent            bool          `json:"annual_event" bson:"annual_event"`
	NewEvent               bool          `json:"new_event" bson:"new_event"`
	Volunteers             bool          `json:"volunteers"`
	Agenda                 []*string `json:"agenda"`
	TargetAudience         string        `json:"target_audience" bson:"target_audience"`
	PartingGifts           []*string  `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterial      []*string  `json:"marketing_material" bson:"marketing_material"`
	Supplies               []*string  `json:"supplies"`
	SpecialOrders          []*string  `json:"special_orders" bson:"special_orders"`
	Performance            string        `json:"performance"`
	Vendors                string        `json:"vendors"`
	FoodAndBeverage        []*string  `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                string        `json:"caterer"`
	FoodHeadCount          int           `json:"food_head_count" bson:"food_head_count"`
	EventTeam              []*string     `json:"event_team" bson:"event_team"`
	VolunteerList          []*string    `json:"volunteer_list" bson:"volunteer_list"`
	Budget                 float64       `json:"budget"`
	AffiliatedOrganization *string       `json:"affiliated_organization" bson:"affiliated_organization"`
	EducationalGoals       []*string      `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes    []*string      `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals             []*string      `json:"grant_goals" bson:"grant_goals"`
	CreatedAt              string        `json:"created_at" bson:"created_at"`
	UpdatedAt              string        `json:"updated_at" bson:"updated_at"`
	Status 			   string        `json:"status" bson:"status"`
}


func (e *Event) Create() {
	collection := database.Db.Collection("events")
	e.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.Status = "pending"
	_, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		panic(err)
	}
}

func (e *Event) Update() {
	collection := database.Db.Collection("events")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.M{"_id": e.ID}
	_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", e}})
	if err != nil {
		panic(err)
	}
}

func (e *Event) Delete() {
	collection := database.Db.Collection("events")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": e.ID})
	if err != nil {
		panic(err)
	}
}

func (e *Event) Approve(id string) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", "approved"},
			{"updated_at", time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Event not found")
	}
}

func (e *Event) Reject(id string) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"status", "rejected"},
			{"updated_at", time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		panic("Event not found")
	}
}