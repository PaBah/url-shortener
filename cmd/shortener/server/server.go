package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/PaBah/url-shortener.git/internal/async"
	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/dto"
	"github.com/PaBah/url-shortener.git/internal/logger"
	"github.com/PaBah/url-shortener.git/internal/middlewares"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

// Server - entity which presents application server
type Server struct {
	options *config.Options
	storage storage.Repository
}

// GetShortURLHandle - handler for list user's shortened URLs
func (s Server) GetShortURLHandle(res http.ResponseWriter, req *http.Request) {
	shortID := chi.URLParam(req, "id")

	shortenURL, _ := s.storage.FindByID(req.Context(), shortID)
	if shortenURL.DeletedFlag {
		res.WriteHeader(http.StatusGone)
		return
	}
	http.Redirect(res, req, shortenURL.OriginalURL, http.StatusTemporaryRedirect)
}

// PostURLHandle - handler for shortening URL
func (s Server) PostURLHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	shortURL := models.NewShortURL(string(body), req.Context().Value(auth.ContextUserKey).(string))
	err = s.storage.Store(req.Context(), shortURL)

	shortenedURL := fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL.UUID)
	res.Header().Set("Content-Type", "")
	res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))

	if errors.Is(err, storage.ErrConflict) {
		res.WriteHeader(http.StatusConflict)
	} else {
		res.WriteHeader(http.StatusCreated)
	}

	_, err = res.Write([]byte(shortenedURL))
	if err != nil {
		logger.Log().Error("Can not send response from PostURLHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// APIShortenHandle - handler for shortening URL via API
func (s Server) APIShortenHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	requestData := &dto.ShortenRequest{}
	err = json.Unmarshal(body, requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURL := models.NewShortURL(requestData.URL, req.Context().Value(auth.ContextUserKey).(string))
	err = s.storage.Store(req.Context(), shortURL)

	res.Header().Set("Content-Type", "application/json")
	if errors.Is(err, storage.ErrConflict) {
		res.WriteHeader(http.StatusConflict)
	} else {
		res.WriteHeader(http.StatusCreated)
	}

	responseData := dto.ShortenResponse{
		Result: fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL.UUID),
	}

	response, err := json.Marshal(responseData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = res.Write(response)
	if err != nil {
		logger.Log().Error("Can not send response from APIShortenHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PingHandle - handler for checking if DB is working
func (s Server) PingHandle(res http.ResponseWriter, req *http.Request) {
	dbStorage, ok := s.storage.(*storage.DBStorage)
	if !ok {
		http.Error(res, "Service working not on top DB storage", http.StatusInternalServerError)
		return
	}

	if err := dbStorage.Ping(req.Context()); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// APIShortenBatchHandle - handler for creation of list of short URLs
func (s Server) APIShortenBatchHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	requestData := make([]dto.BatchShortenRequest, 0)
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	shortURLsMap := make(map[string]models.ShortenURL, len(requestData))
	for _, batchRequest := range requestData {
		shortURL := models.NewShortURL(batchRequest.URL, req.Context().Value(auth.ContextUserKey).(string))
		shortURLsMap[batchRequest.CorrelationID] = shortURL
	}

	err = s.storage.StoreBatch(req.Context(), shortURLsMap)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	var responseData []dto.BatchShortenResponse
	for correlationID, shortenedURL := range shortURLsMap {
		responseData = append(responseData, dto.BatchShortenResponse{
			CorrelationID: correlationID,
			ShortURL:      fmt.Sprintf("%s/%s", s.options.BaseURL, shortenedURL.UUID),
		})
	}

	res.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(responseData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	_, err = res.Write(response)
	if err != nil {
		logger.Log().Error("Can not send response from APIShortenHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UserUrlsHandle - handler for list of short URLs of authorized user
func (s Server) UserUrlsHandle(res http.ResponseWriter, req *http.Request) {
	shortURLs, err := s.storage.GetAllUsers(req.Context())
	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	responseData := make([]dto.UsersURLsResponse, len(shortURLs))
	for _, shortURL := range shortURLs {
		responseData = append(responseData, dto.UsersURLsResponse{
			OriginalURL: shortURL.OriginalURL,
			ShortURL:    fmt.Sprintf("%s/%s", s.options.BaseURL, shortURL.UUID),
		})
	}

	res.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(responseData)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(response)
	if err != nil {
		logger.Log().Error("Can not send response from APIShortenHandle:", zap.Error(err))
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// APIDeleteUsersUrlsHandle - handler for delete short URLs
func (s Server) APIDeleteUsersUrlsHandle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	requestData := dto.DeleteURLsRequest{}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	inputCh := async.BulkDeletionDataGenerator(requestData)

	userID := req.Context().Value(auth.ContextUserKey).(string)

	channels := async.DeletionFanOut(userID, s.storage, inputCh)
	addResultCh := async.DeletionFanIn(channels...)
	async.Delete(s.storage, addResultCh)

	res.WriteHeader(http.StatusAccepted)
}

// NewRouter - creates instance of Server
func NewRouter(options *config.Options, storage *storage.Repository) *chi.Mux {
	r := chi.NewRouter()

	s := Server{
		options: options,
		storage: *storage,
	}
	r.Use(middlewares.GzipMiddleware)
	r.Use(logger.LoggerMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(auth.PublicAuthorizationMiddleware)
		r.Post("/", s.PostURLHandle)
		r.Get("/{id}", s.GetShortURLHandle)
		r.Get("/ping", s.PingHandle)
		r.Post("/api/shorten", s.APIShortenHandle)
		r.Post("/api/shorten/batch", s.APIShortenBatchHandle)
		r.MethodNotAllowed(
			func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusBadRequest)
			},
		)
	})
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthorizedMiddleware)
		r.Get("/api/user/urls", s.UserUrlsHandle)
		r.Delete("/api/user/urls", s.APIDeleteUsersUrlsHandle)
	})
	return r
}
