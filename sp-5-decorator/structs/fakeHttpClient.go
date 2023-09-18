package structs

import (
	"fmt"
	"net/http"
)

type FakeHttpClient struct{}

func NewFakeHttpClient() *FakeHttpClient {
	return &FakeHttpClient{}
}

func (f *FakeHttpClient) SendRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n【模擬發送請求】\n")
	fmt.Printf("FakeHttpClient request: %v\n", r)

	fmt.Println("【模擬接收回應】")
	w.WriteHeader(http.StatusOK)
}
