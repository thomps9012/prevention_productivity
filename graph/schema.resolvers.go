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

	"go.mongodb.org/mongo-driver/bson"
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
		note.Update(id)
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
			users = append(users, &model.User{
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Username:  user.Username,
				IsAdmin:   user.IsAdmin,
			})
		}
		return users, nil
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
	if IsAdmin || logUserID == UserID {
		var notes []*model.Note
		noteCollection := database.Db.Collection("notes")
		noteFilter := bson.D{{"itemid", id}}
		cursor, err := noteCollection.Find(context.TODO(), noteFilter)
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
				ItemID:    note.ItemID,
				UserID:    note.UserID,
				Title:     note.Title,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
			})
		}
		LogWithNotes = &model.LogWithNotes{
			Log: &model.Log{
				ID:        &log.ID,
				UserID:    &log.UserID,
				FocusArea: log.FocusArea,
				Actions:   log.Actions,
				Successes: log.Successes,
				Improvements: log.Improvements,
				NextSteps: log.NextSteps,
				Status:    log.Status,
				CreatedAt: log.CreatedAt,
				UpdatedAt: log.UpdatedAt,
			},
			Notes: notes,
		}
		return LogWithNotes, nil
	} else {
		return nil, fmt.Errorf("Unauthorized")
	}
}

func (r *queryResolver) AllLogs(ctx context.Context) ([]*model.Log, error) {
	IsAdmin := auth.ForAdmin(ctx)
	UserID := auth.ForUserID(ctx)
	var filter bson.D
	if IsAdmin {
		filter = bson.D{}
	} else {
		filter = bson.D{{"user_id", UserID}}
	}
	collection := database.Db.Collection("logs")
	var logs []*model.Log
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
}

func (r *queryResolver) UserLogs(ctx context.Context, userID string) ([]*model.Log, error) {
	IsAdmin := auth.ForAdmin(ctx)
	if IsAdmin {
		var logs []*model.Log
		collection := database.Db.Collection("logs")
		filter := bson.D{{"userid", userID}}
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

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
