package auth

import (
	"context"
	"net/http"
)

type key int

const (
	ContextUserKey key = iota
)

func PublicAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID string
		authCookie, err := r.Cookie("Authorization")

		if err != nil || authCookie == nil {
			//userID = uuid.NewV4().String()
			userID = "48c01326-079e-4092-a903-b994a5b62b21"
			JWTToken, err := BuildJWTString(userID)
			if err != nil {
				http.Error(w, "Can not build auth token", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: JWTToken})
		} else {
			userID = GetUserID(authCookie.Value)
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
