package config

// Options - shortener server configurations
type Options struct {
	ServerAddress   string // ServerAddress - address which system use to run shortener server
	BaseURL         string // BaseURL - host for shortened URLs
	LogsLevel       string // LogsLevel - level of logger
	FileStoragePath string // FileStoragePath - path to file where InFileStorage
	DatabaseDSN     string // DatabaseDSN - DSN path for DB connection
}
