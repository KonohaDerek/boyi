package merchant

import (
	"boyi/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	repo      iface.IRepository
	cacheRepo iface.ICacheRepository
}

type Params struct {
	fx.In

	Repo      iface.IRepository
	CacheRepo iface.ICacheRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.IMercahntService {
	return &service{
		repo:      p.Repo,
		cacheRepo: p.CacheRepo,
	}
}
