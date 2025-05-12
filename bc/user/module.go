package user

import (
	"embed"

	"github.com/google/wire"

	"usercenter/api/user"
	"usercenter/bc/user/repo"
	"usercenter/pkg/module"
	"usercenter/pkg/server"
)

//go:embed sql
var sqlFS embed.FS

type Module struct {
	s     *server.Server
	route user.UserServiceServer
	info  module.Info
}

func (m *Module) Info() module.Info {
	return module.Info{
		Name: "user",
		SQL:  sqlFS,
	}
}

func (m *Module) Start() error {
	user.RegisterUserServiceServer(m.s.Grpc, m.route)
	user.RegisterUserServiceHTTPServer(m.s.Http, m.route)
	return nil
}

func (m *Module) Stop() error {
	return nil
}

func NewModule(app *App, s *server.Server) (*Module, func(), error) {
	m := &Module{
		s:     s,
		route: app,
		info: module.Info{
			Name:     "user",
			SQL:      sqlFS,
			RouteSrv: app,
		},
	}
	return m, func() {
		m.Stop()
	}, nil
}

var Provider = wire.NewSet(
	repo.NewUserRepo,
	NewSrv, NewApp,
	NewModule,
)
