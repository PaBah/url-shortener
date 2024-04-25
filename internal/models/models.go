package models

import (
	"encoding/hex"
	"hash/fnv"
)

type ShortenURL struct {
	UUID        string `json:"uuid"`
	UserID      int    `json:"user_id"`
	OriginalURL string `json:"original_URL"`
}

func NewShortURL(originalURL string, userId int) ShortenURL {
	return ShortenURL{UUID: buildID(originalURL), OriginalURL: originalURL, UserID: userId}
}

func buildID(Value string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Value))
	ID = hex.EncodeToString(h.Sum(nil))
	return
}
