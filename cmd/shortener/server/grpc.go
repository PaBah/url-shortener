package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/PaBah/url-shortener.git/internal/async"
	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	pb "github.com/PaBah/url-shortener.git/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// ShortenerServer shortener gRPC server
type ShortenerServer struct {
	pb.UnimplementedShortenerServer

	options *config.Options
	storage storage.Repository
}

func (s *ShortenerServer) Short(ctx context.Context, in *pb.ShortRequest) (*pb.ShortResponse, error) {
	response := &pb.ShortResponse{}
	shortURL := models.NewShortURL(in.Url, in.UserId)
	err := s.storage.Store(ctx, shortURL)

	response.Result = shortURL.UUID
	if errors.Is(err, storage.ErrConflict) {
		return response, err
	}
	return response, err
}

func (s *ShortenerServer) Expand(ctx context.Context, in *pb.ExpandRequest) (*pb.ExpandResponse, error) {
	response := &pb.ExpandResponse{}

	shortenURL, _ := s.storage.FindByID(ctx, in.ShortId)
	if shortenURL.DeletedFlag {
		return response, errors.New("shorten URL already expanded")
	}
	response.Url = shortenURL.OriginalURL
	return response, nil
}

func (s *ShortenerServer) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	response := new(emptypb.Empty)

	inputCh := async.BulkDeletionDataGenerator(in.Id)

	channels := async.DeletionFanOut(in.UserId, s.storage, inputCh)
	addResultCh := async.DeletionFanIn(channels...)
	async.Delete(s.storage, addResultCh)

	return response, nil
}

func (s *ShortenerServer) GetUserBucket(ctx context.Context, in *pb.GetUserBucketRequest) (*pb.GetUserBucketResponse, error) {
	response := &pb.GetUserBucketResponse{}
	shortURLs, err := s.storage.GetAllUsers(context.WithValue(ctx, auth.ContextUserKey, in.UserId))
	if err != nil {
		return response, err
	}

	for _, shortURL := range shortURLs {
		response.Pair = append(response.Pair, &pb.Pair{
			OriginalUrl: shortURL.OriginalURL,
			ShortUrl:    fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL.UUID),
		})
	}

	return response, nil
}

func (s *ShortenerServer) ShortBatch(ctx context.Context, in *pb.ShortBatchRequest) (*pb.ShortBatchResponse, error) {
	response := &pb.ShortBatchResponse{}

	shortURLsMap := make(map[string]models.ShortenURL, len(in.Original))
	for _, batchRequest := range in.Original {
		shortURL := models.NewShortURL(batchRequest.OriginalUrl, in.UserId)
		shortURLsMap[batchRequest.CorrelationId] = shortURL
	}

	err := s.storage.StoreBatch(ctx, shortURLsMap)
	if err != nil {
		return response, err
	}

	for correlationID, shortenedURL := range shortURLsMap {
		response.Short = append(response.Short, &pb.CorrelatedShortURL{
			CorrelationId: correlationID,
			ShortUrl:      fmt.Sprintf("%s/%s", s.options.BaseURL, shortenedURL.UUID),
		})
	}

	return response, nil
}

func (s *ShortenerServer) Stats(ctx context.Context, in *emptypb.Empty) (*pb.StatsResponse, error) {
	response := &pb.StatsResponse{}
	urls, users, err := s.storage.GetStats(ctx)
	if err != nil {
		return response, err
	}

	response.Urls = int64(urls)
	response.Users = int64(users)

	return response, nil
}

func NewShortenerServer(options *config.Options, storage *storage.Repository) *ShortenerServer {
	s := ShortenerServer{
		options: options,
		storage: *storage,
	}
	return &s
}
