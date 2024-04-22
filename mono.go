package monogo

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

var lookupBasePath = "/v3/lookup"
var accountBasePath = "/v2/accounts"
var telecomBasePath = "/v1/telecom"
var customerBasePath = "/v2/customers"
var telcoAccBasePath = "/accounts"

type MonoClient struct {
	client *http.Client
	req    *http.Request
}

func (m *MonoClient) withMethod(method string) *MonoClient {
	m.req.Method = strings.ToUpper(method)
	return m
}

func (m *MonoClient) withBody(body io.Reader) *MonoClient {
	m.req.Body = io.NopCloser(body)
	return m
}

func (m *MonoClient) withPath(path string) *MonoClient {
	m.req.URL.Path = path
	return m
}

// DefaultMonoClient returns a new instance of MonoClient with default values.
// It takes a secretKey string as a parameter and initializes the client and request objects.
// The client is an HTTP client and the request is an HTTP GET request to the Mono API endpoint for customers.
// The secretKey is used as a header in the request for authentication.
func New(secretKey string) *MonoClient {
	return &MonoClient{
		client: &http.Client{},
		req: &http.Request{
			Method: "GET",
			Body:   nil,
			URL: &url.URL{
				Scheme: "https",
				Host:   "api.withmono.com",
				Path:   "",
			},
			Header: map[string][]string{
				"content-type": {"application/json"},
				"accept":       {"application/json"},
				"mono-sec-key": {secretKey},
			},
		},
	}
}
