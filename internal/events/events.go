package events

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	ID                      string    `json:"id" bson:"_id"`
	EventLead               *string   `json:"event_lead" bson:"event_lead"`
	Co_Planners             []*string `json:"co_planners" bson:"co_planners"`
	Title                   string    `json:"title" bson:"title"`
	Description             string    `json:"description" bson:"description"`
	StartDate               string    `json:"start_date" bson:"start_date"`
	SetUp                   string    `json:"set_up" bson:"set_up"`
	CleanUp                 string    `json:"clean_up" bson:"clean_up"`
	EndDate                 string    `json:"end_date" bson:"end_date"`
	GrantID                 string    `json:"grant_id" bson:"grant_id"`
	PublicEvent             bool      `json:"public_event" bson:"public_event"`
	RSVPRequired            bool      `json:"rsvp_required" bson:"rsvp_required"`
	AnnualEvent             bool      `json:"annual_event" bson:"annual_event"`
	NewEvent                bool      `json:"new_event" bson:"new_event"`
	VolunteersNeeded        bool      `json:"volunteers_needed" bson:"volunteers_needed"`
	Agenda                  []*string `json:"agenda" bson:"agenda"`
	TargetAudience          string    `json:"target_audience" bson:"target_audience"`
	PartingGifts            []*string `json:"parting_gifts" bson:"parting_gifts"`
	MarketingMaterials      []*string `json:"marketing_materials" bson:"marketing_materials"`
	Supplies                []*string `json:"supplies" bson:"supplies"`
	SpecialOrders           []*string `json:"special_orders" bson:"special_orders"`
	Performance             string    `json:"performance" bson:"performance"`
	Vendors                 string    `json:"vendors" bson:"vendors"`
	FoodAndBeverage         []*string `json:"food_and_beverage" bson:"food_and_beverage"`
	Caterer                 string    `json:"caterer" bson:"caterer"`
	FoodHeadCount           int       `json:"food_head_count" bson:"food_head_count"`
	EventTeam               []*string `json:"event_team" bson:"event_team"`
	VolunteerList           []*string `json:"volunteer_list" bson:"volunteer_list"`
	Budget                  float64   `json:"budget" bson:"budget"`
	AffiliatedOrganizations *string   `json:"affiliated_organizations" bson:"affiliated_organizations"`
	EducationalGoals        []*string `json:"educational_goals" bson:"educational_goals"`
	EducationalOutcomes     []*string `json:"educational_outcomes" bson:"educational_outcomes"`
	GrantGoals              []*string `json:"grant_goals" bson:"grant_goals"`
	CreatedAt               string    `json:"created_at" bson:"created_at"`
	UpdatedAt               string    `json:"updated_at" bson:"updated_at"`
	Status                  string    `json:"status" bson:"status"`
}

func (e *Event) Create() (*Event, error) {
	collection := database.Db.Collection("events")
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	e.Status = "pending"
	res, err := collection.InsertOne(context.TODO(), e)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create event")
	}
	return e, nil
}

func (e *Event) Update() (*Event, error) {
	collection := database.Db.Collection("events")
	e.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.M{"_id": e.ID}
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$set", Value: e}}, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Event) Delete() (*Event, error) {
	collection := database.Db.Collection("events")
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": e.ID})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, errors.New("failed to delete event")
	}
	return e, nil
}

func (e *Event) Approve(id string) (bool, error) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": id}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "approved"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}

func (e *Event) Reject(id string) (bool, error) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": id}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "rejected"},
			{Key: "updated_at", Value: time.Now().Format("2006-01-02 15:04:05")},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount == 1, nil
}
