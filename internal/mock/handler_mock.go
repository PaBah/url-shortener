package mock

import (
	"fmt"
	"net/http"
)

func NewHandlerMock(ResponseData string, StatusCode int) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(StatusCode)
		amount, _ := res.Write([]byte(ResponseData))
		fmt.Println(amount)
	}
}
