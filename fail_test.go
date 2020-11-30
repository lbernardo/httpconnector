package httpconnector

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestFail_DoGet(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockFailConnector{}
		res, err := c.DoGet("https://google.com.br", nil)
		if err == nil {
			t.Error("Was expected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Was expected %v and return %v\n", http.StatusBadRequest, res.StatusCode)
		}
	})
}

func TestFail_DoPost(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockFailConnector{}
		res, err := c.DoPost("https://google.com.br", "", nil)
		if err == nil {
			t.Error("Was  expected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Was expected %v and return %v\n", http.StatusBadRequest, res.StatusCode)
		}
	})
}

func TestFail_DoDelete(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockFailConnector{}
		res, err := c.DoDelete("https://google.com.br", "", nil)
		if err == nil {
			t.Error("Was  expected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Was expected %v and return %v\n", http.StatusBadRequest, res.StatusCode)
		}
	})
}

func TestFail_DoPut(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockFailConnector{}
		res, err := c.DoPut("https://google.com.br", "", nil)
		if err == nil {
			t.Error("Was  expected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Was expected %v and return %v\n", http.StatusBadRequest, res.StatusCode)
		}
	})
}

func TestFail_DoRequest(t *testing.T) {
	t.Run("Test request with Response", func(t *testing.T) {
		c := &MockFailConnector{
			BodyResponse: "{\"name\": \"test\"}",
			StatusCode:   http.StatusInternalServerError,
			Error:        "Error",
		}
		res, err := c.DoPost("https://google.com.br", "", nil)
		if err == nil {
			t.Error("Was expected error")
		}
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("Was expected %v and return %v\n", http.StatusInternalServerError, res.StatusCode)
		}
		var b map[string]string
		json.NewDecoder(res.Body).Decode(&b)
		if b["name"] != "test" {
			t.Errorf("Was expected %v and return %v\n", "test", b["name"])
		}
	})
}
