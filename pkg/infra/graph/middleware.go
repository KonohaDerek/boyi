package graph

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"

	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
)

func BuildGraphqlField(ctx context.Context, graphqlCtx *graphql.OperationContext) (operationName, query []string) {
	if graphqlCtx.OperationName == "IntrospectionQuery" {
		return
	}
	if graphqlCtx == nil {
		return
	}
	query = []string{}
	operationName = []string{}
	if graphqlCtx.Operation != nil {
		operationName = make([]string, 0, len(graphqlCtx.Operation.SelectionSet))
		for i := range graphqlCtx.Operation.SelectionSet {
			var (
				count            int
				isStart          bool = false
				end              int
				operationNameEnd int
			)

			position := graphqlCtx.Operation.SelectionSet[i].GetPosition()
			for j := position.Start; j < len(graphqlCtx.RawQuery); j++ {
				if graphqlCtx.RawQuery[j] == '(' || graphqlCtx.RawQuery[j] == '{' {
					if !isStart {
						operationNameEnd = j
					}
					count++
					isStart = true
				} else if graphqlCtx.RawQuery[j] == '}' {
					count--
				}
				if isStart && count == 0 {
					end = j
					break
				}
			}
			if end > position.Start {
				query = append(query, strings.ReplaceAll(graphqlCtx.RawQuery[position.Start:end+1], "\n", ""))
			}

			if position.Start < operationNameEnd {
				operationName = append(operationName, strings.TrimSpace(graphqlCtx.RawQuery[position.Start:operationNameEnd]))
			}
		}
	}

	return
}

func GQLResponseLog(cfg *Config) graphql.ResponseMiddleware {
	return func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		graphqlCtx := graphql.GetOperationContext(ctx)
		startedAt := time.Now()

		operatorName, query := BuildGraphqlField(ctx, graphqlCtx)

		logger := log.With().
			Str("operation_name", strings.Join(operatorName, ", ")).
			Str("trace_id", ctxutil.GetTraceIDFromContext(ctx)).
			Str("device_id", ctxutil.GetDeviceUIDFromContext(ctx)).
			Str("origin", ctxutil.GetOriginFromContext(ctx)).
			Logger()

		if graphqlCtx.Operation != nil && graphqlCtx.Operation.Operation == ast.Subscription {
			return next(logger.WithContext(ctx))
		}

		resp := next(logger.WithContext(ctx))

		if len(query) != 0 {
			zerologCtx := logger.With()
			if resp != nil {
				zerologCtx = zerologCtx.Str("response", string(resp.Data))
			}
			if graphqlCtx.Operation != nil {
				zerologCtx = zerologCtx.Str("operation", string(graphqlCtx.Operation.Operation))
			}
			if len(graphqlCtx.Variables) != 0 {
				zerologCtx = zerologCtx.Interface("variables", graphqlCtx.Variables)
			}

			zerologCtx = zerologCtx.
				Bool("is_graphql", true).
				Time("start_time", startedAt).
				Bool("access_log", true).
				Float64("cost.sec", time.Since(startedAt).Seconds())

			logger = zerologCtx.Logger()
			logger.Debug().Msgf("%s { %s ", graphqlCtx.Operation.Operation, strings.Join(query, ", "))
		}
		return resp
	}
}

// GQLRecoverFunc ...
func GQLRecoverFunc(ctx context.Context, err interface{}) error {
	path := graphql.GetPath(ctx)
	var msg string
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		msg += fmt.Sprintf("%s:%d\n", file, line)
	}
	logger := log.Ctx(ctx).
		With().
		Str("path", path.String()).
		Interface("err", err).Logger()

	logger.Error().Msgf("panic\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", msg)

	_err, ok := err.(error)
	if ok {
	} else {
		debug.PrintStack()
		_err = errors.ErrInternalError
	}
	return _err
}
