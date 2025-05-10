package greeter

import (
	"usercenter/bc/greeter/model"
	"usercenter/bc/greeter/repo"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	model.NewData, repo.NewGreeterRepo, repo.NewGreeterUsecase, NewGreeter,
)
