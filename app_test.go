package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	client := http.Client{/*모든 옵션값을 기본값으로 설정했습니다*/}
	req,errRequset := http.NewRequest(http.MethodGet,
		"http://localhost:5000/hello",nil)
	if errRequset != nil {
		return
	}
	res,err := client.Do(req)
	if err != nil {
		return
	}
	fmt.Printf("http response : %v",res)
	return
}