package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/models"
	"go.uber.org/zap"
)

// InFileStorage - model of Repository storage on top of file
type InFileStorage struct {
	state map[string]models.ShortenURL
	file  *os.File
}

// Store - stores shortened URL to internal field
func (fs *InFileStorage) Store(ctx context.Context, shortURL models.ShortenURL) (err error) {
	_, duplicate := fs.state[shortURL.UUID]
	if duplicate {
		err = ErrConflict
	}
	fs.state[shortURL.UUID] = shortURL
	return
}

// FindByID - filter and returns shortened URL by short ID
func (fs *InFileStorage) FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error) {
	var found bool
	shortURL, found = fs.state[ID]
	if !found {
		err = fmt.Errorf("no value with such ID")
	}
	return
}

// GetAllUsers - returns all shortened URLs of the User from context
func (fs *InFileStorage) GetAllUsers(ctx context.Context) (shortURLs []models.ShortenURL, err error) {
	shortURLs = make([]models.ShortenURL, 0)
	for _, shortURL := range fs.state {
		if shortURL.UserID == ctx.Value(auth.ContextUserKey).(string) {
			shortURLs = append(shortURLs, shortURL)
		}
	}

	if len(shortURLs) == 0 {
		err = errors.New("user don't have any URLs")
	}

	return
}

// StoreBatch - stores batch of shortened URLs in internal field
func (fs *InFileStorage) StoreBatch(ctx context.Context, shortURLs map[string]models.ShortenURL) (err error) {
	for _, shortURL := range shortURLs {
		fs.state[shortURL.UUID] = shortURL
	}
	return
}

// AsyncCheckURLsUserID - async checking if URL belongs to the User from context
func (fs *InFileStorage) AsyncCheckURLsUserID(userID string, shortURLCh chan string) chan string {
	addRes := make(chan string)
	go func() {
		defer close(addRes)

		for data := range shortURLCh {

			shortURL, err := fs.FindByID(context.Background(), data)
			var result string
			if err == nil && shortURL.UserID == userID {
				result = shortURL.UUID
			}

			addRes <- result
		}
	}()
	return addRes
}

// DeleteShortURLs - delete shortened URLs from Data Base
func (fs *InFileStorage) DeleteShortURLs(ctx context.Context, shortURLs []string) (err error) {
	for _, shortURL := range shortURLs {
		shortenedURL := fs.state[shortURL]
		shortenedURL.DeletedFlag = true
		fs.state[shortURL] = shortenedURL
	}
	return
}

func (fs *InFileStorage) initialize(filePath string) {
	fs.file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)

	fs.state = make(map[string]models.ShortenURL)

	decoder := json.NewDecoder(fs.file)
	shortURLRecord := models.ShortenURL{}
	for {
		if err := decoder.Decode(&shortURLRecord); err != nil {
			break
		}
		fs.state[shortURLRecord.UUID] = shortURLRecord
	}
}

func (fs *InFileStorage) writeBackup() error {
	writer := json.NewEncoder(fs.file)
	for _, shortenURL := range fs.state {
		err := writer.Encode(&shortenURL)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close - close file
func (fs *InFileStorage) Close() error {
	err := fs.writeBackup()
	if err != nil {
		logger.Log().Error("can not write backup file", zap.Error(err))
	}
	return fs.file.Close()
}

// NewInFileStorage - create instance of InFileStorage
func NewInFileStorage(filePath string) InFileStorage {
	store := InFileStorage{}
	store.initialize(filePath)
	return store
}
