package main

import (
	"os"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name          string
		expectedValue []string
		envValues     []string
	}{
		{
			name:          "got from ENV",
			expectedValue: []string{":8888", "https://yandex.ru", "info", "/tmp/short-url-db.json", "test", ":3200", "0.0.0.0/0"},
			envValues:     []string{":8888", "https://yandex.ru", "info", "/tmp/short-url-db.json", "test", ":3200", "0.0.0.0/0", "true"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := &config.Options{}
			if tt.envValues != nil {
				os.Setenv("SERVER_ADDRESS", tt.envValues[0])
				os.Setenv("BASE_URL", tt.envValues[1])
				os.Setenv("LOG_LEVEL", tt.envValues[2])
				os.Setenv("FILE_STORAGE_PATH", tt.envValues[3])
				os.Setenv("DATABASE_DSN", tt.envValues[4])
				os.Setenv("GRPC_ADDRESS", tt.envValues[5])
				os.Setenv("TRUSTED_SUBNET", tt.envValues[6])
				os.Setenv("ENABLE_HTTPS", tt.envValues[7])
			}
			ParseFlags(options)
			assert.Equal(t, options.ServerAddress, tt.expectedValue[0], "Правльно распаршеный SERVER_ADDRESS")
			assert.Equal(t, options.BaseURL, tt.expectedValue[1], "Правльно распаршеный BASE_URL")
			assert.Equal(t, options.LogsLevel, tt.expectedValue[2], "Правльно распаршеный LOG_LEVEL")
			assert.Equal(t, options.FileStoragePath, tt.expectedValue[3], "Правльно распаршеный FILE_STORAGE_PATH")
			assert.Equal(t, options.DatabaseDSN, tt.expectedValue[4], "Правльно распаршеный DATABASE_DSN")
			assert.Equal(t, options.GRPCAddress, tt.expectedValue[5], "Правльно распаршеный DATABASE_DSN")
			assert.Equal(t, options.TrustedSubnet, tt.expectedValue[6], "Правльно распаршеный DATABASE_DSN")
			assert.Equal(t, options.EnableHTTPS, true, "Правльно распаршеный DATABASE_DSN")
		})
	}
}
