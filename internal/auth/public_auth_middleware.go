package auth

import (
	"context"
	"net/http"
)

const (
	CONTEXT_USER_ID_KEY = "userID"
)

func PublicAuthorizationMiddleware(next http.Handler) http.Handler {
	knownUsers := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := -1
		authCookie, err := r.Cookie("Authorization")

		if err != nil || authCookie == nil {
			knownUsers++
			userID = knownUsers
			JWTToken, err := BuildJWTString(knownUsers)
			if err != nil {
				http.Error(w, "Can not build auth token", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: JWTToken})
		} else {
			userID = GetUserId(authCookie.Value)
		}

		ctx := context.WithValue(r.Context(), CONTEXT_USER_ID_KEY, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
