package httpconnector

import (
	"bytes"
	"fmt"
	"net/http"
)

type Connector struct {
	Options
	client *http.Client
}

func NewConnector(opt ...Options) *Connector {
	client := &http.Client{}
	if len(opt) > 0 {
		return &Connector{
			Options: opt[0],
			client:  client,
		}
	}
	return &Connector{
		client: client,
	}
}

func (c *Connector) DoGet(url string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, url, "", headers)
}

func (c *Connector) DoPost(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPost, url, body, headers)
}

func (c *Connector) DoPut(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodPut, url, body, headers)
}

func (c *Connector) DoDelete(url string, body string, headers map[string]string) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, url, body, headers)
}

func (c *Connector) DoRequest(method string, url string, body string, headers map[string]string) (*http.Response, error) {
	url = fmt.Sprintf("%v%v", c.BaseURL, url)
	bodyRequest := bytes.NewBufferString(body)
	req, _ := http.NewRequest(method, url, bodyRequest)
	AddHeaders(req, c.Headers)
	AddHeaders(req, headers)
	return c.client.Do(req)
}
