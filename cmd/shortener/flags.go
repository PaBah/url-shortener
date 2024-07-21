package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	//flag.StringVar(&configFilePath, "c", "/Users/paulbahush/projects/yp/url-shortener/config.json", "path to config file")
	flag.StringVar(&configFilePath, "c", "", "path to config file")
	flag.StringVar(&options.ServerAddress, "a", "", "host:port on which server run")
	flag.StringVar(&options.BaseURL, "b", "", "URL for of shortened URLs hosting")
	flag.StringVar(&options.DatabaseDSN, "d", "", "database DSN address")
	flag.StringVar(&options.LogsLevel, "l", "info", "logs level")
	flag.StringVar(&options.FileStoragePath, "f", "", "path to file.json with file storage data")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "enable-https")
	flag.Parse()

	if configFilePath != "" {
		defaultConfig := config.Options{
			ServerAddress:   ":8080",
			BaseURL:         "http://localhost:8080",
			LogsLevel:       "info",
			FileStoragePath: "/tmp/short-url-db.json",
			EnableHTTPS:     false,
			DatabaseDSN:     "host=localhost user=paulbahush dbname=urlshortener password=",
		}
		file, err := os.Open(configFilePath)
		defer file.Close()
		if err == nil {
			err = json.NewDecoder(file).Decode(&defaultConfig)
			fmt.Println(defaultConfig)
			if err == nil {
				if !isFlagPassed("a") {
					options.ServerAddress = defaultConfig.ServerAddress
				}
				if !isFlagPassed("b") {
					options.BaseURL = defaultConfig.BaseURL
				}
				if !isFlagPassed("d") {
					options.DatabaseDSN = defaultConfig.DatabaseDSN
				}
				if !isFlagPassed("f") {
					options.FileStoragePath = defaultConfig.FileStoragePath
				}
				if !isFlagPassed("s") {
					options.EnableHTTPS = defaultConfig.EnableHTTPS
				}
			}
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
