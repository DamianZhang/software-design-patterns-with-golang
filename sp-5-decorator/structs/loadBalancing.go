package structs

import (
	"net/http"
	"strings"
)

type LoadBalancing struct {
	indexOfRoundRobin int
	availableServices *AvailableServices
	*HttpRequestProcessor
}

func NewLoadBalancing(availableServices *AvailableServices, next HttpClient) *LoadBalancing {
	return &LoadBalancing{
		indexOfRoundRobin:    0,
		availableServices:    availableServices,
		HttpRequestProcessor: NewHttpRequestProcessor(next),
	}
}

func (l *LoadBalancing) SendRequest(w http.ResponseWriter, r *http.Request) {
	var (
		hostOfURL        = r.URL.Host
		availableService = l.availableServices.GetAvailableServiceByHostOfURL(hostOfURL)
	)
	if availableService == nil {
		l.next.SendRequest(w, r)
		return
	}

	availableIPs := availableService.GetAvailableIPs()
	if availableIPs == "" {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var (
		hosts      = strings.Split(availableIPs, ",")
		lenOfHosts = len(hosts)
	)
	l.indexOfRoundRobin %= lenOfHosts
	chosenHost := hosts[l.indexOfRoundRobin]
	l.indexOfRoundRobin++

	r.URL.Host = chosenHost
	r.Host = chosenHost
	l.next.SendRequest(w, r)
}
