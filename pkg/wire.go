package pkg

import (
	"github.com/google/wire"
	"github.com/jinzhu/copier"

	"usercenter/config"
	"usercenter/pkg/server"
)

func CopyServerConfig(config *config.Config) (*server.Config, error) {
	var sc server.Config
	if err := copier.Copy(&sc, config); err != nil {
		return nil, err
	}
	return &sc, nil
}

var Provider = wire.NewSet(CopyServerConfig, server.ProviderSet)
