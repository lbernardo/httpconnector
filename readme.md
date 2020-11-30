# httpconnector

#http #go #httpconnector #mock

Use httpconnector for creating your dependencies injections  and mocks for tests

```
go get github.com/lbernardo/httpconnector
```

## Use
```go
package services

import "github.com/lbernardo/httpconnector"
func NewService(conn httpconnector.HttpConnector) {
    req1, err1 := conn.DoGet("/get", nil)
    req2, err2 := conn.DoPost("/post", "{\"name\" : \"body\"}", map[string]string{
        "Content-type" : "application/json"
    })
    // conn.DoPut(url string, body string, headers map[string]string) (*http.Response, error)
    // conn.DoDelete(url string, body string, headers map[string]string) (*http.Response, error)
    // conn.DoRequest(method string, url string, body string, headers map[string]string) (*http.Response, error)
}


```

Container
```go
package main

import "github.com/lbernardo/httpconnector"

func main() {
    c := httpconnector.NewConnector(httpconnector.Options{
        BaseURL: "https://postman-echo.com" // string,
        Headers: map[string]string{
            "x-api":  "123"
        }
    })
    
    services.NewService(c)
}
```

For tests
```go
package services

import "testing"
import "github.com/lbernardo/httpconnector"

func TestGetClient(t *testing.T) {
    s := NewService(&httpconnector.MockSuccessConnector{
        StatusCode: http.StatusOk,
        BodyResponse: "{\"success\": true}",
    })
    res, err := s.GetClient()
    if err != nil {
        t.Error("Was not expected error")
    }
    if res.StatusCode != http.StatusOk {
        t.Errorf("Was expected %v and return %v", http.StatusOk, res.StatusCode)
    }
}

func TestGetClientWithFail(t *testing.T) {
    s := NewService(&httpconnector.MockFailConnector{
        StatusCode: http.StatusInternalServerError,
        BodyResponse: "{\"success\": true}",
        Error: "Error"
    })
    res, err := s.GetClient()
    if err == nil {
        t.Error("Was expected error")
    }
    if res.StatusCode != http.StatusInternalServerError {
        t.Errorf("Was expected %v and return %v", http.StatusInternalServerError, res.StatusCode)
    }
}
```