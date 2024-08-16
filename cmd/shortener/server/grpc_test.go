package server

import (
	"context"
	"errors"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/config"
	pb "github.com/PaBah/url-shortener.git/internal/gen/proto/shortener/v1"
	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Short(t *testing.T) {
	testCases := []struct {
		storage        storage.Repository
		requestURL     string
		expectedResult string
		expectedError  bool
		errorCode      codes.Code
	}{
		{requestURL: "https://practicum.yandex.ru/", expectedError: false, expectedResult: "2187b119"},
		{requestURL: "https://bad.url.ru/", expectedError: true, errorCode: codes.InvalidArgument, expectedResult: "4e76a198"},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		Store(gomock.Any(), gomock.Eq(models.NewShortURL("https://practicum.yandex.ru/", "1"))).
		Return(nil).
		AnyTimes()
	rm.
		EXPECT().
		Store(gomock.Any(), gomock.Eq(models.NewShortURL("https://bad.url.ru/", "1"))).
		Return(storage.ErrConflict).
		AnyTimes()

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run("Store", func(t *testing.T) {
			result, err := sh.Short(context.Background(), &pb.ShortRequest{Url: tc.requestURL, UserId: "1"})

			assert.Equal(t, tc.expectedResult, result.Result, "Expected result get")
			if tc.expectedError {
				e, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			}
		})
	}
}

func Test_Expand(t *testing.T) {
	testCases := []struct {
		storage        storage.Repository
		shortID        string
		expectedResult string
		expectedError  bool
		errorCode      codes.Code
	}{
		{shortID: "2187b119", expectedError: false, expectedResult: "https://practicum.yandex.ru/"},
		{shortID: "4e76a198", expectedError: true, errorCode: codes.InvalidArgument, expectedResult: ""},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		FindByID(gomock.Any(), "2187b119").
		Return(models.NewShortURL("https://practicum.yandex.ru/", "1"), nil).
		AnyTimes()
	rm.
		EXPECT().
		FindByID(gomock.Any(), "4e76a198").
		Return(models.ShortenURL{OriginalURL: "https://practicum.yandex.ru/", UserID: "1", DeletedFlag: true}, nil).
		AnyTimes()

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run(tc.shortID, func(t *testing.T) {
			result, err := sh.Expand(context.Background(), &pb.ExpandRequest{ShortId: tc.shortID})

			assert.Equal(t, tc.expectedResult, result.Url, "Expected result get")
			if tc.expectedError {
				e, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	testCases := []struct {
		storage       storage.Repository
		shortID       string
		expectedError bool
	}{
		{shortID: "2187b119", expectedError: false},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		AsyncCheckURLsUserID(gomock.Eq("1"), gomock.Any()).
		Return(make(chan string)).AnyTimes()

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run(tc.shortID, func(t *testing.T) {
			_, err := sh.Delete(context.Background(), &pb.DeleteRequest{Id: []string{tc.shortID}, UserId: "1"})
			assert.NoError(t, err, "Expected success")
			//if tc.expectedError {
			//	e, ok := status.FromError(err)
			//	require.True(t, ok)
			//	assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			//}
		})
	}
}

func Test_GetUserBucket(t *testing.T) {
	testCases := []struct {
		storage        storage.Repository
		userID         string
		expectedError  bool
		expectedResult string
		errorCode      codes.Code
	}{
		{userID: "1", expectedError: false, expectedResult: "https://practicum.yandex.kz/"},
		{userID: "1", expectedError: true, errorCode: codes.InvalidArgument, expectedResult: ""},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		GetAllUsers(gomock.Any()).
		Return([]models.ShortenURL{models.NewShortURL("https://practicum.yandex.kz/", "1")}, nil).
		Times(1)
	rm.
		EXPECT().
		GetAllUsers(gomock.Any()).
		Return([]models.ShortenURL{}, errors.New("Error")).
		Times(1)

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run(tc.userID, func(t *testing.T) {
			result, err := sh.GetUserBucket(context.Background(), &pb.GetUserBucketRequest{UserId: tc.userID})

			if tc.expectedError {
				e, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			} else {
				assert.Equal(t, tc.expectedResult, result.Data[0].OriginalUrl, "Expected result get")
			}
		})
	}
}

func Test_ShortBatch(t *testing.T) {
	testCases := []struct {
		storage        storage.Repository
		request        string
		expectedError  bool
		expectedResult string
		errorCode      codes.Code
	}{
		{
			request:        "https://practicum.yandex.kz/",
			expectedError:  false,
			expectedResult: "http://localhost:8080/2a49568d",
		},
		{
			request:        "https://practicum.yandex.kz/",
			expectedError:  true,
			errorCode:      codes.InvalidArgument,
			expectedResult: "",
		},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		StoreBatch(gomock.Any(), gomock.Eq(map[string]models.ShortenURL{"1": models.NewShortURL("https://practicum.yandex.kz/", "1")})).
		Return(nil).
		Times(1)
	rm.
		EXPECT().
		StoreBatch(gomock.Any(), gomock.Eq(map[string]models.ShortenURL{"1": models.NewShortURL("https://practicum.yandex.kz/", "1")})).
		Return(errors.New("Error")).
		Times(1)

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run(tc.request, func(t *testing.T) {
			result, err := sh.ShortBatch(context.Background(), &pb.ShortBatchRequest{
				UserId: "1",
				Original: []*pb.CorrelatedOriginalURL{
					{
						CorrelationId: "1",
						OriginalUrl:   tc.request,
					},
				}})

			if tc.expectedError {
				e, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			} else {
				assert.Equal(t, tc.expectedResult, result.Short[0].ShortUrl, "Expected result get")
			}
		})
	}
}

func Test_Stats(t *testing.T) {
	testCases := []struct {
		storage        storage.Repository
		expectedError  bool
		expectedResult []int64
		errorCode      codes.Code
	}{
		{
			expectedError:  false,
			expectedResult: []int64{1, 2},
		},
		{
			expectedError:  true,
			errorCode:      codes.InvalidArgument,
			expectedResult: []int64{},
		},
	}
	options := &config.Options{
		ServerAddress: ":8080",
		BaseURL:       "http://localhost:8080",
		DatabaseDSN:   "wrong DSN",
	}

	var store storage.Repository
	ctrl := gomock.NewController(t)
	rm := mock.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		GetStats(gomock.Any()).
		Return(2, 1, nil).
		Times(1)
	rm.
		EXPECT().
		GetStats(gomock.Any()).
		Return(0, 0, errors.New("Error")).
		Times(1)

	sh := NewShortenerServer(options, &store)

	for _, tc := range testCases {
		t.Run("test", func(t *testing.T) {
			result, err := sh.Stats(context.Background(), &pb.StatsRequest{})

			if tc.expectedError {
				e, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tc.errorCode, e.Code(), "Expected error code get")
			} else {
				assert.Equal(t, tc.expectedResult[0], result.Users, "Expected result users get")
				assert.Equal(t, tc.expectedResult[1], result.Urls, "Expected result urls get")
			}
		})
	}
}
