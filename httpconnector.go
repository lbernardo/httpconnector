package httpconnector

import "net/http"

type HttpConnector interface {
	DoGet(url string, headers map[string]string) (*http.Response, error)
	DoPost(url string, body string, headers map[string]string) (*http.Response, error)
	DoPut(url string, body string, headers map[string]string) (*http.Response, error)
	DoDelete(url string, body string, headers map[string]string) (*http.Response, error)
	DoRequest(method string, url string, body string, headers map[string]string) (*http.Response, error)
}

func AddHeaders(req *http.Request, headers map[string]string) {
	if headers != nil {
		for n, v := range headers {
			req.Header.Add(n, v)
		}
	}
}
