package platform

import (
	"boyi/internal/graph/platform/generated"
	"boyi/pkg/hub"
	"boyi/pkg/iface"
	"boyi/pkg/middleware"
	"context"
	"net/http"
	"time"

	"boyi/pkg/infra/errors"
	"boyi/pkg/infra/graph"
	"boyi/pkg/infra/redis"
	"boyi/pkg/infra/zlog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userSvc     iface.IUserService
	roleSvc     iface.IRoleService
	tagSvc      iface.ITagService
	menuSvc     iface.IMenu
	authSvc     iface.IAuthService
	auditLogSvc iface.IAuditLogService
	supportSvc  iface.ISupportService

	cacheRepo iface.ICacheRepository

	hub *hub.Hub
}

type Params struct {
	fx.In

	UserSvc     iface.IUserService
	RoleSvc     iface.IRoleService
	TagSvc      iface.ITagService
	MenuSvc     iface.IMenu
	AuthSvc     iface.IAuthService
	AuditLogSvc iface.IAuditLogService
	SupportSvc  iface.ISupportService

	CacheRepo iface.ICacheRepository

	HubSvc *hub.Hub
}

var Module = fx.Options(
	fx.Provide(
		createConfig,
		NewResolver,
	),
	fx.Invoke(
		SetResolver,
	),
)

func NewResolver(p Params) *Resolver {
	return &Resolver{
		userSvc:     p.UserSvc,
		menuSvc:     p.MenuSvc,
		roleSvc:     p.RoleSvc,
		tagSvc:      p.TagSvc,
		authSvc:     p.AuthSvc,
		auditLogSvc: p.AuditLogSvc,
		hub:         p.HubSvc,
		supportSvc:  p.SupportSvc,
		cacheRepo:   p.CacheRepo,
	}
}

func createConfig(r *Resolver) generated.Config {
	c := generated.Config{
		Resolvers:  r,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}

	return c
}

func SetResolver(logCfg *zlog.Config, engine *gin.Engine, cfg generated.Config, authSvc iface.IAuthService, auditLogSvc iface.IAuditLogService, redis redis.Redis) error {
	gqlSvc := handler.New(generated.NewExecutableSchema(cfg))
	gqlSvc.AroundResponses(graph.GQLResponseLog(&graph.Config{}))
	gqlSvc.Use(extension.FixedComplexityLimit(3000))
	gqlSvc.SetQueryCache(lru.New(1000))
	gqlSvc.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	gqlSvc.AddTransport(transport.POST{})
	gqlSvc.AddTransport(transport.MultipartForm{})
	gqlSvc.AddTransport(transport.Websocket{
		ErrorFunc: func(ctx context.Context, err error) {
			log.Ctx(ctx).Debug().Msgf("websocket error: %s", err.Error())
		},
		CloseFunc: func(ctx context.Context, closeCode int) {
			log.Ctx(ctx).Debug().Msgf("websocket close: %d", closeCode)
		},
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			HandshakeTimeout: 15 * time.Second,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
				log.Error().Msgf("ws error: %s", reason)
			},
		},
	})
	gqlSvc.AroundResponses(auditLogSvc.RecordAuditLogForGraphql)
	gqlSvc.SetErrorPresenter(errors.GQLErrorPresenter)
	gqlSvc.SetRecoverFunc(graph.GQLRecoverFunc)

	// 加入封鎖IP判斷
	gqlSvc.Use(middleware.HostDenyMiddleware(authSvc))
	gqlSvc.Use(middleware.IPRecordMiddleware(redis))

	engine.Any("/graph/query", authSvc.SetClaims(), gin.WrapH(gqlSvc))

	if logCfg.Environment != "prod" {
		gqlSvc.Use(extension.Introspection{})
		playGround := playground.Handler("GraphQL playground", "/graph/query")
		engine.Any("/graph/playground", gin.WrapH(playGround))
	}
	return nil
}

const (
	refreshTokenTTLSeconds = 3600 * 24 * 30
)
