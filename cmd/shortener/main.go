package main

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
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
		http.Redirect(res, req, responseMessage, http.StatusTemporaryRedirect)
	}

	_, err := res.Write([]byte(responseMessage))
	if err != nil {
		panic("Can not send response!")
	}
}

func AddURL(Data string) (shortURL string) {
	if Urls == nil {
		Urls = make(map[string]string)
	}
	h := fnv.New32()
	h.Write([]byte(Data))
	shortURL = hex.EncodeToString(h.Sum(nil))
	Urls[shortURL] = Data
	fmt.Println(Urls)
	return
}

func FindURL(shorURL string) (string, error) {
	if Urls == nil {
		Urls = make(map[string]string)
	}
	return Urls[shorURL], nil
}

func main() {
	shortenerHandler := ShortenerHandler{}
	http.ListenAndServe(":8080", shortenerHandler)
}
