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
		}

		userID := GetUserID(authCookie.Value)
		ctx := context.WithValue(r.Context(), ContextUserKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
