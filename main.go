package main

import (
	"SeaMaple/protocol"
	"SeaMaple/static"
	"net/http"
)

func init() {

}
//main 함수, 서버 시작, GraceFull Reset 등 정의

func main() {
	serverHandler := http.NewServeMux()

	server := &http.Server{
		Addr:              "localhost:5000",
		Handler:           protocol.InitRouter(serverHandler),
	}
	//TODO add Error Handling
	static.InitLogger()
	l := static.NewLogger()
	r := protocol.NewRouter(server,l)
	r.Run()
}