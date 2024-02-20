package redis_worker

import (
	"boyi/pkg/hub"
	"boyi/pkg/iface"

	"boyi/pkg/infra/redis"

	"github.com/bsm/redislock"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		RegisterHandler,
	),
	fx.Invoke(
		redis.InitSubscriptHandler,
	),
)

// Params ...
type Params struct {
	fx.In

	UserSvc    iface.IUserService
	CacheRepo  iface.ICacheRepository
	SupportSvc iface.ISupportService
	Repo       iface.IRepository
	Hub        *hub.Hub
	RedisLock  *redislock.Client
}

// Handler gRPC handler ...
type handler struct {
	userSvc    iface.IUserService
	cacheRepo  iface.ICacheRepository
	supportSvc iface.ISupportService
	repo       iface.IRepository
	hub        *hub.Hub
	redisLock  *redislock.Client
}

// New gRPC 依賴注入
func New(p Params) (*handler, error) {
	h := handler{
		userSvc:    p.UserSvc,
		supportSvc: p.SupportSvc,
		repo:       p.Repo,
		cacheRepo:  p.CacheRepo,
		hub:        p.Hub,
		redisLock:  p.RedisLock,
	}

	return &h, nil
}
