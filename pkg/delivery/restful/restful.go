package restful

import (
	"boyi/pkg/iface"

	gintool "boyi/pkg/Infra/gin"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/fx"
)

type RegisterAPIRouterParams struct {
	fx.In

	E        *gin.Engine
	Handlers []iface.IHandler `group:"handler"`
}

func RegisterAPIRouter(p RegisterAPIRouterParams) {
	p.E.Static("/public", "public")
	for _, h := range p.Handlers {
		h.Register(p.E.Group("/" + h.Version()))
	}

	p.E.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gintool.RegisterDefaultRoute(p.E)
}
