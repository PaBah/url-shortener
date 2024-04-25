package auth

import (
	"context"
	"net/http"
)

func AuthorizedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("Authorization")

		if err != nil || authCookie == nil || authCookie.Value == "" {
			http.Error(w, "Unauthorized requests forbidden", http.StatusUnauthorized)
			return
		} else {

		}

		userID := GetUserId(authCookie.Value)
		ctx := context.WithValue(r.Context(), CONTEXT_USER_ID_KEY, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
