package httpconnector

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSuccess_DoGet(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockSuccessConnector{}
		res, err := c.DoGet("https://google.com.br", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestSuccess_DoPost(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockSuccessConnector{}
		res, err := c.DoPost("https://google.com.br", "", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestSuccess_DoDelete(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockSuccessConnector{}
		res, err := c.DoDelete("https://google.com.br", "", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestSuccess_DoPut(t *testing.T) {
	t.Run("Test request", func(t *testing.T) {
		c := &MockSuccessConnector{}
		res, err := c.DoPut("https://google.com.br", "", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Was expected %v and return %v\n", http.StatusOK, res.StatusCode)
		}
	})
}

func TestSuccess_DoRequest(t *testing.T) {
	t.Run("Test request with Response", func(t *testing.T) {
		c := &MockSuccessConnector{
			BodyResponse: "{\"name\": \"test\"}",
			StatusCode:   http.StatusAccepted,
		}
		res, err := c.DoPost("https://google.com.br", "", nil)
		if err != nil {
			t.Error("Was not expected error")
		}
		if res.StatusCode != http.StatusAccepted {
			t.Errorf("Was expected %v and return %v\n", http.StatusAccepted, res.StatusCode)
		}
		var b map[string]string
		json.NewDecoder(res.Body).Decode(&b)
		if b["name"] != "test" {
			t.Errorf("Was expected %v and return %v\n", "test", b["name"])
		}
	})
}
