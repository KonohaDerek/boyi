package menu

import (
	"boyi/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo iface.IRepository
}

type Params struct {
	fx.In

	Repo iface.IRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.IMenu {
	s := &service{
		repo: p.Repo,
	}
	return s
}
