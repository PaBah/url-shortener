package storage

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"os"

	"github.com/PaBah/url-shortener.git/internal/models"
)

type Repository interface {
	Store(string) string
	FindByID(string) (string, error)
}

type InFileStorage struct {
	state map[string]string
	file  *os.File
}

func (fs *InFileStorage) Store(Data string) (ID string) {
	if len(fs.state) == 0 {
		fs.state = make(map[string]string)
	}

	ID = fs.buildID(Data)
	prevStateLen := len(fs.state)
	fs.state[ID] = Data
	if prevStateLen < len(fs.state) {
		_ = fs.writeToFile(&models.ShortenURL{UUID: ID, OriginalURL: Data})
	}

	return
}

func (fs *InFileStorage) FindByID(ID string) (Data string, err error) {
	if fs.state == nil {
		fs.state = make(map[string]string)
	}

	Data, found := fs.state[ID]
	if !found {
		return Data, fmt.Errorf("no value with such ID")
	}

	return Data, nil
}

func (fs *InFileStorage) buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}

func (fs *InFileStorage) init(filePath string) {
	fs.file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)

	if fs.state == nil {
		fs.state = make(map[string]string)
	}

	decoder := json.NewDecoder(fs.file)
	shortURLRecord := &models.ShortenURL{}
	for {
		if err := decoder.Decode(&shortURLRecord); err == io.EOF {
			break
		}
		fs.state[shortURLRecord.UUID] = shortURLRecord.OriginalURL
	}
}

func (fs *InFileStorage) writeToFile(shortenURL *models.ShortenURL) error {
	writer := json.NewEncoder(fs.file)
	return writer.Encode(&shortenURL)
}

func (fs *InFileStorage) Close() error {
	return fs.file.Close()
}

func NewInFileStorage(filePath string) InFileStorage {
	store := InFileStorage{}
	store.init(filePath)
	return store
}
