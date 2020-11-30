package httpconnector

type Options struct {
	BaseURL string            `json:"baseUri"`
	Headers map[string]string `json:"headers"`
}
