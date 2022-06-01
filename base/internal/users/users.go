package users

import (
	database "prevention_productivity/base/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"context"
	"fmt"
	"strings"
	"github.com/google/uuid"
)

type User struct {
	ID	   string `json:"id" bson:"_id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsAdmin bool `json:"isAdmin"`
}


func (u *User) Create() {
	collection := database.Db.Collection("users")
	hashed, err := HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
	}
	u.Password = hashed
	u.IsAdmin = false
	u.ID = strings.Replace(uuid.New().String(), "-", "", -1)
	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
}

func (u *User) Authenticate() bool {
	collection := database.Db.Collection("users")
	var user User
	filter := bson.D{{"email", u.Email}}
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
