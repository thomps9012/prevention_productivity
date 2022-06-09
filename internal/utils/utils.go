package utils

import (
	"context"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"fmt"
)

func GetLogs(filter bson.D) ([]*model.AllLogs, error) {
	// this function breaks if logs don't meet the model requirements
	logsCollection := database.Db.Collection("logs")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allLogs []*model.AllLogs

	cursor, err := logsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var log *model.Log
		err := cursor.Decode(&log)
		println(log.FocusArea)
		println(log.UpdatedAt)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", log.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		println(noteCount)
		var user *model.User
		fmt.Println("%v", log.UserID)
		userFilter := bson.D{{"_id", log.UserID}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		singleLog := &model.AllLogs{
			Log: &model.Log{
				ID:        log.ID,
				FocusArea: log.FocusArea,
				Status:    log.Status,
				CreatedAt: log.CreatedAt,
				UpdatedAt: log.UpdatedAt,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
		}
		println(singleLog)
		allLogs = append(allLogs, singleLog)
	}
	return allLogs, nil
}

func GetEvents(filter bson.D) ([]*model.AllEvents, error) {
	// this function breaks if logs don't meet the model requirements
	eventsCollection := database.Db.Collection("events")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allEvents []*model.AllEvents

	cursor, err := eventsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var event *model.Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", event.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", event.EventLead}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		singleEvent := &model.AllEvents{
			Event: &model.Event{
				ID:        event.ID,
				Title: event.Title,
				StartDate: event.StartDate,
				Status:    event.Status,
				CreatedAt: event.CreatedAt,
				UpdatedAt: event.UpdatedAt,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
		}
		allEvents = append(allEvents, singleEvent)
	}
	return allEvents, nil
}
