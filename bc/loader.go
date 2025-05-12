package bc

import (
	"context"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"

	"usercenter/bc/user"
	"usercenter/pkg/module"
	"usercenter/pkg/server"
)

type Loader struct {
	s       *server.Server
	db      *bun.DB
	mux     sync.Mutex
	modules []module.Module
}

func (l *Loader) AddModule(m ...module.Module) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.modules = append(l.modules, m...)
}

func (l *Loader) Migrate() error {
	mig := migrate.NewMigrations()
	for _, m := range l.modules {
		info := m.Info()
		if err := mig.Discover(info.SQL); err != nil {
			return err
		}
	}
	ctx := context.Background()
	m := migrate.NewMigrator(l.db, mig)
	if err := m.Init(ctx); err != nil {
		return err
	}
	if _, err := m.Migrate(ctx); err != nil {
		return err
	}
	return nil
}

func (l *Loader) StartService() {
	for _, m := range l.modules {
		if err := m.Start(); err != nil {
			return
		}
	}
}

func NewLoader(s *server.Server, db *bun.DB,
	um *user.Module,
) (*Loader, func(), error) {
	l := &Loader{
		s:       s,
		db:      db,
		modules: make([]module.Module, 0),
	}
	l.AddModule(um)
	return l, func() {}, nil
}
