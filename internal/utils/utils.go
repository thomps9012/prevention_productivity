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


func GetEventSummaries(filter bson.D) ([]*model.AllEventSummaries, error) {
	// this function breaks if logs don't meet the model requirements
	println("hit")
	eventSummariesCollection := database.Db.Collection("event_summaries")
	eventsCollection := database.Db.Collection("events")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allEventSummaries []*model.AllEventSummaries

	cursor, err := eventSummariesCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var eventSummary *model.EventSummary
		err := cursor.Decode(&eventSummary)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", eventSummary.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", eventSummary.UserID}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		var event *model.Event
		eventFilter := bson.D{{"_id", eventSummary.EventID}}
		err = eventsCollection.FindOne(context.TODO(), eventFilter).Decode(&event)
		if err != nil {
			return nil, err
		}
		singleEventSummary := &model.AllEventSummaries{
			Event: &model.Event{
				ID:        event.ID,
				Title: event.Title,
				StartDate: event.StartDate,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
			EventSummary: &model.EventSummary{
				ID:        eventSummary.ID,
				EventID:   eventSummary.EventID,
				AttendeeCount: eventSummary.AttendeeCount,
				Status:    eventSummary.Status,
				CreatedAt: eventSummary.CreatedAt,
				UpdatedAt: eventSummary.UpdatedAt,
			},
		}
		allEventSummaries = append(allEventSummaries, singleEventSummary)
	}
	return allEventSummaries, nil
}

func GetSchoolReports(filter bson.D) ([]*model.AllSchoolReports, error) {
	// this function breaks if logs don't meet the model requirements
	schoolReportsCollection := database.Db.Collection("school_reports")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allSchoolReports []*model.AllSchoolReports

	cursor, err := schoolReportsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var schoolReport *model.SchoolReport
		err := cursor.Decode(&schoolReport)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", schoolReport.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", schoolReport.UserID}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		singleSchoolReport := &model.AllSchoolReports{
			SchoolReport: &model.SchoolReport{
				ID:        schoolReport.ID,
				Curriculum: schoolReport.Curriculum,
				Status:    schoolReport.Status,
				CreatedAt: schoolReport.CreatedAt,
				UpdatedAt: schoolReport.UpdatedAt,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
		}
		allSchoolReports = append(allSchoolReports, singleSchoolReport)
	}
	return allSchoolReports, nil
}