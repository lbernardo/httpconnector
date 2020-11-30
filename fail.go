package httpconnector

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type MockFailConnector struct {
	StatusCode   int
	BodyResponse string
	Error        string
}

func (c *MockFailConnector) DoGet(url string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, url, "", headers)
}

func (c *MockFailConnector) DoPost(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPost, url, body, headers)
}

func (c *MockFailConnector) DoPut(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPut, url, body, headers)
}

func (c *MockFailConnector) DoDelete(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, url, body, headers)
}

func (c *MockFailConnector) DoRequest(method string, url string, body string, headers map[string]string) (*http.Response, error) {
	response := &http.Response{}
	response.StatusCode = http.StatusBadRequest
	if c.BodyResponse != "" {
		response.Body = ioutil.NopCloser(strings.NewReader(c.BodyResponse))
	}
	if c.StatusCode != 0 {
		response.StatusCode = c.StatusCode
	}
	return response, errors.New(c.Error)
}
