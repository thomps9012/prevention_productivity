package jwt 

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("my_secret_key")
)

type Claims struct {
	username string
	isAdmin bool
}

func GenerateToken(username string, isAdmin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return &Claims{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &Claims{}, fmt.Errorf("invalid token")
	}
	username := claims["username"].(string)
	isAdmin := claims["isAdmin"].(bool)
	return &Claims{username, isAdmin}, nil
}

