package auth

import (
	"context"
	"net/http"
//	"strconv"
"go.mongodb.org/mongo-driver/bson"
	"prevention_productivity/base/internal/jwt"
	"prevention_productivity/base/internal/users"
	database "prevention_productivity/base/internal/db"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	username string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			email, err := jwt.ParseToken(tokenStr)
			println(email)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			collection := database.Db.Collection("users")
			var user users.User
			filter := bson.D{{"email", email}}
			err = collection.FindOne(context.TODO(), filter).Decode(&user)
			if err != nil {
				http.Error(w, "User not found", http.StatusForbidden)
				return
			}
			println(user.Email)
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}