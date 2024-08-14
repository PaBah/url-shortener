package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"go.uber.org/zap"
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
	_, b, _, _ := runtime.Caller(0)
	configPath := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(b))), "config.json")

	var specified bool
	var serverAddress, baseURL, logsLevel, fileStoragePath, databaseDSN, enableHTTPS, configFilePath, trustedSubnet string
	var gRPCAddress string

	flag.StringVar(&configFilePath, "c", configPath, "path to config file")
	flag.StringVar(&options.ServerAddress, "a", ":8080", "host:port on which server run")
	flag.StringVar(&options.GRPCAddress, "g", ":3200", "host:port on which gRPC run")
	flag.StringVar(&options.BaseURL, "b", "http://localhost:8080", "URL for of shortened URLs hosting")
	flag.StringVar(&options.DatabaseDSN, "d", "host=localhost user=paulbahush dbname=urlshortener password=", "database DSN address")
	flag.StringVar(&options.LogsLevel, "l", "info", "logs level")
	flag.StringVar(&options.FileStoragePath, "f", "/tmp/short-url-db.json", "path to file.json with file storage data")
	flag.StringVar(&options.TrustedSubnet, "t", "", "CIDR address of allowed subnet")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "enable-https")
	flag.Parse()

	var fileConfig config.Options
	file, err := os.Open(configFilePath)
	if err == nil {
		err = json.NewDecoder(file).Decode(&fileConfig)
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				logger.Log().Error("can not close file", zap.Error(err))
			}
		}(file)
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
			if !isFlagPassed("t") {
				options.TrustedSubnet = fileConfig.TrustedSubnet
			}
			if !isFlagPassed("g") {
				options.GRPCAddress = fileConfig.GRPCAddress
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

	gRPCAddress, specified = os.LookupEnv("GRPC_ADDRESS")
	if specified {
		options.GRPCAddress = gRPCAddress
	}

	trustedSubnet, specified = os.LookupEnv("TRUSTED_SUBNET")
	if specified {
		options.TrustedSubnet = trustedSubnet
	}

	enableHTTPS, specified = os.LookupEnv("ENABLE_HTTPS")
	if specified {
		options.EnableHTTPS, _ = strconv.ParseBool(enableHTTPS)
	}

}
