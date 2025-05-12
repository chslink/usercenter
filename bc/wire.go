package bc

import (
	"github.com/google/wire"

	"usercenter/bc/user"
)

var Provider = wire.NewSet(
	user.Provider,
	NewLoader,
)
