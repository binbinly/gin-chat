package server

import (
	"gin-chat/internal/router"
	"gin-chat/pkg/app"
	"gin-chat/pkg/transport/http"
)

// NewHTTPServer http server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	r := router.NewRouter(true)

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = r

	return srv
}
