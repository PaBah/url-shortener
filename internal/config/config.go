package config

// Options - shortener server configurations
type Options struct {
	ServerAddress   string `json:"server_address"` // ServerAddress - address which system use to run shortener server
	BaseURL         string `json:"base_url"`       // BaseURL - host for shortened URLs
	LogsLevel       string // LogsLevel - level of logger
	FileStoragePath string `json:"file_storage_path"` // FileStoragePath - path to file where InFileStorage
	EnableHTTPS     bool   `json:"enable_https"`      // EnableHTTPS - flag to enable HTTPS server mode
	DatabaseDSN     string `json:"database_dsn"`      // DatabaseDSN - DSN path for DB connection
	TrustedSubnet   string `json:"trusted_subnet"`    // TrustedSubnet - CIDR address of allowed subnet
}
