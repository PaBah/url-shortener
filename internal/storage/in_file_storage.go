package storage

import (
	"context"
	"database/sql"
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
	db    *sql.DB
}

func (fs *InFileStorage) Store(ctx context.Context, Data string) (ID string) {
	if len(fs.state) == 0 {
		fs.state = make(map[string]string)
	}

	ID = buildID(Data)
	fs.state[ID] = Data

	return
}

func (fs *InFileStorage) FindByID(ctx context.Context, ID string) (Data string, err error) {
	if fs.state == nil {
		fs.state = make(map[string]string)
	}

	Data, found := fs.state[ID]
	if !found {
		return Data, fmt.Errorf("no value with such ID")
	}

	return Data, nil
}

func (fs *InFileStorage) init(filePath string) {
	fs.file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)

	fs.state = map[string]string{}

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
	store.init(filePath)
	return store
}