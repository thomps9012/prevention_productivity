package jwt 

import (
	"time"
	"fmt"
	"log"
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("my_secret_key")
)

type Claims struct {
	userID string
	email string
	isAdmin bool
	jwt.StandardClaims
}

func GenerateToken(email string, isAdmin bool, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
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

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*Claims)
	return claims, nil
}