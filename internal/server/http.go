package server

import (
	"gin-chat/pkg/app"
	"gin-chat/pkg/transport/http"
)

// NewHTTPServer http server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	return srv
}
