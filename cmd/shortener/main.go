package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/PaBah/url-shortener.git/internal/tls"

	pb "github.com/PaBah/url-shortener.git/internal/gen/proto/shortener/v1"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
	options := &config.Options{}
	ParseFlags(options)

	if err := logger.Initialize(options.LogsLevel); err != nil {
		fmt.Printf("Logger can not be initialized %s", err)
		return
	}

	var store storage.Repository
	dbStore, err := storage.NewDBStorage(context.Background(), options.DatabaseDSN)
	if err != nil {
		logger.Log().Error("Database error with start", zap.Error(err))
		inFileStore := storage.NewInFileStorage(options.FileStoragePath)
		store = &inFileStore

		defer inFileStore.Close()
	} else {
		store = &dbStore
		defer dbStore.Close()
	}

	newServer := server.NewRouter(options, &store)
	newGRPCServer := server.NewShortenerServer(options, &store)

	logger.Log().Info("Start server on", zap.String("address", options.ServerAddress))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	go func() {
		listen, err := net.Listen("tcp", options.GRPCAddress)
		if err != nil {
			log.Fatal(err)
		}
		s := grpc.NewServer()
		pb.RegisterShortenerServiceServer(s, newGRPCServer)

		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if options.EnableHTTPS {
			const (
				certFilePath = "cert.pem" // certFilePath - path to TLS certificate
				keyFilePath  = "key.pem"  // keyFilePath - path to TLS key
			)
			err = tls.CreateTLSCert(certFilePath, keyFilePath)
			err = http.ListenAndServeTLS(options.ServerAddress, certFilePath, keyFilePath, newServer)
		} else {
			err = http.ListenAndServe(options.ServerAddress, newServer)
		}
		if err != nil {
			logger.Log().Fatal("Error starting server", zap.Error(err))
		}
	}()

	<-ctx.Done()
}
