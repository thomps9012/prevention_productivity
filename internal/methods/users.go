package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/internal/auth"
	database "thomps9012/prevention_productivity/internal/db"
	"thomps9012/prevention_productivity/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hash), err
}
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		println(err.Error())
	}
	return err == nil
}
func CreateUser(new_user models.NewUser) (*models.LoginRes, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: new_user.Email}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("account already associated with that email")
	}
	hashed, hashErr := HashPassword(new_user.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	user := models.User{
		ID:        uuid.New().String(),
		Email:     new_user.Email,
		FirstName: new_user.FirstName,
		LastName:  new_user.LastName,
		Password:  hashed,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: bson.TypeNull.String(),
		DeletedAt: bson.TypeNull.String(),
		Active:    true,
		Admin:     false,
	}
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert user")
	}
	token, err := auth.GenerateToken(user.Email, user.Admin, user.ID)
	if err != nil {
		return nil, err
	}
	return &models.LoginRes{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Active:    user.Active,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}, nil
}
func LoginUser(login models.LoginInput) (*models.LoginRes, error) {
	collection := database.Db.Collection("users")
	var user models.User
	filter := bson.D{{Key: "email", Value: login.Email}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	match := CheckPasswordHash(user.Password, login.Password)
	if !match {
		return nil, errors.New("incorrect login credentials")
	}
	token, err := auth.GenerateToken(user.Email, user.Admin, user.ID)
	if err != nil {
		return nil, err
	}
	return &models.LoginRes{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Active:    user.Active,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}, nil
}
func UpdateUser(update models.UpdateUser) (*models.UserUpdateRes, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: update.ID}}
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	var update_args bson.D
	if len(update.Password) <= 0 {
		update_args = bson.D{{Key: "$set", Value: bson.M{"first_name": update.FirstName, "last_name": update.LastName, "email": update.Email, "updated_at": updated_at, "active": update.Active, "admin": update.Admin}}}
	} else {
		hashed, hashErr := HashPassword(update.Password)
		if hashErr != nil {
			return nil, hashErr
		}
		update_args = bson.D{{Key: "$set", Value: bson.M{"password": hashed, "first_name": update.FirstName, "last_name": update.LastName, "email": update.Email, "updated_at": updated_at, "active": update.Active, "admin": update.Admin}}}
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var u models.UserUpdateRes
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func DeleteUser(user_id string) (*models.UserUpdateRes, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	update := bson.D{{Key: "active", Value: false}, {Key: "deleted_at", Value: now}, {Key: "updated_at", Value: now}}
	var u models.UserUpdateRes
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
