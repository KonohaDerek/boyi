package middleware

import (
	"context"
	"fmt"
	"time"

	"boyi/pkg/Infra/ctxutil"
	"boyi/pkg/Infra/redis"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type IPRecordFunc func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error

func (f IPRecordFunc) ExtensionName() string {
	return "HostsDenyFunc"
}

func (f IPRecordFunc) Validate(schema graphql.ExecutableSchema) error {
	if f == nil {
		return fmt.Errorf("HostsDenyFunc can not be nil")
	}
	return nil
}

func (f IPRecordFunc) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	return f(ctx, rc)
}

func IPRecordMiddleware(redis redis.Redis) IPRecordFunc {
	return IPRecordFunc(func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
		ip := ctxutil.GetRealIPFromContext(ctx)
		if ip == "" {
			return nil
		}
		key := fmt.Sprintf("ip_count:%s", time.Now().Format("2006-01-02"))
		if err := redis.ZIncrBy(ctx, key, 1, ip).Err(); err != nil {
			log.Error().Msgf("IPRecordMiddleware error: %v", err)
		}

		return nil
	})
}
