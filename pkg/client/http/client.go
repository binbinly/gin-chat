package http

import (
	"context"
	"net/http"
	"time"
)

// defaultTimeout default http request timeout
const defaultTimeout = 10 * time.Second

// httpSettings http request options
type httpSettings struct {
	headers map[string]string
	cookies []*http.Cookie
	close   bool
	timeout time.Duration
}

//ClientOption HTTPOption configures how we set up the http request
type ClientOption func(s *httpSettings)

// WithHTTPHeader specifies the headers to http request.
func WithHTTPHeader(key, value string) ClientOption {
	return func(s *httpSettings) {
		s.headers[key] = value
	}
}

// WithHTTPCookies specifies the cookies to http request.
func WithHTTPCookies(cookies ...*http.Cookie) ClientOption {
	return func(s *httpSettings) {
		s.cookies = cookies
	}
}

// WithHTTPClose specifies close the connection after
// replying to this request (for servers) or after sending this
// request and reading its response (for clients).
func WithHTTPClose() ClientOption {
	return func(s *httpSettings) {
		s.close = true
	}
}

// WithHTTPTimeout specifies the timeout to http request.
func WithHTTPTimeout(timeout time.Duration) ClientOption {
	return func(s *httpSettings) {
		s.timeout = timeout
	}
}

// Client 定义 http client 接口
type Client interface {
	// Get sends an HTTP get request
	Get(ctx context.Context, reqURL string, options ...ClientOption) ([]byte, error)

	// Post sends an HTTP post request
	Post(ctx context.Context, reqURL string, data map[string]string, options ...ClientOption) ([]byte, error)

	// PostJSON sends an HTTP post request with json
	PostJSON(ctx context.Context, reqURL string, body []byte, options ...ClientOption) ([]byte, error)

	// Upload sends an HTTP post request for uploading media
	Upload(ctx context.Context, reqURL string, form UploadForm, options ...ClientOption) ([]byte, error)
}
