package audit_log

import (
	"boyi/pkg/iface"

	"go.uber.org/fx"
)

// service ...
type service struct {
	Repo iface.IRepository
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

func New(p Params) iface.IAuditLogService {
	return &service{
		Repo: p.Repo,
	}
}
