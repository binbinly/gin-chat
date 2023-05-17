package app

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/binbinly/pkg/logger"
	"github.com/gin-gonic/gin"
)

// 反向代理（Reverse Proxy）

// NewReverseProxy takes target host and creates a reverse proxy
func NewReverseProxy(targetHost string) (*httputil.ReverseProxy, error) {
	remote, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}

	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(targetHost string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy, err := NewReverseProxy(targetHost)
		if err != nil {
			log.Fatalf("failed to create reverse proxy: %v", err)
		}
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			req.Header = r.Header
		}
		proxy.ServeHTTP(w, r)
	}
}

// ProxyGinHandler handles the http request using proxy
func ProxyGinHandler(targetHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy, err := NewReverseProxy(targetHost)
		if err != nil {
			log.Fatalf("failed to create reverse proxy: %v", err)
		}
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			req.Header = c.Request.Header
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		logger.Warnf("Got error while modifying request: %v", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return nil
	}
}
