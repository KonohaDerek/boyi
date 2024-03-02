package service

import (
	"boyi/pkg/service/audit_log"
	"boyi/pkg/service/auth"
	"boyi/pkg/service/menu"
	"boyi/pkg/service/merchant"
	"boyi/pkg/service/role"
	"boyi/pkg/service/support"
	"boyi/pkg/service/tag"
	"boyi/pkg/service/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
	menu.Module,
	user.Module,
	role.Module,
	tag.Module,
	audit_log.Module,
	support.Module,
	merchant.Module,
)
