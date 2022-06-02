package auth

import (
	"context"
	"net/http"
	"prevention_productivity/internal/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type contextInfo struct {
	email string `json:"email"`
	isAdmin bool `json:"isAdmin"`
	id string `json:"id"`
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
			token, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			email := token["email"].(string)
			isAdmin := token["isAdmin"].(bool)
			userID := token["userID"].(string)
			contextInfo := &contextInfo{email, isAdmin, userID}
			ctx := context.WithValue(r.Context(), userCtxKey, contextInfo)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForUserID(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.id
}

func ForAdmin(ctx context.Context) bool {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.isAdmin
}