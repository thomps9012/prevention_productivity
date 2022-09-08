package utils

import (
	"context"
	"fmt"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

func GetLogs(filter bson.D) ([]*model.AllLogs, error) {
	// this times out
	logsCollection := database.Db.Collection("logs")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	cursor, err := logsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var allLogs []*model.AllLogs
	for cursor.Next(context.TODO()) {
		println("logs found")
		var log *model.Log
		err := cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", log.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", log.UserID}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		singleLog := &model.AllLogs{
			Log: &model.Log{
				ID:        log.ID,
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
		allLogs = append(allLogs, singleLog)
	}
	fmt.Printf("%v logs being returned", allLogs)
	return allLogs, nil
}

func GetEvents(filter bson.M) ([]*model.AllEvents, error) {
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
				Title:     event.Title,
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
				Title:     event.Title,
				StartDate: event.StartDate,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
			EventSummary: &model.EventSummary{
				ID:            eventSummary.ID,
				EventID:       eventSummary.EventID,
				AttendeeCount: eventSummary.AttendeeCount,
				Status:        eventSummary.Status,
				CreatedAt:     eventSummary.CreatedAt,
				UpdatedAt:     eventSummary.UpdatedAt,
			},
		}
		allEventSummaries = append(allEventSummaries, singleEventSummary)
	}
	return allEventSummaries, nil
}

func GetSchoolReportPlans(filter bson.M) ([]*model.AllSchoolReportPlans, error) {
	// this function breaks if logs don't meet the model requirements
	schoolReportsCollection := database.Db.Collection("school_report_plans")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allSchoolReportPlans []*model.AllSchoolReportPlans
	fmt.Printf("func filter %v", filter)
	cursor, err := schoolReportsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	println("hit1")
	for cursor.Next(context.TODO()) {
		println("hit2")
		var schoolReportPlan *model.SchoolReportPlan
		err := cursor.Decode(&schoolReportPlan)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", schoolReportPlan.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", schoolReportPlan.UserID}}
		usererr := userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if usererr != nil {
			return nil, usererr
		}
		singleSchoolReportPlan := &model.AllSchoolReportPlans{
			SchoolReportPlan: &model.SchoolReportPlan{
				ID:         schoolReportPlan.ID,
				Curriculum: schoolReportPlan.Curriculum,
				Status:     schoolReportPlan.Status,
				CreatedAt:  schoolReportPlan.CreatedAt,
				UpdatedAt:  schoolReportPlan.UpdatedAt,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
		}
		allSchoolReportPlans = append(allSchoolReportPlans, singleSchoolReportPlan)
	}
	return allSchoolReportPlans, nil
}
func GetSchoolReportDebriefs(filter bson.D) ([]*model.AllSchoolReportDebriefs, error) {
	// this function breaks if logs don't meet the model requirements
	schoolReportsCollection := database.Db.Collection("school_report_debriefs")
	notesCollection := database.Db.Collection("notes")
	userCollection := database.Db.Collection("users")
	var allSchoolReportDebriefs []*model.AllSchoolReportDebriefs

	cursor, err := schoolReportsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var schoolReportDebrief *model.SchoolReportDebrief
		err := cursor.Decode(&schoolReportDebrief)
		if err != nil {
			return nil, err
		}
		noteFilter := bson.D{{"item_id", schoolReportDebrief.ID}}
		noteCount, noteErr := notesCollection.CountDocuments(context.TODO(), noteFilter)
		if noteErr != nil {
			return nil, err
		}
		intNoteCount := int(noteCount)
		var user *model.User
		userFilter := bson.D{{"_id", schoolReportDebrief.UserID}}
		err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
		if err != nil {
			return nil, err
		}
		singleSchoolReportDebrief := &model.AllSchoolReportDebriefs{
			SchoolReportDebrief: &model.SchoolReportDebrief{
				ID:           schoolReportDebrief.ID,
				LessonPlanID: schoolReportDebrief.LessonPlanID,
				Status:       schoolReportDebrief.Status,
				CreatedAt:    schoolReportDebrief.CreatedAt,
				UpdatedAt:    schoolReportDebrief.UpdatedAt,
			},
			User: &model.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			NoteCount: &intNoteCount,
		}
		allSchoolReportDebriefs = append(allSchoolReportDebriefs, singleSchoolReportDebrief)
	}
	return allSchoolReportDebriefs, nil
}
