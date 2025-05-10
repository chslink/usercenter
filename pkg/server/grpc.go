package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *Config, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	grpcConf := c.Server.Grpc
	//if grpcConf.Network != "" {
	//	opts = append(opts, grpc.Network(c.Grpc.Network))
	//}
	if grpcConf.Addr != "" {
		opts = append(opts, grpc.Address(grpcConf.Addr))
	}
	//if grpcConf.Timeout != "" {
	//	opts = append(opts, grpc.Timeout(grpcConf.Timeout))
	//}
	srv := grpc.NewServer(opts...)
	//v1.RegisterGreeterServer(srv, greeter)
	return srv
}
