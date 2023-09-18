package structs

import (
	"net/http"
)

type Blocklist struct {
	blockServices *BlockServices
	*HttpRequestProcessor
}

func NewBlocklist(blockServices *BlockServices, next HttpClient) *Blocklist {
	return &Blocklist{
		blockServices:        blockServices,
		HttpRequestProcessor: NewHttpRequestProcessor(next),
	}
}

func (b *Blocklist) SendRequest(w http.ResponseWriter, r *http.Request) {
	hostOfURL := r.URL.Host
	for _, blockService := range *(b.blockServices) {
		if blockService.DomainName() == hostOfURL {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		for _, ip := range blockService.IPs() {
			if ip == hostOfURL {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
		}
	}

	b.next.SendRequest(w, r)
}
