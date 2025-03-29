// pkg/utils/http_client.go
package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// HTTPClient defines the interface for HTTP clients
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error)
}

// httpClient implements HTTPClient
type httpClient struct {
	client *http.Client
	logger Logger
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(timeout time.Duration, logger Logger) HTTPClient {
	return &httpClient{
		client: &http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}

// Do executes an HTTP request
func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// Get executes an HTTP GET request
func (c *httpClient) Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return c.Do(req)
}

// Post executes an HTTP POST request
func (c *httpClient) Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error) {
	var bodyReader io.Reader

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Set content type if not already set
	if _, ok := headers["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.Do(req)
}

