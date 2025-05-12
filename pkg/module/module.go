package module

import (
	"embed"
)

type Info struct {
	Name     string
	SQL      embed.FS
	RouteSrv any
}

type Module interface {
	Info() Info

	Start() error
	Stop() error
}
