//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"usercenter/bc"
	"usercenter/config"
	"usercenter/pkg"
)

// wireApp init kratos application.
func wireApp(*config.Config, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(pkg.Provider, bc.Provider, newApp))
}
