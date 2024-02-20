package platform

import (
	"boyi/configuration"
	"boyi/pkg/hub"
	"boyi/pkg/iface"

	"boyi/pkg/infra/storage"

	"go.uber.org/fx"
)

// handler ...
type handler struct {
	appConfig  *configuration.App
	hub        *hub.Hub
	menuSvc    iface.IMenu
	authSvc    iface.IAuthService
	s3Svc      storage.StorageS3
	userSvc    iface.IUserService
	supportSvc iface.ISupportService
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type Params struct {
	fx.In

	AppCfg     *configuration.App
	Hub        *hub.Hub
	MenuSvc    iface.IMenu
	S3Svc      storage.StorageS3
	AuthSvc    iface.IAuthService
	UserSvc    iface.IUserService
	SupportSvc iface.ISupportService
}

func New(p Params) iface.HandlerResult {
	return iface.HandlerResult{
		Handler: &handler{
			appConfig:  p.AppCfg,
			hub:        p.Hub,
			menuSvc:    p.MenuSvc,
			s3Svc:      p.S3Svc,
			userSvc:    p.UserSvc,
			authSvc:    p.AuthSvc,
			supportSvc: p.SupportSvc,
		},
	}
}

func (h *handler) Version() string {
	return "/b/apis/v1"
}
