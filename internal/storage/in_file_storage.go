package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/models"
	"go.uber.org/zap"
)

type InFileStorage struct {
	state map[string]string
	file  *os.File
}

func (fs *InFileStorage) Store(ctx context.Context, shortURL models.ShortenURL) (err error) {
	_, duplicate := fs.state[shortURL.UUID]
	if duplicate {
		err = ErrConflict
	}
	fs.state[shortURL.UUID] = shortURL.OriginalURL
	return
}

func (fs *InFileStorage) FindByID(ctx context.Context, ID string) (shortURL models.ShortenURL, err error) {
	var (
		found bool
		URL   string
	)
	URL, found = fs.state[ID]
	shortURL = models.NewShortURL(URL)
	if !found {
		err = fmt.Errorf("no value with such ID")
	}
	return
}

func (fs *InFileStorage) StoreBatch(ctx context.Context, shortURLs map[string]models.ShortenURL) (err error) {
	for _, shortURL := range shortURLs {
		fs.state[shortURL.UUID] = shortURL.OriginalURL
	}
	return
}

func (fs *InFileStorage) initialize(filePath string) {
	fs.file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)

	fs.state = make(map[string]string)

	decoder := json.NewDecoder(fs.file)
	shortURLRecord := &models.ShortenURL{}
	for {
		if err := decoder.Decode(&shortURLRecord); err != nil {

			break
		}
		fs.state[shortURLRecord.UUID] = shortURLRecord.OriginalURL
	}
}

func (fs *InFileStorage) writeBackup() error {
	writer := json.NewEncoder(fs.file)
	for k, v := range fs.state {
		shortenURL := models.ShortenURL{
			UUID:        k,
			OriginalURL: v,
		}
		err := writer.Encode(&shortenURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fs *InFileStorage) Close() error {
	err := fs.writeBackup()
	if err != nil {
		logger.Log().Error("can not write backup file", zap.Error(err))
	}
	return fs.file.Close()
}

func NewInFileStorage(filePath string) InFileStorage {
	store := InFileStorage{}
	store.initialize(filePath)
	return store
}
