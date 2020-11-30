package httpconnector

import (
	"net/http"
	"testing"
)

func TestNewConnector(t *testing.T) {
	t.Run("Test without Options", func(t *testing.T) {
		c := NewConnector()
		if c.BaseURL != "" {
			t.Error("Was exepected BaseUrl empty")
		}

	})
	t.Run("Test with Options", func(t *testing.T) {
		c := NewConnector(Options{
			BaseURL: "https://example.com",
		})
		if c.BaseURL == "" {
			t.Error("Was not exepected BaseUrl empty")
		}
	})
	t.Run("Test return HttpConnector Interface", func(t *testing.T) {
		f := func(h HttpConnector) {}
		c := NewConnector()
		f(c)
	})
}

func TestNewConnector_DoGet(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := NewConnector(Options{BaseURL: "https://postman-echo.com"})
		res, err := c.DoGet("/get?foo=bar", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestNewConnector_DoPost(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := NewConnector(Options{BaseURL: "https://postman-echo.com"})
		res, err := c.DoPost("/post", "{\"name\": \"test\"}", map[string]string{
			"Content-type": "application/json",
		})
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestNewConnector_DoDelete(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := NewConnector(Options{BaseURL: "https://postman-echo.com"})
		res, err := c.DoDelete("/delete", "{\"name\": \"test\"}", map[string]string{
			"Content-type": "application/json",
		})
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestNewConnector_DoPut(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := NewConnector(Options{BaseURL: "https://postman-echo.com"})
		res, err := c.DoPut("/put", "{\"name\": \"test\"}", map[string]string{
			"Content-type": "application/json",
		})
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestNewConnector_DoRequest(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := NewConnector(Options{BaseURL: "https://postman-echo.com"})
		res, err := c.DoRequest(http.MethodPost, "/post", "{\"name\": \"test\"}", map[string]string{
			"Content-type": "application/json",
		})
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}
