package pkg

import (
	"github.com/google/wire"
	"github.com/jinzhu/copier"

	"usercenter/config"
	"usercenter/pkg/db"
	"usercenter/pkg/server"
)

func NewServerConfig(config *config.Config) (*server.Config, error) {
	var sc server.Config
	if err := copier.Copy(&sc, config); err != nil {
		return nil, err
	}
	return &sc, nil
}

func NewDbConfig(config *config.Config) (*db.Config, error) {
	var dc db.Config
	if err := copier.Copy(&dc, config); err != nil {
		return nil, err
	}
	return &dc, nil
}

var Provider = wire.NewSet(
	NewServerConfig, server.Provider,
	NewDbConfig, db.New,
)
