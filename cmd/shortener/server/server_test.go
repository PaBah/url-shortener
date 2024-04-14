package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/mock"
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
		{method: http.MethodGet, path: "/2187b119", requestBody: "", expectedCode: http.StatusTemporaryRedirect, expectedBody: ""},
		{method: http.MethodGet, path: "/2a49568d", requestBody: "", expectedCode: http.StatusTemporaryRedirect, expectedBody: ""},
		{method: http.MethodPut, path: "/2187b119", requestBody: "https://practicum.yandex.ru/", expectedCode: http.StatusBadRequest, expectedBody: ""},
		{method: http.MethodPost, path: "/api/shorten", requestBody: `{"url": "https://practicum.yandex.kz/"}`, expectedCode: http.StatusCreated, expectedBody: `{"result":"http://localhost:8080/2a49568d"}`},
		{method: http.MethodGet, path: "/ping", requestBody: "", expectedCode: http.StatusInternalServerError, expectedBody: ""},
		{
			method:       http.MethodPost,
			path:         "/api/shorten/batch",
			requestBody:  `[{"correlation_id": "1","original_url": "https://practicum.yandex.kz/"}]`,
			expectedCode: http.StatusCreated,
			expectedBody: `[{"correlation_id":"1","short_url":"http://localhost:8080/2a49568d"}]`,
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
		Store(gomock.Any(), gomock.Eq("https://practicum.yandex.ru/")).
		Return("2187b119").
		AnyTimes()
	rm.
		EXPECT().
		FindByID(gomock.Any(), "2187b119").
		Return("https://practicum.yandex.ru/", nil).
		AnyTimes()
	rm.
		EXPECT().
		Store(gomock.Any(), gomock.Eq("https://practicum.yandex.kz/")).
		Return("2a49568d").
		AnyTimes()
	rm.
		EXPECT().
		FindByID(gomock.Any(), "2a49568d").
		Return("https://practicum.yandex.kz/", nil).
		AnyTimes()
	rm.
		EXPECT().
		StoreBatch(gomock.Any(), gomock.Eq(map[string]string{"1": "https://practicum.yandex.kz/"})).
		Return(map[string]string{"1": "2a49568d"}, nil).
		AnyTimes()
	sh := NewRouter(options, &store)

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.requestBody != "" {
				r = httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.requestBody))
			}
			w := httptest.NewRecorder()

			sh.ServeHTTP(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
