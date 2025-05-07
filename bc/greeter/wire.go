package greeter

import (
	"usercenter/bc/greeter/repo"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewGreeterUsecase,
)
