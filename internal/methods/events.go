package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindEventDetails(event_id string) (*model.EventWithNotes, error) {
	return nil, errors.New("method unimplemented")
}
func FindEvents(filter bson.D) ([]*model.EventOverview, error) {
	return nil, errors.New("method unimplemented")
}
func FindUserEvents(user_id string) ([]*model.EventOverview, error) {
	return nil, errors.New("method unimplemented")
}

func CreateEvent(new_event model.NewEvent, event_creator string) (*model.EventRes, error) {
	collection := database.Db.Collection("events")
	event := model.Event{
		ID:                      uuid.New().String(),
		Title:                   new_event.Title,
		Description:             new_event.Description,
		UserID:                  event_creator,
		CoPlanners:              new_event.CoPlanners,
		StartDate:               new_event.StartDate,
		SetUp:                   new_event.SetUp,
		CleanUp:                 new_event.CleanUp,
		EndDate:                 new_event.EndDate,
		GrantID:                 new_event.GrantID,
		PublicEvent:             new_event.PublicEvent,
		RsvpRequired:            new_event.RsvpRequired,
		AnnualEvent:             new_event.AnnualEvent,
		NewEvent:                new_event.NewEvent,
		VolunteersNeeded:        new_event.VolunteersNeeded,
		Agenda:                  new_event.Agenda,
		TargetAudience:          new_event.TargetAudience,
		PartingGifts:            new_event.PartingGifts,
		MarketingMaterial:       new_event.MarketingMaterials,
		Supplies:                new_event.Supplies,
		SpecialOrders:           new_event.SpecialOrders,
		Performance:             *new_event.Performance,
		Vendors:                 *new_event.Vendors,
		FoodAndBeverage:         new_event.FoodAndBeverage,
		Caterer:                 *new_event.Caterer,
		FoodHeadCount:           new_event.FoodHeadCount,
		EventTeam:               new_event.EventTeam,
		VolunteerList:           new_event.VolunteerList,
		Budget:                  new_event.Budget,
		AffiliatedOrganizations: new_event.AffiliatedOrganizations,
		EducationalGoals:        new_event.EducationalGoals,
		EducationalOutcomes:     new_event.EducationalOutcomes,
		GrantGoals:              new_event.GrantGoals,
		Status:                  "pending",
		CreatedAt:               time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:               bson.TypeNull.String(),
	}
	var event_lead_info model.UserOverview
	opts := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})
	res, err := collection.InsertOne(context.TODO(), event)
	if err != nil {
		return nil, err
	}
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: event_creator}}, opts).Decode(&event_lead_info)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to create event")
	}
	return &model.EventRes{
		ID:        event.ID,
		UserID:    &event_lead_info,
		Title:     event.Title,
		StartDate: event.StartDate,
		Status:    event.Status,
		CreatedAt: event.CreatedAt,
	}, nil
}
func UpdateEvent(update model.UpdateEvent, filter bson.D) (*model.Event, error) {
	collection := database.Db.Collection("events")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	after := options.After
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	var e model.Event
	err := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$set", Value: update}, {Key: "$set", Value: bson.D{{Key: "updated_at", Value: updated_at}}}}, &opts).Decode(&e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func DeleteEvent(filter bson.D) (bool, error) {
	collection := database.Db.Collection("events")
	res, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return res.DeletedCount == 1, nil
}
func ApproveEvent(event_id string) (bool, error) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": event_id}
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
func RejectEvent(event_id string) (bool, error) {
	collection := database.Db.Collection("events")
	filter := bson.M{"_id": event_id}
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
