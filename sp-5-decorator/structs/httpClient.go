package structs

import "net/http"

type HttpClient interface {
	SendRequest(w http.ResponseWriter, r *http.Request)
}
