package srv

import (
	"context"

	"usercenter/bc/user/model"
)

type UserSrv interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}
