package user

import (
	"boyi/pkg/iface"

	"boyi/pkg/Infra/qqzeng_ip"
	"boyi/pkg/Infra/storage"

	"github.com/bsm/redislock"

	"go.uber.org/fx"
)

// service ...
type service struct {
	redisLock *redislock.Client
	repo      iface.IRepository
	cacheRepo iface.ICacheRepository
	ipRepo    qqzeng_ip.IPSearch

	s3Svc storage.StorageS3

	supportSvc iface.ISupportService
}

type Params struct {
	fx.In

	RedisLock  *redislock.Client
	Repo       iface.IRepository
	CacheRepo  iface.ICacheRepository
	IPRepo     qqzeng_ip.IPSearch
	S3Svc      storage.StorageS3
	SupportSvc iface.ISupportService
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

var _ iface.IUserService = (*service)(nil)

func New(p Params) (iface.IUserService, error) {
	s := &service{
		redisLock:  p.RedisLock,
		repo:       p.Repo,
		cacheRepo:  p.CacheRepo,
		ipRepo:     p.IPRepo,
		s3Svc:      p.S3Svc,
		supportSvc: p.SupportSvc,
	}

	return s, nil
}
