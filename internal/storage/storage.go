package storage

import (
	"context"
	"errors"

	"github.com/PaBah/url-shortener.git/internal/models"
)

// ErrConflict - error when user tries to save already existing data
var ErrConflict = errors.New("data conflict")

// Repository - interface over Repository pattern for system storage
type Repository interface {
	Store(ctx context.Context, shortURL models.ShortenURL) (err error)
	FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error)
	GetAllUsers(ctx context.Context) (shortURLs []models.ShortenURL, err error)
	StoreBatch(ctx context.Context, shortURLsMap map[string]models.ShortenURL) (err error)
	AsyncCheckURLsUserID(usedID string, shortURL chan string) chan string
	DeleteShortURLs(ctx context.Context, shortURLs []string) (err error)
	GetStats(ctx context.Context) (urls int, users int, err error)
}
