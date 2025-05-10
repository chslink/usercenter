package bc

import (
	"github.com/google/wire"

	"usercenter/bc/greeter"
)

var Provider = wire.NewSet(
	greeter.Provider,
	LoadServices,
)
