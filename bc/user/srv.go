package user

import (
	"context"

	"usercenter/bc/user/model"
	"usercenter/bc/user/srv"
)

type Srv struct {
}

func (s *Srv) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewSrv() srv.UserSrv {
	return &Srv{}
}
