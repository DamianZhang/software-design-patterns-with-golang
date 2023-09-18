package structs

type HttpRequestProcessor struct {
	next HttpClient
}

func NewHttpRequestProcessor(next HttpClient) *HttpRequestProcessor {
	return &HttpRequestProcessor{next: next}
}
