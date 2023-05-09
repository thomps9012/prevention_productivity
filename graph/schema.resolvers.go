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
	"go.mongodb.org/mongo-driver/mongo"
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
	token, err := jwt.GenerateToken(user.Email, user.Admin, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser, id string) (*model.User, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if !isAdmin && id != userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	var user users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName
	user.Email = updateUser.Email
	user.Password = updateUser.Password
	user.Active = updateUser.Active
	if isAdmin {
		user.Admin = updateUser.Admin
	}
	count, err := collection.CountDocuments(context.TODO(), bson.D{{Key: "email", Value: updateUser.Email}})
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
		Admin:     user.Admin,
		Active:    user.Active,
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "email", Value: login.Email}}
	var userDB users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&userDB)
	println(userDB.Admin)
	println(userDB.ID)
	if !userDB.Active {
		return "User is not active", fmt.Errorf("User is not active")
	}
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(userDB.Email, userDB.Admin, userDB.ID)
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
	grant.CreatedBy = &userID
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
		CreatedBy:   *grant.CreatedBy,
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
		Active:      grant.Active,
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
		Active:      grant.Active,
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
	if contact.ID == "" {
		return &model.Contact{}, fmt.Errorf("Contact already exists")
	}
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
		Active:    contact.Active,
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
	if contact.Active {
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	log.DailyActivity = newLog.DailyActivity
	log.Positives = newLog.Positives
	log.Improvements = newLog.Improvements
	log.NextSteps = newLog.NextSteps
	log.Create()
	return &model.Log{
		ID:            &log.ID,
		UserID:        &log.UserID,
		DailyActivity: log.DailyActivity,
		Positives:     log.Positives,
		Improvements:  log.Improvements,
		NextSteps:     log.NextSteps,
		Status:        log.Status,
		CreatedAt:     log.CreatedAt,
	}, nil
}

func (r *mutationResolver) UpdateLog(ctx context.Context, id string, updateLog model.UpdateLog) (*model.Log, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: id}}
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
		log.DailyActivity = updateLog.DailyActivity
		log.Status = updateLog.Status
		log.Positives = updateLog.Positives
		log.Improvements = updateLog.Improvements
		log.NextSteps = updateLog.NextSteps
		log.Update(id)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
	return &model.Log{
		ID:            &log.ID,
		UserID:        &log.UserID,
		DailyActivity: log.DailyActivity,
		Positives:     log.Positives,
		Improvements:  log.Improvements,
		NextSteps:     log.NextSteps,
		Status:        log.Status,
		CreatedAt:     log.CreatedAt,
		UpdatedAt:     log.UpdatedAt,
	}, nil
}

func (r *mutationResolver) RemoveLog(ctx context.Context, id string) (bool, error) {
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	event.EventLead = &userID
	event.Coplanners = newEvent.Coplanners
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
		EventLead:   event.EventLead,
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
	filter := bson.D{{Key: "_id", Value: id}}
	var event events.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		return nil, err
	}
	if isAdmin || event.EventLead == &userID {
		event.Title = *updateEvent.Title
		event.Coplanners = updateEvent.Coplanners
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
		EventLead:   event.EventLead,
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	// possibly add in backend logic here for coplanner ids and event lead id here
	var eventSummary eventSummaries.EventSummary
	eventSummary.EventID = *newEventSummary.EventID
	eventSummary.UserID = userID
	eventSummary.Coplanners = newEventSummary.Coplanners
	eventSummary.AttendeeCount = newEventSummary.AttendeeCount
	eventSummary.Challenges = newEventSummary.Challenges
	eventSummary.Successes = newEventSummary.Successes
	eventSummary.Improvements = newEventSummary.Improvements
	eventSummary.Create()
	return &model.EventSummary{
		ID:            eventSummary.ID,
		EventID:       eventSummary.EventID,
		UserID:        eventSummary.UserID,
		Coplanners:    eventSummary.Coplanners,
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
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&eventSummary)
	if err != nil {
		return nil, err
	}
	if !isAdmin && eventSummary.UserID != userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	eventSummary.Coplanners = updateEventSummary.Coplanners
	eventSummary.AttendeeCount = updateEventSummary.AttendeeCount
	eventSummary.Challenges = updateEventSummary.Challenges
	eventSummary.Successes = updateEventSummary.Successes
	eventSummary.Improvements = updateEventSummary.Improvements
	eventSummary.Update(id)
	return &model.EventSummary{
		ID:            eventSummary.ID,
		EventID:       eventSummary.EventID,
		UserID:        eventSummary.UserID,
		Coplanners:    eventSummary.Coplanners,
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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
	filter := bson.D{{Key: "_id", Value: id}}
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

func (r *mutationResolver) CreateSchoolReportPlan(ctx context.Context, newSchoolReportPlan model.NewSchoolReportPlan) (*model.SchoolReportPlan, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var schoolReportPlan schoolReports.SchoolReportPlan
	schoolReportPlan.Curriculum = newSchoolReportPlan.Curriculum
	schoolReportPlan.Cofacilitators = newSchoolReportPlan.Cofacilitators
	schoolReportPlan.LessonTopics = newSchoolReportPlan.LessonTopics
	schoolReportPlan.UserID = &userID
	schoolReportPlan.School = newSchoolReportPlan.School
	schoolReportPlan.Create()
	return &model.SchoolReportPlan{
		ID:             &schoolReportPlan.ID,
		UserID:         schoolReportPlan.UserID,
		Cofacilitators: schoolReportPlan.Cofacilitators,
		Curriculum:     schoolReportPlan.Curriculum,
		LessonTopics:   schoolReportPlan.LessonTopics,
		School:         schoolReportPlan.School,
		CreatedAt:      schoolReportPlan.CreatedAt,
		UpdatedAt:      schoolReportPlan.UpdatedAt,
		Status:         schoolReportPlan.Status,
	}, nil
}

func (r *mutationResolver) UpdateSchoolReportPlan(ctx context.Context, id string, updateSchoolReportPlan model.UpdateSchoolReportPlan) (*model.SchoolReportPlan, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReportPlan schoolReports.SchoolReportPlan
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportPlan)
	if err != nil {
		return nil, err
	}
	if !isAdmin && schoolReportPlan.UserID != &userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	schoolReportPlan.Curriculum = updateSchoolReportPlan.Curriculum
	schoolReportPlan.Cofacilitators = updateSchoolReportPlan.Cofacilitators
	schoolReportPlan.LessonTopics = updateSchoolReportPlan.LessonTopics
	schoolReportPlan.School = updateSchoolReportPlan.School
	schoolReportPlan.Status = *updateSchoolReportPlan.Status
	schoolReportPlan.Update(id)
	return &model.SchoolReportPlan{
		ID:             &schoolReportPlan.ID,
		UserID:         schoolReportPlan.UserID,
		Cofacilitators: schoolReportPlan.Cofacilitators,
		Curriculum:     schoolReportPlan.Curriculum,
		LessonTopics:   schoolReportPlan.LessonTopics,
		School:         schoolReportPlan.School,
		CreatedAt:      schoolReportPlan.CreatedAt,
		UpdatedAt:      schoolReportPlan.UpdatedAt,
		Status:         schoolReportPlan.Status,
	}, nil
}

func (r *mutationResolver) RemoveSchoolReportPlan(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportPlan schoolReports.SchoolReportPlan
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportPlan)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportPlan.Delete(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) ApproveSchoolReportPlan(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportPlan schoolReports.SchoolReportPlan
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportPlan)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportPlan.Approve(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) RejectSchoolReportPlan(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportPlan schoolReports.SchoolReportPlan
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportPlan)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportPlan.Reject(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) CreateSchoolReportDebrief(ctx context.Context, newSchoolReportDebrief model.NewSchoolReportDebrief) (*model.SchoolReportDebrief, error) {
	userID := auth.ForUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	// possibly add in backend logic here for cofacilitator ids and lesson lead here
	var schoolReportDebrief schoolReports.SchoolReportDebrief
	schoolReportDebrief.LessonPlanID = newSchoolReportDebrief.LessonPlanID
	schoolReportDebrief.StudentCount = *newSchoolReportDebrief.StudentCount
	schoolReportDebrief.Positives = newSchoolReportDebrief.Positives
	schoolReportDebrief.Discussion = newSchoolReportDebrief.Discussion
	schoolReportDebrief.StudentList = newSchoolReportDebrief.StudentList
	schoolReportDebrief.ChallengesImprovements = newSchoolReportDebrief.ChallengesImprovements
	schoolReportDebrief.UserID = &userID
	schoolReportDebrief.Create()
	return &model.SchoolReportDebrief{
		ID:                     schoolReportDebrief.ID,
		UserID:                 *schoolReportDebrief.UserID,
		LessonPlanID:           schoolReportDebrief.LessonPlanID,
		StudentCount:           &schoolReportDebrief.StudentCount,
		StudentList:            schoolReportDebrief.StudentList,
		ChallengesImprovements: schoolReportDebrief.ChallengesImprovements,
		Positives:              schoolReportDebrief.Positives,
		Discussion:             schoolReportDebrief.Discussion,
		CreatedAt:              schoolReportDebrief.CreatedAt,
		UpdatedAt:              schoolReportDebrief.UpdatedAt,
		Status:                 schoolReportDebrief.Status,
	}, nil
}

func (r *mutationResolver) UpdateSchoolReportDebrief(ctx context.Context, id string, updateSchoolReportDebrief model.UpdateSchoolReportDebrief) (*model.SchoolReportDebrief, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReportPlan schoolReports.SchoolReportPlan
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportPlan)
	if err != nil {
		return nil, err
	}
	if !isAdmin && schoolReportPlan.UserID != &userID {
		return nil, fmt.Errorf("Unauthorized")
	}
	var schoolReportDebrief schoolReports.SchoolReportDebrief
	schoolReportDebrief.StudentCount = *updateSchoolReportDebrief.StudentCount
	schoolReportDebrief.Positives = updateSchoolReportDebrief.Positives
	schoolReportDebrief.Discussion = updateSchoolReportDebrief.Discussion
	schoolReportDebrief.StudentList = updateSchoolReportDebrief.StudentList
	schoolReportDebrief.ChallengesImprovements = updateSchoolReportDebrief.ChallengesImprovements
	schoolReportDebrief.Status = *updateSchoolReportDebrief.Status
	schoolReportDebrief.Update(id)
	return &model.SchoolReportDebrief{
		ID:                     schoolReportDebrief.ID,
		UserID:                 *schoolReportDebrief.UserID,
		LessonPlanID:           schoolReportDebrief.LessonPlanID,
		StudentCount:           &schoolReportDebrief.StudentCount,
		StudentList:            schoolReportDebrief.StudentList,
		ChallengesImprovements: schoolReportDebrief.ChallengesImprovements,
		Positives:              schoolReportDebrief.Positives,
		Discussion:             schoolReportDebrief.Discussion,
		CreatedAt:              schoolReportDebrief.CreatedAt,
		UpdatedAt:              schoolReportDebrief.UpdatedAt,
		Status:                 schoolReportDebrief.Status,
	}, nil
}

func (r *mutationResolver) RemoveSchoolReportDebrief(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportDebrief schoolReports.SchoolReportDebrief
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportDebrief)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportDebrief.Delete(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) ApproveSchoolReportDebrief(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportDebrief schoolReports.SchoolReportDebrief
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportDebrief)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportDebrief.Approve(id)
	result = true
	return &result, nil
}

func (r *mutationResolver) RejectSchoolReportDebrief(ctx context.Context, id string) (*bool, error) {
	isAdmin := auth.ForAdmin(ctx)
	var result bool
	if !isAdmin {
		result = false
		return &result, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "_id", Value: id}}
	var schoolReportDebrief schoolReports.SchoolReportDebrief
	err := collection.FindOne(context.TODO(), filter).Decode(&schoolReportDebrief)
	if err != nil {
		result = false
		return &result, err
	}
	schoolReportDebrief.Reject(id)
	result = true
	return &result, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if !IsAdmin && UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var users []*model.User
	collection := database.Db.Collection("users")
	var projection *options.FindOptions
	if IsAdmin {
		projection = options.Find().SetProjection(bson.D{{}})
	} else {
		projection = options.Find().SetProjection(bson.M{"id": 1, "first_name": 1, "last_name": 1, "active": 1})
	}
	cursor, err := collection.Find(context.TODO(), bson.D{}, projection)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if !IsAdmin && UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var user *model.User
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: UserID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if !IsAdmin && UserID != id {
		return nil, fmt.Errorf("Unauthorized")
	}
	var user *model.User
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) ItemNotes(ctx context.Context, itemID string) ([]*model.Note, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if !IsAdmin && UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var notes []*model.Note
	noteCollection := database.Db.Collection("notes")
	noteFilter := bson.D{{Key: "item_id", Value: itemID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := noteCollection.Find(context.TODO(), noteFilter, findOptions)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *queryResolver) Note(ctx context.Context, id string) (*model.Note, error) {
	isAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	var note *model.Note
	noteCollection := database.Db.Collection("notes")
	noteFilter := bson.D{{Key: "_id", Value: id}}
	err := noteCollection.FindOne(context.TODO(), noteFilter).Decode(&note)
	if err != nil {
		return nil, err
	}
	if !isAdmin && *note.UserID != UserID {
		return nil, fmt.Errorf("Unauthorized")
	}
	return note, nil
}

func (r *queryResolver) Log(ctx context.Context, id string) (*model.LogWithNotes, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	var LogWithNotes *model.LogWithNotes
	logCollection := database.Db.Collection("logs")
	logFilter := bson.D{{Key: "_id", Value: id}}
	notes_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "notes"}}}}
	pipeline := mongo.Pipeline{logFilter, notes_stage}
	cursor, err := logCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &LogWithNotes)
	if err != nil {
		return nil, err
	}
	if *LogWithNotes.Log.UserID != UserID && !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	return LogWithNotes, nil
}

func (r *queryResolver) AllLogs(ctx context.Context) ([]*model.AllLogs, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	if UserID == "" {
		return nil, fmt.Errorf("Unauthorized")
	}
	var filter bson.D
	if IsAdmin {
		filter = bson.D{{}}
	} else {
		filter = bson.D{{Key: "user_id", Value: UserID}}
	}
	logsCollection := database.Db.Collection("logs")
	sort_stage := bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "user"}, {
		Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$project", Value: bson.D{{Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}, {Key: "_id", Value: 1}}}},
		},
	}}}}
	note_stage := bson.D{{Key: "$count", Value: bson.D{{Key: "from", Value: "notes"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "item_id"}, {Key: "as", Value: "note_count"}}}}
	log_pipeline := mongo.Pipeline{filter, sort_stage, user_stage, note_stage}
	cursor, err := logsCollection.Aggregate(context.TODO(), log_pipeline)
	if err != nil {
		return nil, err
	}
	var allLogs []*model.AllLogs
	err = cursor.All(context.TODO(), &allLogs)
	if err != nil {
		return nil, err
	}
	return allLogs, nil
}

func (r *queryResolver) UserLogs(ctx context.Context, userID string) ([]*model.Log, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var logs []*model.Log
	collection := database.Db.Collection("logs")
	filter := bson.D{{Key: "user_id", Value: userID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.EventWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var eventWithNotes *model.EventWithNotes
	var event *model.Event
	eventCollection := database.Db.Collection("events")
	eventFilter := bson.D{{Key: "_id", Value: id}}
	err := eventCollection.FindOne(context.TODO(), eventFilter).Decode(&event)
	if err != nil {
		return nil, err
	}
	eventLead := event.EventLead
	eventCoplanners := event.Coplanners
	isCoplanner := utils.Exists(userID, eventCoplanners)
	if isAdmin || eventLead == &userID || isCoplanner {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{Key: "item_id", Value: id}}
		findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
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
				Coplanners:             event.Coplanners,
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
	eventSummaryFilter := bson.D{{Key: "_id", Value: id}}
	err := eventSummaryCollection.FindOne(context.TODO(), eventSummaryFilter).Decode(&eventSummary)
	if err != nil {
		return nil, err
	}
	summaryAuthor := eventSummary.UserID
	if isAdmin || summaryAuthor == userID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{Key: "item_id", Value: id}}
		findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
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
				Coplanners:    eventSummary.Coplanners,
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

func (r *queryResolver) SchoolReportPlan(ctx context.Context, id string) (*model.SchoolReportPlanWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReportPlanWithNotes *model.SchoolReportPlanWithNotes
	var schoolReportPlan *model.SchoolReportPlan
	schoolReportPlanCollection := database.Db.Collection("school_report_plans")
	schoolReportPlanFilter := bson.D{{Key: "_id", Value: id}}
	err := schoolReportPlanCollection.FindOne(context.TODO(), schoolReportPlanFilter).Decode(&schoolReportPlan)
	if err != nil {
		return nil, err
	}
	reportAuthor := schoolReportPlan.UserID
	cofacilitators := schoolReportPlan.Cofacilitators
	isCofacilitator := utils.Exists(userID, cofacilitators)
	fmt.Printf("\ntesting report author%v\n", *reportAuthor)
	fmt.Printf("\ntesting user id%v\n", userID)
	fmt.Printf("\ntesting report author == user id%v\n", *reportAuthor == userID)
	fmt.Printf("\ntesting is cofacilitator%v\n", isCofacilitator)
	fmt.Printf("\ntesting is admin%v\n", isAdmin)
	fmt.Printf("\ntesting condition%v\n", isAdmin || *reportAuthor == userID || isCofacilitator)
	if isAdmin || *reportAuthor == userID || isCofacilitator {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{Key: "item_id", Value: id}}
		findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
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
		schoolReportPlanWithNotes = &model.SchoolReportPlanWithNotes{
			SchoolReportPlan: &model.SchoolReportPlan{
				ID:             schoolReportPlan.ID,
				UserID:         schoolReportPlan.UserID,
				Cofacilitators: schoolReportPlan.Cofacilitators,
				Curriculum:     schoolReportPlan.Curriculum,
				School:         schoolReportPlan.School,
				LessonTopics:   schoolReportPlan.LessonTopics,
				Status:         schoolReportPlan.Status,
				CreatedAt:      schoolReportPlan.CreatedAt,
				UpdatedAt:      schoolReportPlan.UpdatedAt,
			},
			Notes: notes,
		}
		return schoolReportPlanWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) SchoolReportDebrief(ctx context.Context, id string) (*model.SchoolReportDebriefWithNotes, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	var schoolReportDebriefWithNotes *model.SchoolReportDebriefWithNotes
	var schoolReportDebrief *model.SchoolReportDebrief
	schoolReportDebriefCollection := database.Db.Collection("school_report_debriefs")
	schoolReportDebriefFilter := bson.D{{Key: "_id", Value: id}}
	err := schoolReportDebriefCollection.FindOne(context.TODO(), schoolReportDebriefFilter).Decode(&schoolReportDebrief)
	if err != nil {
		return nil, err
	}
	reportAuthor := schoolReportDebrief.UserID
	if isAdmin || reportAuthor == userID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{Key: "item_id", Value: id}}
		findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
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
		schoolReportDebriefWithNotes = &model.SchoolReportDebriefWithNotes{
			SchoolReportDebrief: &model.SchoolReportDebrief{
				ID:                     schoolReportDebrief.ID,
				UserID:                 schoolReportDebrief.UserID,
				LessonPlanID:           schoolReportDebrief.LessonPlanID,
				StudentCount:           schoolReportDebrief.StudentCount,
				StudentList:            schoolReportDebrief.StudentList,
				ChallengesImprovements: schoolReportDebrief.ChallengesImprovements,
				Positives:              schoolReportDebrief.Positives,
				Discussion:             schoolReportDebrief.Discussion,
				Status:                 schoolReportDebrief.Status,
				CreatedAt:              schoolReportDebrief.CreatedAt,
				UpdatedAt:              schoolReportDebrief.UpdatedAt,
			},
			Notes: notes,
		}
		return schoolReportDebriefWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.AllEvents, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.M{}
		return utils.GetEvents(filter)
	} else if userID != "" {
		filter := bson.M{"$or": bson.A{bson.M{"coplanners": userID}, bson.M{"event_lead": userID}}}
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
		// add in check for coplanners here
		filter := bson.D{{Key: "user_id", Value: userID}}
		return utils.GetEventSummaries(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) SchoolReportPlans(ctx context.Context) ([]*model.AllSchoolReportPlans, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.M{}
		return utils.GetSchoolReportPlans(filter)
	} else if userID != "" {
		filter := bson.M{"$or": bson.A{bson.M{"cofacilitators": userID}, bson.M{"user_id": userID}}}
		return utils.GetSchoolReportPlans(filter)
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) SchoolReportDebriefs(ctx context.Context) ([]*model.AllSchoolReportDebriefs, error) {
	isAdmin := auth.ForAdmin(ctx)
	userID := auth.ForUserID(ctx)
	if isAdmin {
		filter := bson.D{}
		return utils.GetSchoolReportDebriefs(filter)
	} else if userID != "" {
		filter := bson.D{{Key: "user_id", Value: userID}}
		return utils.GetSchoolReportDebriefs(filter)
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
			Active:      grant.Active,
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
	err := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&grant)
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
		filter := bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "type", Value: "Student"}}, bson.D{{Key: "type", Value: "Parent"}}}}}
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
				Active:    contact.Active,
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
				Active:    contact.Active,
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
	err := contactsColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&contact)
	if err != nil {
		return nil, err
	}
	err = usersColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: contact.CreatedBy}}).Decode(&user)
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
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var events []*model.Event
	collection := database.Db.Collection("events")
	filter := bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "event_lead", Value: userID}}, bson.D{{Key: "coplanners", Value: userID}}}}}
	findOptions := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var event *model.Event
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &model.Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			Status:      event.Status,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
		})
	}
	return events, nil
}

func (r *queryResolver) UserEventSummaries(ctx context.Context, userID string) ([]*model.EventSummary, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var events []*model.EventSummary
	collection := database.Db.Collection("event_summaries")
	filter := bson.D{{Key: "user_id", Value: userID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var event *model.EventSummary
		err := cursor.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, &model.EventSummary{
			ID:            event.ID,
			EventID:       event.EventID,
			AttendeeCount: event.AttendeeCount,
			Status:        event.Status,
			CreatedAt:     event.CreatedAt,
			UpdatedAt:     event.UpdatedAt,
		})
	}
	return events, nil
}

func (r *queryResolver) UserSchoolReportPlans(ctx context.Context, userID string) ([]*model.SchoolReportPlan, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var reports []*model.SchoolReportPlan
	collection := database.Db.Collection("school_report_plans")
	filter := bson.D{{Key: "user_id", Value: userID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var report *model.SchoolReportPlan
		err := cursor.Decode(&report)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &model.SchoolReportPlan{
			ID:           report.ID,
			UserID:       report.UserID,
			Curriculum:   report.Curriculum,
			School:       report.School,
			LessonTopics: report.LessonTopics,
			Status:       report.Status,
			CreatedAt:    report.CreatedAt,
			UpdatedAt:    report.UpdatedAt,
		})
	}
	return reports, nil
}

func (r *queryResolver) UserSchoolReportDebriefs(ctx context.Context, userID string) ([]*model.SchoolReportDebrief, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var reports []*model.SchoolReportDebrief
	collection := database.Db.Collection("school_report_debriefs")
	filter := bson.D{{Key: "user_id", Value: userID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var report *model.SchoolReportDebrief
		err := cursor.Decode(&report)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &model.SchoolReportDebrief{
			ID:           report.ID,
			UserID:       report.UserID,
			LessonPlanID: report.LessonPlanID,
			StudentCount: report.StudentCount,
			Status:       report.Status,
			CreatedAt:    report.CreatedAt,
			UpdatedAt:    report.UpdatedAt,
		})
	}
	return reports, nil
}

func (r *queryResolver) UserNotes(ctx context.Context, userID string) ([]*model.Note, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if !IsAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	var notes []*model.Note
	collection := database.Db.Collection("notes")
	filter := bson.D{{Key: "user_id", Value: userID}}
	findOptions := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(10)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
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
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		})
	}
	return notes, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
