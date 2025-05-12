package user

import (
	"context"

	"usercenter/api/common"
	"usercenter/api/user"
	"usercenter/bc/user/srv"
)

type App struct {
	user.UnimplementedUserServiceServer
	userSrv srv.UserSrv
}

func (a *App) GetCurrentUser(ctx context.Context, empty *common.Empty) (*common.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *App) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*common.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewApp(userSrv srv.UserSrv) *App {
	return &App{
		userSrv: userSrv,
	}
}
