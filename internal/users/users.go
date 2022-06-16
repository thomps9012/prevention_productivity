package users

import (
	database "thomps9012/prevention_productivity/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"context"
	"fmt"
	// "strings"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID	   string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"first_name"`
	LastName string `json:"lastName" bson:"last_name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	IsAdmin bool `json:"isAdmin"`
	IsActive bool `json:"isActive"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
}


func (u *User) Create() {
	collection := database.Db.Collection("users")
	filter := bson.D{{"email", u.Email}}
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		u.ID = uuid.New().String()
		u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		u.IsActive = true
		u.IsAdmin = false
		hashed, hashErr := HashPassword(u.Password)
		u.Password = hashed
		if hashErr != nil {
			fmt.Println(hashErr)
		}
		_, err := collection.InsertOne(context.TODO(), u)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Errorf("user already exists")
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

func (u *User) Update(id string) {
	collection := database.Db.Collection("users")
	filter := bson.D{{"_id", id}}
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	println("update pw", u.Password)
	if(len(u.Password) <= 0) {
		println("not updating pw")
		fmt.Printf("%v\n", u)
		result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.M{"first_name": u.FirstName, "last_name": u.LastName, "email": u.Email, "updated_at": u.UpdatedAt}}})
		if err != nil {
			panic(err)
		}
		println(result.ModifiedCount)
		} else {
			
			hashed, hashErr := HashPassword(u.Password)
			u.Password = hashed
			println("update pw", u.Password)
			fmt.Printf("%v\n", u)
			if hashErr != nil {
				fmt.Println(hashErr)
			}
			result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", u}})
			if err != nil {
				panic(err)
			}
		println(result.ModifiedCount)
	}
}

func (u *User) Delete() {
	collection := database.Db.Collection("users")
	filter := bson.D{{"email", u.Email}}
	u.IsActive = false
	_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", u}})
	if err != nil {
		panic(err)
	}
}
