package cache

import (
	"boyi/pkg/iface"

	"boyi/pkg/Infra/redis"

	"go.uber.org/fx"
)

type repository struct {
	redis redis.Redis
}

type Params struct {
	fx.In

	RedisConn redis.Redis
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

var _ iface.ICacheRepository = (*repository)(nil)

func New(p Params) (iface.ICacheRepository, error) {
	repo := &repository{
		redis: p.RedisConn,
	}
	return repo, nil
}
