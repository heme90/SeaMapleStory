package protocol

import (
	"net/http"
)

func InitRouter(server *http.ServeMux) *http.ServeMux{
	server.HandleFunc("/hello", GET(Logging(Hello)))
	//add routing function
	//ex) server.HandleFunc("/path/of/api", POST(Logging(func(){})))
	return server
}