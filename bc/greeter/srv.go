package greeter

import (
	"context"

	v1 "usercenter/api/helloworld/v1"
)

type Greeter struct {
	v1.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(ctx context.Context, request *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: request.Name}, nil
}

func NewGreeter() *Greeter {
	return &Greeter{}
}
