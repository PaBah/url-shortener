package main

import (
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/PaBah/url-shortener.git/internal/config"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddURL(t *testing.T) {
	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
	testCases := []struct {
		method       string
		requestBody  string
		path         string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodPost, path: "/", requestBody: "https://practicum.yandex.ru/", expectedCode: http.StatusCreated, expectedBody: "http://localhost:8080/2187b119"},
		{method: http.MethodGet, path: "/2187b119", requestBody: "", expectedCode: http.StatusTemporaryRedirect, expectedBody: ""},
		{method: http.MethodPut, path: "/2187b119", requestBody: "https://practicum.yandex.ru/", expectedCode: http.StatusBadRequest, expectedBody: ""},
	}

	options := &config.Options{}
	ParseFlags(options)

	ctrl := gomock.NewController(t)
	var store storage.Repository
	rm := storage.NewMockRepository(ctrl)
	store = rm

	rm.
		EXPECT().
		Store(gomock.Eq("https://practicum.yandex.ru/")).
		Return("2187b119").
		AnyTimes()
	rm.
		EXPECT().
		FindByID("2187b119").
		Return("https://practicum.yandex.ru/", nil).
		AnyTimes()

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.requestBody != "" {
				r = httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.requestBody))
			}
			w := httptest.NewRecorder()

			sh := server.NewRouter(options, &store)
			sh.ServeHTTP(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
