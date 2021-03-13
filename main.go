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
		//protocol.InitRouter 함수에서 라우팅 테이블 초기화
		Handler:           protocol.InitRouter(serverHandler),
	}
	//TODO add Error Handling
	static.InitLogger()
	l := static.NewLogger()
	r := protocol.NewRouter(server,l)
	r.Run()
}