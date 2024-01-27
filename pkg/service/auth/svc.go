package auth

import (
	"boyi/configuration"
	"boyi/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	jwtConfig  *configuration.Jwt
	userSvc    iface.IUserService
	supportSvc iface.ISupportService

	cacheRepo iface.ICacheRepository
}

type Params struct {
	fx.In

	JwtConfig  *configuration.Jwt
	UserSvc    iface.IUserService
	SupportSvc iface.ISupportService
	CacheRepo  iface.ICacheRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

var _ iface.IAuthService = (*service)(nil)

func New(p Params) iface.IAuthService {
	return &service{
		jwtConfig:  p.JwtConfig,
		userSvc:    p.UserSvc,
		supportSvc: p.SupportSvc,
		cacheRepo:  p.CacheRepo,
	}
}
