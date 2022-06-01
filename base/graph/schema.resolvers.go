package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"prevention_productivity/base/graph/generated"
	"prevention_productivity/base/graph/model"
	"prevention_productivity/base/internal/auth"
	database "prevention_productivity/base/internal/db"
	"prevention_productivity/base/internal/jwt"
	"prevention_productivity/base/internal/logs"
	"prevention_productivity/base/internal/users"

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
		return "", fmt.Errorf("Incorrect email or password")
	}
	collection := database.Db.Collection("users")
	filter := bson.D{{"email", login.Email}}
	var userDB users.User
	err := collection.FindOne(context.TODO(), filter).Decode(&userDB)
	if err != nil {
		return "", err
	}
	if !correct {
		return "", &users.WrongEmailOrPassword{}
	}
	token, err := jwt.GenerateToken(user.Email, user.IsAdmin, userDB.ID)
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
	log.Action = newLog.Action
	log.Create()
	return &model.Log{
		ID:       &log.ID,
		UserID:   &log.UserID,
		Action:   log.Action,
		CreatedAt: log.CreatedAt,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
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
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Logs(ctx context.Context) ([]*model.Log, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
