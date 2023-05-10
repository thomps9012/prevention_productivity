package auth

import (
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type contextInfo struct {
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	UserID string `json:"user_id"`
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
			token, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			email := token["email"].(string)
			admin := token["admin"].(bool)
			user_id := token["user_id"].(string)
			contextInfo := &contextInfo{email, admin, user_id}
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
	return raw.UserID
}

func ForAdmin(ctx context.Context) bool {
	raw, _ := ctx.Value(userCtxKey).(*contextInfo)
	return raw.Admin
}
