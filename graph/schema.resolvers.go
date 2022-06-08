package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "thomps9012/prevention_productivity/graph/generated"
	"thomps9012/prevention_productivity/graph/model"
	"thomps9012/prevention_productivity/internal/auth"
	database "thomps9012/prevention_productivity/internal/db"
	"thomps9012/prevention_productivity/internal/jwt"
	"thomps9012/prevention_productivity/internal/logs"
	"thomps9012/prevention_productivity/internal/notes"
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
	token, err := jwt.GenerateToken(user.Email, user.IsAdmin, user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser, id string) (*model.User, error) {
	isAdmin := auth.ForAdmin(ctx)
	if !isAdmin {
		return nil, fmt.Errorf("Unauthorized")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"_id", id}}
	var user users.User
	println("userid", id)
	println("first_name", updateUser.FirstName)
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
		return "", &users.WrongEmailOrPassword{}
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"email", login.Email}}
	var userDB users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&userDB)
	println(userDB.IsAdmin)
	println(userDB.ID)
	if !userDB.IsActive {
		return "", &users.UserNotActive{}
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateGrant(ctx context.Context, id string, updateGrant model.UpdateGrant) (*model.Grant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveGrant(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateContact(ctx context.Context, newContact model.NewContact) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id string, updateContact model.UpdateContact) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveContact(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, updateEvent model.UpdateEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveEvent(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveEvent(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectEvent(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEventSummary(ctx context.Context, newEventSummary model.NewEventSummary) (*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEventSummary(ctx context.Context, id string, updateEventSummary model.UpdateEventSummary) (*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveEventSummary(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveEventSummary(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectEventSummary(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSchoolReport(ctx context.Context, newSchoolReport model.NewSchoolReport) (*model.SchoolReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSchoolReport(ctx context.Context, id string, updateSchoolReport model.UpdateSchoolReport) (*model.SchoolReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveSchoolReport(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveSchoolReport(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectSchoolReport(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EventSummary(ctx context.Context, id string) (*model.EventSummaryWithNotes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SchoolReport(ctx context.Context, id string) (*model.SchoolReportWithNotes, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EventSummaries(ctx context.Context) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SchoolReports(ctx context.Context) ([]*model.SchoolReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Grants(ctx context.Context) ([]*model.Grant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserEvents(ctx context.Context, userID string) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserEventSummaries(ctx context.Context, userID string) ([]*model.EventSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserSchoolReports(ctx context.Context, userID string) ([]*model.SchoolReport, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
