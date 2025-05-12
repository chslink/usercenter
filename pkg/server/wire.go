package server

import (
	"github.com/google/wire"
)

// Provider is server providers.
var Provider = wire.NewSet(NewGRPCServer, NewHTTPServer, NewServer)
