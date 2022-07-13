package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "thomps9012/prevention_productivity/graph/generated"
	"thomps9012/prevention_productivity/graph/model"
	"thomps9012/prevention_productivity/internal/auth"
	"thomps9012/prevention_productivity/internal/contacts"
	database "thomps9012/prevention_productivity/internal/db"
	"thomps9012/prevention_productivity/internal/eventSummaries"
	"thomps9012/prevention_productivity/internal/events"
	"thomps9012/prevention_productivity/internal/grants"
	"thomps9012/prevention_productivity/internal/jwt"
	"thomps9012/prevention_productivity/internal/logs"
	"thomps9012/prevention_productivity/internal/notes"
	"thomps9012/prevention_productivity/internal/schoolReports"
	"thomps9012/prevention_productivity/internal/users"
	"thomps9012/prevention_productivity/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (string, error) {
	var user users.User
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.Email = newUser.Email
	user.Password = newUser.Password
	user.Create()
	if user.ID == "" {
		return "", fmt.Errorf("User already exists")
	}
	token, err := jwt.GenerateToken(user.Email, user.IsAdmin, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser, id string) (*model.User, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	println("userid", id)
	println("editorID", userID)
	if !isAdmin && id != userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"_id", id}}
	var user users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email
	user.Password = updateUser.Password
	user.IsActive = updateUser.IsActive
	user.IsAdmin = updateUser.IsAdmin
	count, err := collection.CountDocuments(context.TODO(), bson.D{{"email", updateUser.Email}})
	if err != nil {
		fmt.Println(err)
	}
	if count > 1 {
		return nil, fmt.Errorf("user already exists")
	}
	user.Update(id)
	return &model.User{
		ID:        &user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &user.DeletedAt,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return false, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"_id", id}}
	var user users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false, err
	}
	user.Delete()
	return true, nil
}

func (r *mutationResolver) Login(ctx context.Context, login model.LoginInput) (string, error) {
	var user users.User
	user.Email = login.Email
	user.Password = login.Password
	correct := user.Authenticate()
	if !correct {
		return "Invalid email or password", fmt.Errorf("Invalid email or password")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"email", login.Email}}
	var userDB users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&userDB)
	println(userDB.IsAdmin)
	println(userDB.ID)
	if !userDB.IsActive {
		return "User is not active", fmt.Errorf("User is not active")
	}
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(userDB.Email, userDB.IsAdmin, userDB.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken model.RefreshTokenInput) (string, error) {
	token, err := jwt.ParseToken(refreshToken.Token)
	if err != nil {
		return "", err
	}
	newToken, err := jwt.GenerateToken(token["email"].(string), token["isAdmin"].(bool), token["userID"].(string))
	if err != nil {
		return "", err
	}
	return newToken, nil
}

func (r *mutationResolver) CreateGrant(ctx context.Context, newGrant model.NewGrant) (*model.Grant, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var grant grants.Grant
	grant.Name = *newGrant.Name
	grant.CreatedBy = userID
	grant.Description = *newGrant.Description
	grant.Goals = newGrant.Goals
	grant.Objectives = newGrant.Objectives
	grant.StartDate = *newGrant.StartDate
	grant.EndDate = *newGrant.EndDate
	grant.Budget = *newGrant.Budget
	grant.AwardNumber = *newGrant.AwardNumber
	grant.AwardDate = *newGrant.AwardDate
	grant.Create()
	return &model.Grant{
		ID:          &grant.ID,
		CreatedBy:   grant.CreatedBy,
		Name:        grant.Name,
		Description: grant.Description,
		Goals:       grant.Goals,
		Objectives:  grant.Objectives,
		StartDate:   grant.StartDate,
		EndDate:     grant.EndDate,
		Budget:      &grant.Budget,
		AwardNumber: grant.AwardNumber,
		AwardDate:   &grant.AwardDate,
		CreatedAt:   grant.CreatedAt,
		UpdatedAt:   grant.UpdatedAt,
		IsActive:    grant.IsActive,
	}, nil
}

func (r *mutationResolver) UpdateGrant(ctx context.Context, id string, updateGrant model.UpdateGrant) (*model.Grant, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var grant grants.Grant
	grant.ID = id
	grant.Name = *updateGrant.Name
	grant.Description = *updateGrant.Description
	grant.Goals = updateGrant.Goals
	grant.Objectives = updateGrant.Objectives
	grant.StartDate = *updateGrant.StartDate
	grant.EndDate = *updateGrant.EndDate
	grant.Budget = *updateGrant.Budget
	grant.AwardNumber = *updateGrant.AwardNumber
	grant.AwardDate = *updateGrant.AwardDate
	grant.Update(id)
	return &model.Grant{
		ID:          &grant.ID,
		Name:        grant.Name,
		Description: grant.Description,
		Goals:       grant.Goals,
		Objectives:  grant.Objectives,
		StartDate:   grant.StartDate,
		EndDate:     grant.EndDate,
		Budget:      &grant.Budget,
		AwardNumber: grant.AwardNumber,
		AwardDate:   &grant.AwardDate,
		UpdatedAt:   grant.UpdatedAt,
		IsActive:    grant.IsActive,
	}, nil
}

func (r *mutationResolver) RemoveGrant(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	removed := false
	if !isAdmin {
		return &removed, fmt.Errorf("Unauthorized")
	}
	var grant grants.Grant
	grant.ID = id
	grant.Delete(id)
	removed = true
	return &removed, nil
}

func (r *mutationResolver) CreateContact(ctx context.Context, newContact model.NewContact) (*model.Contact, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var contact contacts.Contact
	contact.Name = *newContact.Name
	contact.Type = *newContact.Type
	contact.Email = *newContact.Email
	contact.Phone = *newContact.Phone
	contact.Notes = *newContact.Notes
	contact.CreatedBy = userID
	contact.Create()
	return &model.Contact{
		ID:        &contact.ID,
		Name:      &contact.Name,
		Type:      &contact.Type,
		Email:     &contact.Email,
		Phone:     &contact.Phone,
		Notes:     &contact.Notes,
		CreatedBy: contact.CreatedBy,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
		IsActive:  contact.IsActive,
	}, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id string, updateContact model.UpdateContact) (*model.Contact, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var contact contacts.Contact
	contact.ID = id
	contact.Name = *updateContact.Name
	contact.Type = *updateContact.Type
	contact.Email = *updateContact.Email
	contact.Phone = *updateContact.Phone
	contact.Notes = *updateContact.Notes
	contact.Update(id)
	return &model.Contact{
		ID:        &contact.ID,
		Name:      &contact.Name,
		Type:      &contact.Type,
		Email:     &contact.Email,
		Phone:     &contact.Phone,
		Notes:     &contact.Notes,
		UpdatedAt: contact.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveContact(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var contact contacts.Contact
	contact.Delete(id)
	// returns false if contact was deactivated
	var result bool
	if contact.IsActive {
		result = false
	} else {
		result = true
	}
	return &result, nil
}

func (r *mutationResolver) CreateNote(ctx context.Context, newNote model.NewNote) (*model.Note, error) {
	UserID := auth.ForUserID(ctx)
	if UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var note notes.Note
	note.UserID = UserID
	note.ItemID = newNote.ItemID
	note.Title = newNote.Title
	note.Content = newNote.Content
	note.Create()
	return &model.Note{
		ID:        &note.ID,
		ItemID:    &note.ItemID,
		UserID:    &note.UserID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}, nil
}

func (r *mutationResolver) UpdateNote(ctx context.Context, id string, updateNote model.UpdateNote) (*model.Note, error) {
	UserID := auth.ForUserID(ctx)
	IsAdmin := auth.ForAdmin(ctx)
	if UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("notes")
	filter := bson.D{{"_id", id}}
	var note notes.Note
	err := collection.FindOne(context.TODO(), filter).Decode(&note)
	if err != nil {
		return nil, err
	}
	if IsAdmin || note.UserID == UserID {
		note.Title = updateNote.Title
		note.Content = updateNote.Content
		note.Update()
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
	return &model.Note{
		ID:        &note.ID,
		UserID:    &note.UserID,
		ItemID:    &note.ItemID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveNote(ctx context.Context, id string) (bool, error) {
	UserID := auth.ForUserID(ctx)
	IsAdmin := auth.ForAdmin(ctx)
	if UserID == "" {
		return false, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("notes")
	filter := bson.D{{"_id", id}}
	var note notes.Note
	err := collection.FindOne(context.TODO(), filter).Decode(&note)
	if err != nil {
		return false, err
	}
	if IsAdmin || note.UserID == UserID {
		note.Remove(id)
	} else {
		return false, fmt.Errorf("Unauthorized")
	}
	return true, nil
}

func (r *mutationResolver) CreateLog(ctx context.Context, newLog model.NewLog) (*model.Log, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var log logs.Log
	log.UserID = userID
	log.FocusArea = newLog.FocusArea
	log.Actions = newLog.Actions
	log.Successes = newLog.Successes
	log.Improvements = newLog.Improvements
	log.NextSteps = newLog.NextSteps
	log.Create()
	return &model.Log{
		ID:           &log.ID,
		UserID:       &log.UserID,
		FocusArea:    log.FocusArea,
		Actions:      log.Actions,
		Successes:    log.Successes,
		Improvements: log.Improvements,
		NextSteps:    log.NextSteps,
		Status:       log.Status,
		CreatedAt:    log.CreatedAt,
	}, nil
}

func (r *mutationResolver) UpdateLog(ctx context.Context, id string, updateLog model.UpdateLog) (*model.Log, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	userIsAdmin := auth.ForAdmin(ctx)
	var log logs.Log
	err := collection.FindOne(context.TODO(), filter).Decode(&log)
	if err != nil {
		return nil, err
	}
	if userIsAdmin || log.UserID == userID {
		log.FocusArea = updateLog.FocusArea
		log.Status = updateLog.Status
		log.Actions = updateLog.Actions
		log.Successes = updateLog.Successes
		log.Improvements = updateLog.Improvements
		log.NextSteps = updateLog.NextSteps
		log.Update(id)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
	return &model.Log{
		ID:           &log.ID,
		UserID:       &log.UserID,
		FocusArea:    log.FocusArea,
		Actions:      log.Actions,
		Successes:    log.Successes,
		Improvements: log.Improvements,
		NextSteps:    log.NextSteps,
		Status:       log.Status,
		CreatedAt:    log.CreatedAt,
		UpdatedAt:    log.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveLog(ctx context.Context, id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return false, fmt.Errorf("Unauthorized")
	}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, fmt.Errorf("Log not found")
	}
	return true, nil
}

func (r *mutationResolver) ApproveLog(ctx context.Context, id string) (bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return false, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	var log logs.Log
	err := collection.FindOne(context.TODO(), filter).Decode(&log)
	if err != nil {
		return false, err
	}
	log.Approve(id)
	return true, nil
}

func (r *mutationResolver) RejectLog(ctx context.Context, id string) (bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return false, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("logs")
	filter := bson.D{{"_id", id}}
	var log logs.Log
	err := collection.FindOne(context.TODO(), filter).Decode(&log)
	if err != nil {
		return false, err
	}
	log.Reject(id)
	return true, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, newEvent model.NewEvent) (*model.Event, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var event events.Event
	event.EventLead = userID
	event.Title = *newEvent.Title
	event.Description = *newEvent.Description
	event.StartDate = *newEvent.StartDate
	event.SetUp = *newEvent.SetUp
	event.CleanUp = *newEvent.CleanUp
	event.EndDate = *newEvent.EndDate
	event.GrantID = *newEvent.GrantID
	event.Public = *newEvent.Public
	event.Rsvp = *newEvent.Rsvp
	event.AnnualEvent = *newEvent.AnnualEvent
	event.NewEvent = *newEvent.NewEvent
	event.Volunteers = *newEvent.Volunteers
	event.TargetAudience = *newEvent.TargetAudience
	event.Vendors = *newEvent.Vendors
	event.Caterer = *newEvent.Caterer
	event.FoodHeadCount = *newEvent.FoodHeadCount
	event.EventTeam = newEvent.EventTeam
	event.Budget = *newEvent.Budget
	event.AffiliatedOrganization = newEvent.AffiliatedOrganization
	event.EducationalGoals = newEvent.EducationalGoals
	event.EducationalOutcomes = newEvent.EducationalOutcomes
	event.GrantGoals = newEvent.GrantGoals
	event.Agenda = newEvent.Agenda
	event.PartingGifts = newEvent.PartingGifts
	event.VolunteerList = newEvent.VolunteerList
	event.Performance = *newEvent.Performance
	event.MarketingMaterial = newEvent.MarketingMaterial
	event.Supplies = newEvent.Supplies
	event.SpecialOrders = newEvent.SpecialOrders
	event.FoodAndBeverage = newEvent.FoodAndBeverage
	event.Create()
	return &model.Event{
		ID:          &event.ID,
		EventLead:   &event.EventLead,
		GrantID:     event.GrantID,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		CreatedAt:   event.CreatedAt,
		Status:      event.Status,
	}, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, updateEvent model.UpdateEvent) (*model.Event, error) {
	userID := auth.ForUserID(ctx)
	isAdmin := auth.ForAdmin(ctx)
	if userID == "" && !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("events")
	filter := bson.D{{"_id", id}}
	println(id)
	var event events.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		return nil, err
	}
	println(userID)
	println(event.EventLead)
	if isAdmin || event.EventLead == userID {
		event.Title = *updateEvent.Title
		event.Description = *updateEvent.Description
		event.StartDate = *updateEvent.StartDate
		event.SetUp = *updateEvent.SetUp
		event.CleanUp = *updateEvent.CleanUp
		event.EndDate = *updateEvent.EndDate
		event.GrantID = *updateEvent.GrantID
		event.Public = *updateEvent.Public
		event.Rsvp = *updateEvent.Rsvp
		event.AnnualEvent = *updateEvent.AnnualEvent
		event.NewEvent = *updateEvent.NewEvent
		event.Volunteers = *updateEvent.Volunteers
		event.Agenda = updateEvent.Agenda
		event.TargetAudience = *updateEvent.TargetAudience
		event.PartingGifts = updateEvent.PartingGifts
		event.MarketingMaterial = updateEvent.MarketingMaterial
		event.Supplies = updateEvent.Supplies
		event.SpecialOrders = updateEvent.SpecialOrders
		event.Performance = *updateEvent.Performance
		event.Vendors = *updateEvent.Vendors
		event.FoodAndBeverage = updateEvent.FoodAndBeverage
		event.Caterer = *updateEvent.Caterer
		event.FoodHeadCount = *updateEvent.FoodHeadCount
		event.EventTeam = updateEvent.EventTeam
		event.VolunteerList = updateEvent.VolunteerList
		event.Budget = *updateEvent.Budget
		event.AffiliatedOrganization = updateEvent.AffiliatedOrganization
		event.EducationalGoals = updateEvent.EducationalGoals
		event.EducationalOutcomes = updateEvent.EducationalOutcomes
		event.GrantGoals = updateEvent.GrantGoals
		event.Update()
	}
	return &model.Event{
		ID:          &event.ID,
		EventLead:   &event.EventLead,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Status:      event.Status,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveEvent(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("events")
	filter := bson.D{{"_id", id}}
	var event events.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		result = false
		return &result, err
	}
	event.Delete()
	result = true
	return &result, nil
}

func (r *mutationResolver) ApproveEvent(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("events")
	filter := bson.D{{"_id", id}}
	var event events.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		result = false
		return &result, err
	}
	event.Approve(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) RejectEvent(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("events")
	filter := bson.D{{"_id", id}}
	var event events.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		result = false
		return &result, err
	}
	event.Reject(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) CreateEventSummary(ctx context.Context, newEventSummary model.NewEventSummary) (*model.EventSummary, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var eventSummary eventSummaries.EventSummary
	eventSummary.EventID = *newEventSummary.EventID
	eventSummary.UserID = userID
	eventSummary.AttendeeCount = *newEventSummary.AttendeeCount
	eventSummary.Challenges = *newEventSummary.Challenges
	eventSummary.Successes = *newEventSummary.Successes
	eventSummary.Improvements = *newEventSummary.Improvements
	eventSummary.Create()
	return &model.EventSummary{
		ID:            eventSummary.ID,
		EventID:       eventSummary.EventID,
		UserID:        eventSummary.UserID,
		AttendeeCount: eventSummary.AttendeeCount,
		Challenges:    eventSummary.Challenges,
		Successes:     eventSummary.Successes,
		Improvements:  eventSummary.Improvements,
		Status:        eventSummary.Status,
		CreatedAt:     eventSummary.CreatedAt,
		UpdatedAt:     eventSummary.UpdatedAt,
	}, nil
}

func (r *mutationResolver) UpdateEventSummary(ctx context.Context, id string, updateEventSummary model.UpdateEventSummary) (*model.EventSummary, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var eventSummary eventSummaries.EventSummary
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&eventSummary)
	if err != nil {
		return nil, err
	}
	if !isAdmin && eventSummary.UserID != userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	eventSummary.AttendeeCount = *updateEventSummary.AttendeeCount
	eventSummary.Challenges = *updateEventSummary.Challenges
	eventSummary.Successes = *updateEventSummary.Successes
	eventSummary.Improvements = *updateEventSummary.Improvements
	eventSummary.Update(id)
	return &model.EventSummary{
		ID:            eventSummary.ID,
		EventID:       eventSummary.EventID,
		UserID:        eventSummary.UserID,
		AttendeeCount: eventSummary.AttendeeCount,
		Challenges:    eventSummary.Challenges,
		Successes:     eventSummary.Successes,
		Improvements:  eventSummary.Improvements,
		Status:        eventSummary.Status,
		CreatedAt:     eventSummary.CreatedAt,
		UpdatedAt:     eventSummary.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveEventSummary(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{"_id", id}}
	var eventSummary eventSummaries.EventSummary
	err := collection.FindOne(context.TODO(), filter).Decode(&eventSummary)
	if err != nil {
		result = false
		return &result, err
	}
	eventSummary.Delete(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) ApproveEventSummary(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{"_id", id}}
	var eventSummary eventSummaries.EventSummary
	err := collection.FindOne(context.TODO(), filter).Decode(&eventSummary)
	if err != nil {
		result = false
		return &result, err
	}
	eventSummary.Approve(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) RejectEventSummary(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{"_id", id}}
	var eventSummary eventSummaries.EventSummary
	err := collection.FindOne(context.TODO(), filter).Decode(&eventSummary)
	if err != nil {
		result = false
		return &result, err
	}
	eventSummary.Reject(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) CreateSchoolReport(ctx context.Context, newSchoolReport model.NewSchoolReport) (*model.SchoolReport, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var schoolReport schoolReports.SchoolReport
	schoolReport.Curriculum = *newSchoolReport.Curriculum
	schoolReport.LessonPlan = *newSchoolReport.LessonPlan
	schoolReport.UserID = &userID
	schoolReport.School = *newSchoolReport.School
	schoolReport.Topics = *newSchoolReport.Topics
	schoolReport.StudentCount = *newSchoolReport.StudentCount
	schoolReport.StudentList = newSchoolReport.StudentList
	schoolReport.Challenges = *newSchoolReport.Challenges
	schoolReport.Successes = *newSchoolReport.Successes
	schoolReport.Improvements = *newSchoolReport.Improvements
	schoolReport.Create()
	return &model.SchoolReport{
		ID:           &schoolReport.ID,
		UserID:       schoolReport.UserID,
		Curriculum:   schoolReport.Curriculum,
		LessonPlan:   schoolReport.LessonPlan,
		School:       schoolReport.School,
		Topics:       schoolReport.Topics,
		StudentCount: schoolReport.StudentCount,
		StudentList:  schoolReport.StudentList,
		Challenges:   schoolReport.Challenges,
		Successes:    schoolReport.Successes,
		Improvements: schoolReport.Improvements,
		CreatedAt:    schoolReport.CreatedAt,
		UpdatedAt:    schoolReport.UpdatedAt,
		Status:       schoolReport.Status,
	}, nil
}

func (r *mutationResolver) UpdateSchoolReport(ctx context.Context, id string, updateSchoolReport model.UpdateSchoolReport) (*model.SchoolReport, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReport schoolReports.SchoolReport
	collection := database.Db.Collection("school_reports")
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReport)
	if err != nil {
		return nil, err
	}
	if !isAdmin && schoolReport.UserID != &userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	schoolReport.Curriculum = *updateSchoolReport.Curriculum
	schoolReport.LessonPlan = *updateSchoolReport.LessonPlan
	schoolReport.School = *updateSchoolReport.School
	schoolReport.Topics = *updateSchoolReport.Topics
	schoolReport.StudentCount = *updateSchoolReport.StudentCount
	schoolReport.StudentList = updateSchoolReport.StudentList
	schoolReport.Challenges = *updateSchoolReport.Challenges
	schoolReport.Successes = *updateSchoolReport.Successes
	schoolReport.Improvements = *updateSchoolReport.Improvements
	schoolReport.Update(id)
	return &model.SchoolReport{
		ID:           &schoolReport.ID,
		UserID:       schoolReport.UserID,
		Curriculum:   schoolReport.Curriculum,
		LessonPlan:   schoolReport.LessonPlan,
		School:       schoolReport.School,
		Topics:       schoolReport.Topics,
		StudentCount: schoolReport.StudentCount,
		StudentList:  schoolReport.StudentList,
		Challenges:   schoolReport.Challenges,
		Successes:    schoolReport.Successes,
		Improvements: schoolReport.Improvements,
		CreatedAt:    schoolReport.CreatedAt,
		UpdatedAt:    schoolReport.UpdatedAt,
		Status:       schoolReport.Status,
	}, nil
}

func (r *mutationResolver) RemoveSchoolReport(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorizaed")
	}
	collection := database.Db.Collection("school_reports")
	filter := bson.D{{"_id", id}}
	var schoolReport schoolReports.SchoolReport
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReport)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReport.Delete(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) ApproveSchoolReport(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorizaed")
	}
	collection := database.Db.Collection("school_reports")
	filter := bson.D{{"_id", id}}
	var schoolReport schoolReports.SchoolReport
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReport)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReport.Approve(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) RejectSchoolReport(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorizaed")
	}
	collection := database.Db.Collection("school_reports")
	filter := bson.D{{"_id", id}}
	var schoolReport schoolReports.SchoolReport
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReport)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReport.Reject(id)
	result = true
	return &result, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if IsAdmin {
		var users []*model.User
		collection := database.Db.Collection("users")
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var user *model.User
			err := cursor.Decode(&user)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
		return users, nil
	}
	return nil, fmt.Errorf("Unauthorized")
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if IsAdmin || UserID != "" {
		var user *model.User
		collection := database.Db.Collection("users")
		filter := bson.D{{"_id", UserID}}
		err := collection.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if IsAdmin || UserID == id {
		var user *model.User
		collection := database.Db.Collection("users")
		filter := bson.D{{"_id", id}}
		err := collection.FindOne(context.TODO(), filter).Decode(&user)
		if err != nil {
			return nil, err
		}
		return &model.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			IsAdmin:   user.IsAdmin,
			IsActive:  user.IsActive,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		}, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) ItemNotes(ctx context.Context, itemID string) ([]*model.Note, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if IsAdmin || UserID != "" {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"item_id", itemID}}
		findOptions := options.Find().SetSort(bson.D{{"created_at", -1}})
		cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var note *model.Note
			err := cursor.Decode(&note)
			if err != nil {
				return nil, err
			}
			notes = append(notes, &model.Note{
				ID:        note.ID,
				UserID:    note.UserID,
				ItemID:    note.ItemID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		return notes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) Note(ctx context.Context, id string) (*model.Note, error) {
	isAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	var note *model.Note
	noteCollection := database.Db.Collection("notes")
	noteFilter := bson.D{{"_id", id}}
	err := noteCollection.FindOne(context.TODO(), noteFilter).Decode(&note)
	if err != nil {
		return nil, err
	}
	if isAdmin || *note.UserID == UserID {
		return &model.Note{
			ID:        note.ID,
			UserID:    note.UserID,
			ItemID:    note.ItemID,
			Title:     note.Title,
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		}, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}

}

func (r *queryResolver) Log(ctx context.Context, id string) (*model.LogWithNotes, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	var LogWithNotes *model.LogWithNotes
	log := logs.Log{}
	logCollection := database.Db.Collection("logs")
	logFilter := bson.D{{"_id", id}}
	err := logCollection.FindOne(context.TODO(), logFilter).Decode(&log)
	if err != nil {
		return nil, err
	}
	logUserID := log.UserID
	println(log.FocusArea)
	println(log.UpdatedAt)
	if IsAdmin || logUserID == UserID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"item_id", id}}
		findOptions := options.Find().SetSort(bson.D{{"created_at", -1}})
		cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var note *model.Note
			err := cursor.Decode(&note)
			if err != nil {
				return nil, err
			}
			notes = append(notes, &model.Note{
				ID:        note.ID,
				UserID:    note.UserID,
				ItemID:    note.ItemID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		LogWithNotes = &model.LogWithNotes{
			Log: &model.Log{
				ID:           &log.ID,
				UserID:       &log.UserID,
				FocusArea:    log.FocusArea,
				Actions:      log.Actions,
				Successes:    log.Successes,
				Improvements: log.Improvements,
				NextSteps:    log.NextSteps,
				Status:       log.Status,
				CreatedAt:    log.CreatedAt,
				UpdatedAt:    log.UpdatedAt,
			},
			Notes: notes,
		}
		return LogWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) AllLogs(ctx context.Context) ([]*model.AllLogs, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	println("hit")
	if IsAdmin {
		filter := bson.D{}
		return utils.GetLogs(filter)
	} else if UserID != "" {
		filter := bson.D{{"user_id", UserID}}
		return utils.GetLogs(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) UserLogs(ctx context.Context, userID string) ([]*model.Log, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if IsAdmin {
		var logs []*model.Log
		collection := database.Db.Collection("logs")
		filter := bson.D{{"user_id", userID}}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var log *model.Log
			err := cursor.Decode(&log)
			if err != nil {
				return nil, err
			}
			logs = append(logs, &model.Log{
				ID:           log.ID,
				UserID:       log.UserID,
				FocusArea:    log.FocusArea,
				Actions:      log.Actions,
				Successes:    log.Successes,
				Improvements: log.Improvements,
				NextSteps:    log.NextSteps,
				Status:       log.Status,
				CreatedAt:    log.CreatedAt,
			})
		}
		return logs, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.EventWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var eventWithNotes *model.EventWithNotes
	var event *model.Event
	eventCollection := database.Db.Collection("events")
	eventFilter := bson.D{{"_id", id}}
	err := eventCollection.FindOne(context.TODO(), eventFilter).Decode(&event)
	if err != nil {
		return nil, err
	}
	eventLead := event.EventLead
	if isAdmin || eventLead == &userID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"item_id", id}}
		findOptions := options.Find().SetSort(bson.D{{"created_at", -1}})
		cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var note *model.Note
			err := cursor.Decode(&note)
			if err != nil {
				return nil, err
			}
			notes = append(notes, &model.Note{
				ID:        note.ID,
				UserID:    note.UserID,
				ItemID:    note.ItemID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		eventWithNotes = &model.EventWithNotes{
			Event: &model.Event{
				ID:                     event.ID,
				EventLead:              event.EventLead,
				Title:                  event.Title,
				Description:            event.Description,
				StartDate:              event.StartDate,
				SetUp:                  event.SetUp,
				CleanUp:                event.CleanUp,
				EndDate:                event.EndDate,
				GrantID:                event.GrantID,
				Public:                 event.Public,
				Rsvp:                   event.Rsvp,
				AnnualEvent:            event.AnnualEvent,
				NewEvent:               event.NewEvent,
				Volunteers:             event.Volunteers,
				Agenda:                 event.Agenda,
				TargetAudience:         event.TargetAudience,
				PartingGifts:           event.PartingGifts,
				MarketingMaterial:      event.MarketingMaterial,
				Supplies:               event.Supplies,
				SpecialOrders:          event.SpecialOrders,
				Performance:            event.Performance,
				Vendors:                event.Vendors,
				FoodAndBeverage:        event.FoodAndBeverage,
				Caterer:                event.Caterer,
				FoodHeadCount:          event.FoodHeadCount,
				EventTeam:              event.EventTeam,
				VolunteerList:          event.VolunteerList,
				Budget:                 event.Budget,
				AffiliatedOrganization: event.AffiliatedOrganization,
				EducationalGoals:       event.EducationalGoals,
				EducationalOutcomes:    event.EducationalOutcomes,
				GrantGoals:             event.GrantGoals,
				CreatedAt:              event.CreatedAt,
				UpdatedAt:              event.UpdatedAt,
				Status:                 event.Status,
			},
			Notes: notes,
		}
		return eventWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) EventSummary(ctx context.Context, id string) (*model.EventSummaryWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var eventSummaryWithNotes *model.EventSummaryWithNotes
	var eventSummary *model.EventSummary
	eventSummaryCollection := database.Db.Collection("event_summaries")
	eventSummaryFilter := bson.D{{"_id", id}}
	err := eventSummaryCollection.FindOne(context.TODO(), eventSummaryFilter).Decode(&eventSummary)
	if err != nil {
		return nil, err
	}
	summaryAuthor := eventSummary.UserID
	if isAdmin || summaryAuthor == userID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"item_id", id}}
		findOptions := options.Find().SetSort(bson.D{{"created_at", -1}})
		cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var note *model.Note
			err := cursor.Decode(&note)
			if err != nil {
				return nil, err
			}
			notes = append(notes, &model.Note{
				ID:        note.ID,
				UserID:    note.UserID,
				ItemID:    note.ItemID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		eventSummaryWithNotes = &model.EventSummaryWithNotes{
			EventSummary: &model.EventSummary{
				ID:            eventSummary.ID,
				UserID:        eventSummary.UserID,
				EventID:       eventSummary.EventID,
				AttendeeCount: eventSummary.AttendeeCount,
				Challenges:    eventSummary.Challenges,
				Successes:     eventSummary.Successes,
				Improvements:  eventSummary.Improvements,
				Status:        eventSummary.Status,
				CreatedAt:     eventSummary.CreatedAt,
				UpdatedAt:     eventSummary.UpdatedAt,
			},
			Notes: notes,
		}
		return eventSummaryWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) SchoolReport(ctx context.Context, id string) (*model.SchoolReportWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReportWithNotes *model.SchoolReportWithNotes
	var schoolReport *model.SchoolReport
	schoolReportCollection := database.Db.Collection("school_reports")
	schoolReportFilter := bson.D{{"_id", id}}
	err := schoolReportCollection.FindOne(context.TODO(), schoolReportFilter).Decode(&schoolReport)
	if err != nil {
		return nil, err
	}
	reportAuthor := schoolReport.UserID
	if isAdmin || reportAuthor == &userID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"item_id", id}}
		findOptions := options.Find().SetSort(bson.D{{"created_at", -1}})
		cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var note *model.Note
			err := cursor.Decode(&note)
			if err != nil {
				return nil, err
			}
			notes = append(notes, &model.Note{
				ID:        note.ID,
				UserID:    note.UserID,
				ItemID:    note.ItemID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		schoolReportWithNotes = &model.SchoolReportWithNotes{
			SchoolReport: &model.SchoolReport{
				ID:           schoolReport.ID,
				UserID:       schoolReport.UserID,
				Curriculum:   schoolReport.Curriculum,
				LessonPlan:   schoolReport.LessonPlan,
				School:       schoolReport.School,
				Topics:       schoolReport.Topics,
				StudentCount: schoolReport.StudentCount,
				StudentList:  schoolReport.StudentList,
				Challenges:   schoolReport.Challenges,
				Successes:    schoolReport.Successes,
				Improvements: schoolReport.Improvements,
				Status:       schoolReport.Status,
				CreatedAt:    schoolReport.CreatedAt,
				UpdatedAt:    schoolReport.UpdatedAt,
			},
			Notes: notes,
		}
		return schoolReportWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.AllEvents, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.D{}
		return utils.GetEvents(filter)
	} else if userID != "" {
		filter := bson.D{{"event_lead", userID}}
		return utils.GetEvents(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) EventSummaries(ctx context.Context) ([]*model.AllEventSummaries, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.D{}
		return utils.GetEventSummaries(filter)
	} else if userID != "" {
		filter := bson.D{{"user_id", userID}}
		return utils.GetEventSummaries(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) SchoolReports(ctx context.Context) ([]*model.AllSchoolReports, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.D{}
		return utils.GetSchoolReports(filter)
	} else if userID != "" {
		filter := bson.D{{"user_id", userID}}
		return utils.GetSchoolReports(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) Grants(ctx context.Context) ([]*model.Grant, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin && userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var grants []*model.Grant
	collection := database.Db.Collection("grants")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var grant *model.Grant
		err := cursor.Decode(&grant)
		if err != nil {
			return nil, err
		}
		grants = append(grants, &model.Grant{
			ID:          grant.ID,
			Name:        grant.Name,
			Description: grant.Description,
			Goals:       grant.Goals,
			Objectives:  grant.Objectives,
			StartDate:   grant.StartDate,
			AwardDate:   grant.AwardDate,
			EndDate:     grant.EndDate,
			AwardNumber: grant.AwardNumber,
			Budget:      grant.Budget,
			CreatedAt:   grant.CreatedAt,
			CreatedBy:   grant.CreatedBy,
			UpdatedAt:   grant.UpdatedAt,
			IsActive:    grant.IsActive,
		})
	}
	return grants, nil
}

func (r *queryResolver) Grant(ctx context.Context, id string) (*model.Grant, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin && userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var grant *model.Grant
	collection := database.Db.Collection("grants")
	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&grant)
	if err != nil {
		return nil, err
	}
	return grant, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin && userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var contacts []*model.Contact
	collection := database.Db.Collection("contacts")
	if !isAdmin {
		filter := bson.D{{"$or", bson.A{bson.D{{"type", "Student"}}, bson.D{{"type", "Parent"}}}}}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var contact *model.Contact
			err := cursor.Decode(&contact)
			if err != nil {
				return nil, err
			}
			contacts = append(contacts, &model.Contact{
				ID:        contact.ID,
				Name:      contact.Name,
				Type:      contact.Type,
				Email:     contact.Email,
				Phone:     contact.Phone,
				Notes:     contact.Notes,
				CreatedAt: contact.CreatedAt,
				CreatedBy: contact.CreatedBy,
				UpdatedAt: contact.UpdatedAt,
				IsActive:  contact.IsActive,
			})
		}
	} else {
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			return nil, err
		}
		for cursor.Next(context.TODO()) {
			var contact *model.Contact
			err := cursor.Decode(&contact)
			if err != nil {
				return nil, err
			}
			contacts = append(contacts, &model.Contact{
				ID:        contact.ID,
				Name:      contact.Name,
				Type:      contact.Type,
				Email:     contact.Email,
				Phone:     contact.Phone,
				Notes:     contact.Notes,
				CreatedAt: contact.CreatedAt,
				CreatedBy: contact.CreatedBy,
				UpdatedAt: contact.UpdatedAt,
				IsActive:  contact.IsActive,
				DeletedAt: contact.DeletedAt,
			})
		}
	}
	return contacts, nil
}

func (r *queryResolver) ContactInfo(ctx context.Context, id string) (*model.ContactInfo, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin && userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	contactsColl := database.Db.Collection("contacts")
	usersColl := database.Db.Collection("users")
	var contactInfo *model.ContactInfo
	var contact *model.Contact
	var user *model.User
	err := contactsColl.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&contact)
	if err != nil {
		return nil, err
	}
	err = usersColl.FindOne(context.TODO(), bson.D{{"_id", contact.CreatedBy}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	if !isAdmin && (*contact.Type != "Parent" && *contact.Type != "Student") {
		return nil, fmt.Errorf("Unauthorized")
	}
	contactInfo = &model.ContactInfo{
		Contact:        contact,
		ContactCreator: user,
	}
	return contactInfo, nil
}

func (r *queryResolver) UserEvents(ctx context.Context, userID string) ([]*model.Event, error) {
	// build out an admin search functionality
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserEventSummaries(ctx context.Context, userID string) ([]*model.EventSummary, error) {
	// build out an admin search functionality
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserSchoolReports(ctx context.Context, userID string) ([]*model.SchoolReport, error) {
	// build out an admin search functionality
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
