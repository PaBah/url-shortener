package server

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/PaBah/url-shortener.git/internal/async"
	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"

	pb "github.com/PaBah/url-shortener.git/internal/gen/proto/shortener/v1"
)

// ShortenerServer shortener gRPC server
type ShortenerServer struct {
	pb.UnimplementedShortenerServer

	options *config.Options
	storage storage.Repository
}

// Short - handler for shortening URL
func (s *ShortenerServer) Short(ctx context.Context, in *pb.ShortRequest) (*pb.ShortResponse, error) {
	response := &pb.ShortResponse{}
	shortURL := models.NewShortURL(in.Url, in.UserId)
	err := s.storage.Store(ctx, shortURL)

	response.Result = shortURL.UUID
	if errors.Is(err, storage.ErrConflict) {
		return response, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return response, err
}

// Expand - handler for list user's shortened URLs
func (s *ShortenerServer) Expand(ctx context.Context, in *pb.ExpandRequest) (*pb.ExpandResponse, error) {
	response := &pb.ExpandResponse{}

	shortenURL, _ := s.storage.FindByID(ctx, in.ShortId)
	if shortenURL.DeletedFlag {
		return response, status.Errorf(codes.InvalidArgument, "shorten URL already expanded")
	}
	response.Url = shortenURL.OriginalURL
	return response, nil
}

// Delete - handler for delete short URLs
func (s *ShortenerServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	response := new(emptypb.Empty)

	inputCh := async.BulkDeletionDataGenerator(in.Id)

	channels := async.DeletionFanOut(in.UserId, s.storage, inputCh)
	addResultCh := async.DeletionFanIn(channels...)
	async.Delete(s.storage, addResultCh)

	return response, nil
}

// GetUserBucket - handler for list of short URLs of authorized user
func (s *ShortenerServer) GetUserBucket(ctx context.Context, in *pb.GetUserBucketRequest) (*pb.GetUserBucketResponse, error) {
	response := &pb.GetUserBucketResponse{}
	shortURLs, err := s.storage.GetAllUsers(context.WithValue(ctx, auth.ContextUserKey, in.UserId))
	if err != nil {
		return response, status.Errorf(codes.InvalidArgument, err.Error())
	}

	for _, shortURL := range shortURLs {
		response.Data = append(response.Data, &pb.OriginalAndShort{
			OriginalUrl: shortURL.OriginalURL,
			ShortUrl:    fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL.UUID),
		})
	}

	return response, nil
}

// ShortBatch - handler for creation of list of short URLs
func (s *ShortenerServer) ShortBatch(ctx context.Context, in *pb.ShortBatchRequest) (*pb.ShortBatchResponse, error) {
	response := &pb.ShortBatchResponse{}

	shortURLsMap := make(map[string]models.ShortenURL, len(in.Original))
	for _, batchRequest := range in.Original {
		shortURL := models.NewShortURL(batchRequest.OriginalUrl, in.UserId)
		shortURLsMap[batchRequest.CorrelationId] = shortURL
	}

	err := s.storage.StoreBatch(ctx, shortURLsMap)
	if err != nil {
		return response, status.Errorf(codes.InvalidArgument, err.Error())
	}

	for correlationID, shortenedURL := range shortURLsMap {
		response.Short = append(response.Short, &pb.CorrelatedShortURL{
			CorrelationId: correlationID,
			ShortUrl:      fmt.Sprintf("%s/%s", s.options.BaseURL, shortenedURL.UUID),
		})
	}

	return response, nil
}

// Stats - handler to check internal service stats
func (s *ShortenerServer) Stats(ctx context.Context, in *emptypb.Empty) (*pb.StatsResponse, error) {
	response := &pb.StatsResponse{}
	urls, users, err := s.storage.GetStats(ctx)
	if err != nil {
		return response, status.Errorf(codes.InvalidArgument, err.Error())
	}

	response.Urls = int64(urls)
	response.Users = int64(users)

	return response, nil
}

// NewShortenerServer - creates new gRPC server instance
func NewShortenerServer(options *config.Options, storage *storage.Repository) *ShortenerServer {
	s := ShortenerServer{
		options: options,
		storage: *storage,
	}
	return &s
}
