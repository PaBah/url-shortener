package main

import (
	"fmt"
	"github.com/PaBah/url-shortener.git/cmd/shortener/storage"
	"io"
	"net/http"
	"strconv"
)

type ShortenerHandler struct{}

var Urls map[string]string

func (sh ShortenerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var responseMessage string
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		responseMessage = "Unsupported HTTP Method"
	}

	if req.Method == http.MethodPost {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			responseMessage = "Invalid body"
		}

		shortURL := AddURL(string(body))
		shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortURL)
		res.Header().Set("Content-Type", "")
		res.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
		res.WriteHeader(http.StatusCreated)
		responseMessage = shortenedURL
	}

	if req.Method == http.MethodGet {
		responseMessage, _ = FindURL(req.URL.EscapedPath()[1:])
		fmt.Println(req.URL)
		fmt.Println(responseMessage)
		http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
	}

	_, err := res.Write([]byte(responseMessage))
	if err != nil {
		panic("Can not send response!")
	}
}

func AddURL(Url string) (shortURL string) {
	cs := storage.CringeStorage{}
	shortURL = cs.Store(Url)
	return
}

func FindURL(shorURL string) (string, error) {
	cs := storage.CringeStorage{}
	result, _ := cs.FindByID(shorURL)
	return result, nil
}

func main() {
	shortenerHandler := ShortenerHandler{}
	http.ListenAndServe(":8080", shortenerHandler)
}
