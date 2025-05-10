package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *Config, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	httpConf := c.Server.Http
	//if httpConf.Network != "" {
	//	opts = append(opts, http.Network(c.Http.Network))
	//}
	if httpConf.Addr != "" {
		opts = append(opts, http.Address(httpConf.Addr))
	}
	//if c.Http.Timeout != nil {
	//	opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	//}
	srv := http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
