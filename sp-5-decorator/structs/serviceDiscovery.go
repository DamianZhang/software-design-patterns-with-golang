package structs

import (
	"math/rand"
	"net/http"
	"time"
)

var (
	Rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type ServiceDiscovery struct {
	availableServices *AvailableServices
	*HttpRequestProcessor
}

func NewServiceDiscovery(availableServices *AvailableServices, next HttpClient) *ServiceDiscovery {
	return &ServiceDiscovery{
		availableServices:    availableServices,
		HttpRequestProcessor: NewHttpRequestProcessor(next),
	}
}

func (s *ServiceDiscovery) SendRequest(w http.ResponseWriter, r *http.Request) {
	var (
		hostOfURL        = r.URL.Host
		availableService = s.availableServices.GetAvailableServiceByHostOfURL(hostOfURL)
	)
	if availableService == nil {
		s.next.SendRequest(w, r)
		return
	}

	// 服務探索
	for _, ip := range availableService.IPs() {
		if ip == "" {
			continue
		}

		if !s.isAvailableIP(ip) {
			// 必須以 ip 的值產生新的 deletedIP
			// 否則最後 callback 裡的 ip 都會指向同一個記憶體位置
			// 也就是 for 迴圈最後那一個 ip 的記憶體位置
			// 導致最後只會刪除跟新增 "最後那一個 ip"
			deletedIP := ip
			availableService.DeleteIP(deletedIP)
			time.AfterFunc(5*time.Second, func() {
				availableService.AddIP(deletedIP)
			})
		}
	}

	availableIP := availableService.GetAvailableIP()
	if availableIP == "" {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	r.URL.Host = availableIP
	r.Host = availableIP
	s.next.SendRequest(w, r)
}

// 模擬 Service 有 3/100 的機率失效
func (s *ServiceDiscovery) isAvailableIP(ip string) bool {
	return Rander.Intn(100) >= 3
}
