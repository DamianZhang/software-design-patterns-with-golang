package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sp-5-decorator/structs"
)

func main() {
	fmt.Println("Http Client is starting...")
	// testURLs := []string{
	// 	"https://957.com",
	// 	"https://blocklist.tw",
	// 	"https://50.0.0.1",
	// 	"https://957tw.com",
	// 	"https://35.0.0.1",
	// }

	// Code Review 後補充 Trade Off:
	// 方法 1. ServiceDiscovery, LoadBalancing 管理各自的 AvailableServices
	// 	解釋: 因為各自管理(兩者毫無關聯)，所以具備高內聚力。
	// 	Resulting Context: 解決 F-SRP，但有 F-CD。
	//
	// 方法 2. ServiceDiscovery, LoadBalancing 共同管理 AvailableServices
	// 	解釋: 因為共同管理，所以不具備高內聚力。
	// 	當有新需求時，此處的程式會有改不動的問題。
	// 	Resulting Context: 解決 F-CD，但有 F-SRP。
	//
	// F-SRP = Force - Single Responsibility Principle = 單一職責原則
	// F-CD = Force - Code Duplication = 重複程式碼
	//
	// 此處情境要求不能有 F-CD，所以採用方法 2
	availableServices, err := structs.NewAvailableServices("./AvailableServices.txt")
	if err != nil {
		fmt.Println("new available services failed:", err)
		return
	}

	blockServices, err := structs.NewBlockServices("./BlockServices.txt")
	if err != nil {
		fmt.Println("new block services failed:", err)
		return
	}

	httpClient := structs.NewServiceDiscovery(
		availableServices,
		structs.NewLoadBalancing(
			availableServices,
			structs.NewBlocklist(
				blockServices,
				structs.NewFakeHttpClient())))

	scanURLToSendRequest(httpClient)
}

func scanURLToSendRequest(httpClient structs.HttpClient) {
	var url string
	for {
		fmt.Println("please input URL:")
		fmt.Scanf("%s", &url)

		var (
			request, _ = http.NewRequest(http.MethodGet, url, nil)
			response   = httptest.NewRecorder()
		)
		httpClient.SendRequest(response, request)

		fmt.Printf("\n＊套件使用者接收回應＊\n")
		fmt.Printf("response.Code: %v\n\n", response.Code)
	}
}
