package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/auth"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServer(t *testing.T) {
	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
	testCases := []struct {
		method       string
		requestBody  string
		path         string
		expectedCode int
		expectedBody string
		storage      storage.Repository
	}{
		{method: http.MethodPost, path: "/", requestBody: "https://practicum.yandex.ru/", expectedCode: http.StatusCreated, expectedBody: "http://localhost:8080/2187b119"},
		{method: http.MethodPost, path: "/", requestBody: "http://prjdzevto8.yandex", expectedCode: http.StatusConflict, expectedBody: "http://localhost:8080/a033a480"},
		{method: http.MethodGet, path: "/2187b119", requestBody: "", expectedCode: http.StatusTemporaryRedirect, expectedBody: ""},
		{method: http.MethodGet, path: "/2a49568d", requestBody: "", expectedCode: http.StatusTemporaryRedirect, expectedBody: ""},
		{method: http.MethodPut, path: "/2187b119", requestBody: "https://practicum.yandex.ru/", expectedCode: http.StatusBadRequest, expectedBody: ""},
		{method: http.MethodPost, path: "/api/shorten", requestBody: `{"url": "https://practicum.yandex.kz/"}`, expectedCode: http.StatusCreated, expectedBody: `{"result":"http://localhost:8080/2a49568d"}`},
		{method: http.MethodPost, path: "/api/shorten", requestBody: `{"url": "https://practicum.yande`, expectedCode: http.StatusInternalServerError, expectedBody: ""},
		{method: http.MethodPost, path: "/api/shorten", requestBody: `{"url": "http://prjdzevto8.yandex"}`, expectedCode: http.StatusConflict, expectedBody: `{"result":"http://localhost:8080/a033a480"}`},
		{method: http.MethodGet, path: "/ping", requestBody: "", expectedCode: http.StatusInternalServerError, expectedBody: ""},
		{method: http.MethodPost, path: "/api/shorten/batch", requestBody: `[x.kz/"}]`, expectedCode: http.StatusInternalServerError, expectedBody: ""},
		{
			method:       http.MethodPost,
			path:         "/api/shorten/batch",
			requestBody:  `[{"correlation_id": "1","original_url": "https://practicum.yandex.kz/"}]`,
			expectedCode: http.StatusCreated,
			expectedBody: `[{"correlation_id":"1","short_url":"http://localhost:8080/2a49568d"}]`,
		},
		{
			method:       http.MethodPost,
			path:         "/api/shorten/batch",
			requestBody:  `[{"correlation_id": "1","original_url": "https://practicum.kz/"}]`,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
		{
			method:       http.MethodGet,
			path:         "/api/user/urls",
			requestBody:  `[{"short_url": "http://localhost:8080/2a49568d","original_url": "https://practicum.yandex.kz/"}]`,
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			method:       http.MethodGet,
			path:         "/api/user/urls",
			requestBody:  ``,
			expectedCode: http.StatusNoContent,
			expectedBody: "",
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
		Store(gomock.Any(), gomock.Eq(models.NewShortURL("https://practicum.yandex.ru/", "1"))).
		Return(nil).
		AnyTimes()
	rm.
		EXPECT().
		Store(gomock.Any(), gomock.Eq(models.NewShortURL("http://prjdzevto8.yandex", "1"))).
		Return(storage.ErrConflict).
		AnyTimes()
	rm.
		EXPECT().
		FindByID(gomock.Any(), "2187b119").
		Return(models.NewShortURL("https://practicum.yandex.ru/", "1"), nil).
		AnyTimes()
	rm.
		EXPECT().
		Store(gomock.Any(), gomock.Eq(models.NewShortURL("https://practicum.yandex.kz/", "1"))).
		Return(nil).
		AnyTimes()
	rm.
		EXPECT().
		FindByID(gomock.Any(), "2a49568d").
		Return(models.NewShortURL("https://practicum.yandex.kz/", "1"), nil).
		AnyTimes()
	rm.
		EXPECT().
		StoreBatch(gomock.Any(), gomock.Eq(map[string]models.ShortenURL{"1": models.NewShortURL("https://practicum.yandex.kz/", "1")})).
		Return(nil).
		AnyTimes()
	err := errors.New("Error")
	rm.
		EXPECT().
		StoreBatch(gomock.Any(), gomock.Eq(map[string]models.ShortenURL{"1": models.NewShortURL("https://practicum.kz/", "1")})).
		Return(err).
		AnyTimes()
	rm.
		EXPECT().
		GetAllUsers(gomock.Any()).
		Return([]models.ShortenURL{models.NewShortURL("https://practicum.kz/", "1")}, nil).
		Times(1)
	rm.
		EXPECT().
		GetAllUsers(gomock.Any()).
		Return([]models.ShortenURL{}, err).
		Times(1)
	sh := NewRouter(options, &store)

	//for i, tc := range testCases {
	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {

			r := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.requestBody != "" {
				r = httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.requestBody))
			}
			w := httptest.NewRecorder()
			JWTToken, _ := auth.BuildJWTString("1")
			r.Header.Set("Cookie", "Authorization="+JWTToken)

			sh.ServeHTTP(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
