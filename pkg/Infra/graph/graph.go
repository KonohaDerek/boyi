package graph

import (
	"net/http"
	"time"

	"boyi/pkg/Infra/errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

type Config struct {
	Playground              bool `json:"playground" mapstructure:"playground"`
	IsResponseDump          bool `json:"is_response_dump" mapstructure:"is_response_dump"`
	QueryCache              int  `json:"query_cache" mapstructure:"query_cache"`
	FixedComplexityLimit    int  `json:"fixed_complexity_limit" mapstructure:"fixed_complexity_limit"`
	AutomaticPersistedQuery int  `json:"automatic_persisted_query" mapstructure:"automatic_persisted_query"`
}

func New(cfg *Config, es graphql.ExecutableSchema) *handler.Server {
	gqlSvc := handler.New(es)
	gqlSvc.AroundResponses(GQLResponseLog(cfg))
	if cfg.FixedComplexityLimit == 0 {
		cfg.FixedComplexityLimit = 1000
	}
	gqlSvc.Use(extension.FixedComplexityLimit(cfg.FixedComplexityLimit))
	if cfg.QueryCache == 0 {
		cfg.QueryCache = 1000
	}
	gqlSvc.SetQueryCache(lru.New(cfg.QueryCache))

	if cfg.AutomaticPersistedQuery == 0 {
		cfg.AutomaticPersistedQuery = 100
	}
	gqlSvc.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(cfg.AutomaticPersistedQuery),
	})
	gqlSvc.AddTransport(transport.POST{})
	gqlSvc.AddTransport(transport.MultipartForm{})
	gqlSvc.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	gqlSvc.SetErrorPresenter(errors.GQLErrorPresenter)
	gqlSvc.SetRecoverFunc(GQLRecoverFunc)

	return gqlSvc
}
