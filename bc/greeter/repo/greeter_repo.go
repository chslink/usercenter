package repo

import (
	"context"

	"usercenter/bc/greeter/model"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *model.Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *model.Data, logger log.Logger) GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*Greeter, error) {
	return nil, nil
}
