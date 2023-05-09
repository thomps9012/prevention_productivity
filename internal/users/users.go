package users

import (
	"context"
	"errors"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"first_name"`
	LastName  string `json:"lastName" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Admin     bool
	Active    bool
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}

func (u *User) Create() (*User, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: u.Email}}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("account already associated with that email")
	}
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.Active = true
	u.Admin = false
	hashed, hashErr := HashPassword(u.Password)
	u.Password = hashed
	if hashErr != nil {
		return nil, hashErr
	}
	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Authenticate() bool {
	collection := database.Db.Collection("users")
	var user User
	filter := bson.D{{Key: "email", Value: u.Email}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false
	}
	match := CheckPasswordHash(user.Password, u.Password)
	return match
}

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

func (u *User) Update(id string) (*User, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	var update bson.D
	if len(u.Password) <= 0 {
		update = bson.D{{Key: "$set", Value: bson.M{"first_name": u.FirstName, "last_name": u.LastName, "email": u.Email, "updated_at": u.UpdatedAt, "active": u.Active, "admin": u.Admin}}}
	} else {
		hashed, hashErr := HashPassword(u.Password)
		if hashErr != nil {
			return nil, hashErr
		}
		u.Password = hashed
		update = bson.D{{Key: "$set", Value: u}}
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Delete() (*User, error) {
	collection := database.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: u.Email}}
	u.Active = false
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$set", Value: u}}, &opts).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
