package models

type ShortenURL struct {
	UUID        string `json:"uuid"`
	OriginalURL string `json:"original_URL"`
}
