package main

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"

	"github.com/PaBah/url-shortener.git/internal/config"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// ParseFlags - initializer system configuration
func ParseFlags(options *config.Options) {
	var specified bool
	var serverAddress, baseURL, logsLevel, fileStoragePath, databaseDSN, enableHTTPS, configFilePath string

	flag.StringVar(&configFilePath, "c", "", "path to config file")
	flag.StringVar(&options.ServerAddress, "a", ":8080", "host:port on which server run")
	flag.StringVar(&options.BaseURL, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.StringVar(&options.DatabaseDSN, "d", "host=localhost user=paulbahush dbname=urlshortener password=", "database DSN address")
	flag.StringVar(&options.LogsLevel, "l", "info", "logs level")
	flag.StringVar(&options.FileStoragePath, "f", "/tmp/short-url-db.json", "path to file.json with file storage data")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "enable-https")
	flag.Parse()

	var fileConfig config.Options
	if configFilePath != "" {
		file, err := os.Open(configFilePath)
		if err == nil {
			err = json.NewDecoder(file).Decode(&fileConfig)
			if err == nil {
				if !isFlagPassed("a") {
					options.ServerAddress = fileConfig.ServerAddress
				}
				if !isFlagPassed("b") {
					options.BaseURL = fileConfig.BaseURL
				}
				if !isFlagPassed("d") {
					options.DatabaseDSN = fileConfig.DatabaseDSN
				}
				if !isFlagPassed("f") {
					options.FileStoragePath = fileConfig.FileStoragePath
				}
				if !isFlagPassed("s") {
					options.EnableHTTPS = fileConfig.EnableHTTPS
				}
			}
		}
		err = file.Close()
		if err != nil {
			return
		}
	}

	serverAddress, specified = os.LookupEnv("SERVER_ADDRESS")
	if specified {
		options.ServerAddress = serverAddress
	}

	baseURL, specified = os.LookupEnv("BASE_URL")
	if specified {
		options.BaseURL = baseURL
	}

	logsLevel, specified = os.LookupEnv("LOG_LEVEL")
	if specified {
		options.LogsLevel = logsLevel
	}

	fileStoragePath, specified = os.LookupEnv("FILE_STORAGE_PATH")
	if specified {
		options.FileStoragePath = fileStoragePath
	}

	databaseDSN, specified = os.LookupEnv("DATABASE_DSN")
	if specified {
		options.DatabaseDSN = databaseDSN
	}

	enableHTTPS, specified = os.LookupEnv("ENABLE_HTTPS")
	if specified {
		options.EnableHTTPS, _ = strconv.ParseBool(enableHTTPS)
	}
}
