package main

import (
	"encoding/hex"
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/server"
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type RepositoryMock struct {
	State map[string]string
}

func (rm *RepositoryMock) Store(Data string) (ID string) {
	h := fnv.New32()
	h.Write([]byte(Data))
	ID = hex.EncodeToString(h.Sum(nil))
	rm.State[ID] = Data
	return ID
}

func (rm *RepositoryMock) FindByID(ID string) (Data string, err error) {
	Data = rm.State[ID]
	return Data, nil
}

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
		{method: http.MethodPut, path: "/2187b119", requestBody: "https://practicum.yandex.kz/", expectedCode: http.StatusBadRequest, expectedBody: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.path, nil)
			if tc.requestBody != "" {
				r = httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.requestBody))
			}
			w := httptest.NewRecorder()
			rm := RepositoryMock{State: map[string]string{"2187b119": "https://practicum.yandex.ru/"}}
			sh := server.NewServer(&rm)
			sh.ServeHTTP(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			if tc.expectedBody != "" {
				fmt.Println(w.Body.String())
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
