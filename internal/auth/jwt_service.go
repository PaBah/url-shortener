package auth

import (
	"fmt"
	"time"

	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// Claims - present JWT claims (customised with UserID)
type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

// Parameters for JWT tokens generation/parsing
const (
	// TokenExp - JWT token expiration time in microseconds
	TokenExp = time.Hour * 3

	// SecretKey - key for JWT encryption
	SecretKey = "supersecretkey"
)

// BuildJWTString - generate JWT string from UserID
func BuildJWTString(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserID - parse JWT string and return UserID
func GetUserID(tokenString string) string {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		return ""
	}

	if !token.Valid {
		logger.Log().Error("Token is not valid", zap.String("token", token.Raw))
		return ""
	}

	return claims.UserID
}
