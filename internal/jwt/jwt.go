package jwt 

import (
	"time"
	"log"
	"os"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

var key = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	userID string
	email string
	isAdmin bool
	jwt.StandardClaims
}

func GenerateToken(email string, isAdmin bool, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	println(isAdmin)
	claims["email"] = email
	claims["isAdmin"] = isAdmin
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal("Error while signing the token")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var info = map[string]interface{}{
			"email": claims["email"].(string),
			"isAdmin": claims["isAdmin"],
			"userID": claims["userID"].(string),
		}
		return info, nil
	}
	return nil, fmt.Errorf("Token is invalid")
}
