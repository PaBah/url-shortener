package storage

import (
	"context"
	"errors"

	"github.com/PaBah/url-shortener.git/internal/models"
)

var ErrConflict = errors.New("data conflict")

type Repository interface {
	Store(ctx context.Context, shortURL models.ShortenURL) (err error)
	FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error)
	GetAllUsers(ctx context.Context) (shortURLs []models.ShortenURL, err error)
	StoreBatch(ctx context.Context, shortURLsMap map[string]models.ShortenURL) (err error)
}
