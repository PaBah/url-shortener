package models

import (
	"encoding/hex"
	"hash/fnv"
)

// ShortenURL - model entity of shortened URL
type ShortenURL struct {
	UUID        string `json:"uuid"`
	UserID      string `json:"user_id"`
	OriginalURL string `json:"original_URL"`
	DeletedFlag bool   `json:"is_deleted"`
}

// NewShortURL - create  instance of ShortenURL
func NewShortURL(originalURL string, userID string) ShortenURL {
	return ShortenURL{UUID: buildID(originalURL), OriginalURL: originalURL, UserID: userID}
}

func buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
