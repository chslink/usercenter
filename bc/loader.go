package bc

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "usercenter/api/helloworld/v1"
	"usercenter/bc/greeter"
)

type Loader struct {
}

func LoadServices(httpSrv *http.Server, grpcSrv *grpc.Server, logger log.Logger,
	gs *greeter.Greeter,
) *Loader {
	v1.RegisterGreeterHTTPServer(httpSrv, gs)
	v1.RegisterGreeterServer(grpcSrv, gs)
	return &Loader{}
}
