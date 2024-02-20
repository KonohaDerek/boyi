package support

import (
	"boyi/pkg/iface"

	"boyi/pkg/infra/storage"

	"go.uber.org/fx"
)

// service ...
type service struct {
	s3Svc storage.StorageS3

	repo      iface.IRepository
	cacheRepo iface.ICacheRepository
}

type Params struct {
	fx.In

	S3Svc     storage.StorageS3
	CacheRepo iface.ICacheRepository
	Repo      iface.IRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

func New(p Params) iface.ISupportService {
	return &service{
		repo:      p.Repo,
		cacheRepo: p.CacheRepo,
		s3Svc:     p.S3Svc,
	}
}
