package httpconnector

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type MockSuccessConnector struct {
	StatusCode   int
	BodyResponse string
}

func (c *MockSuccessConnector) DoGet(url string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, url, "", headers)
}

func (c *MockSuccessConnector) DoPost(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPost, url, body, headers)
}

func (c *MockSuccessConnector) DoPut(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPut, url, body, headers)
}

func (c *MockSuccessConnector) DoDelete(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, url, body, headers)
}

func (c *MockSuccessConnector) DoRequest(method string, url string, body string, headers map[string]string) (*http.Response, error) {
	response := &http.Response{}
	response.StatusCode = http.StatusOK
	if c.BodyResponse != "" {
		response.Body = ioutil.NopCloser(strings.NewReader(c.BodyResponse))
	}
	if c.StatusCode != 0 {
		response.StatusCode = c.StatusCode
	}
	return response, nil
}
