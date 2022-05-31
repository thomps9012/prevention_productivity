package users

import (
	database "prevention_productivity/base/internal/db"
	"golang.org/x/crypto/bcrypt"
	"context"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin bool `json:"isAdmin"`
}


func (u *User) Create()  {
	collection := database.Db.Collection("users")
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		panic(err)
	}
	u.Password = hashedPassword
	u.IsAdmin = false
	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}