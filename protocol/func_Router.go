package protocol

import (
	"fmt"
	"net/http"
)

//컨트롤러 매핑 함수 정의

func Logging(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter , request *http.Request ) {
		fmt.Printf("request Query : %v",request.URL.Query())
		h(writer,request)
	}
}
func CheckAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter , request *http.Request ) {
		if request.Header.Get("apiKey") == "" {
			ErrorResponse(writer, "apikey is empty")
		}
		h(writer,request)
	}
}
func GET(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter , request *http.Request ) {
		if request.Method != "GET" {
			ErrorResponse(writer, "unsupported http method")
		}
		h(writer,request)
	}
}

func ErrorResponse(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

//example function
func Hello(writer http.ResponseWriter , request *http.Request ) {
	writer.Write([]byte("Hello World"))
}