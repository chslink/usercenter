package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Server struct {
	Http *http.Server
	Grpc *grpc.Server
}

func NewServer(httpSrv *http.Server, grpcSrv *grpc.Server) *Server {
	return &Server{
		Http: httpSrv,
		Grpc: grpcSrv,
	}
}
