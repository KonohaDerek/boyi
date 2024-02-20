package audit_log

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
	"encoding/json"
	"strings"
	"time"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/helper"

	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	// AuditLog 忽略的 method
	ignoreMethod map[string]struct{} = map[string]struct{}{
		"postMessage":            {},
		"createConsultationRoom": {},
		"Login":                  {},
	}
)

// GetAuditLog 取得AuditLog的資訊
func (s *service) GetAuditLog(ctx context.Context, opt *option.AuditLogWhereOption) (dto.AuditLog, error) {
	var (
		result dto.AuditLog
	)
	if err := s.Repo.Get(ctx, nil, &result, opt); err != nil {
		return result, err
	}
	return result, nil
}

// CreateAuditLog 建立AuditLog
func (s *service) CreateAuditLog(ctx context.Context, data *dto.AuditLog) error {
	return s.Repo.Create(ctx, nil, data)
}

// ListAuditLogs 列出AuditLog
func (s *service) ListAuditLogs(ctx context.Context, opt *option.AuditLogWhereOption) ([]dto.AuditLog, int64, error) {
	var (
		result []dto.AuditLog
	)
	total, err := s.Repo.List(ctx, nil, &result, opt)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

// UpdateAuditLog 更新AuditLog
func (s *service) UpdateAuditLog(ctx context.Context, opt *option.AuditLogWhereOption, col *option.AuditLogUpdateColumn) error {
	err := s.Repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAuditLog 刪除AuditLog
func (s *service) DeleteAuditLog(ctx context.Context, opt *option.AuditLogWhereOption) error {
	return s.Repo.Delete(ctx, nil, &dto.AuditLog{}, opt)
}

func (s *service) RecordAuditLogForGraphql(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	go func(ctx context.Context) {
		opCtx := graphql.GetOperationContext(ctx)
		u, err := claims.GetClaims(ctx)
		if err != nil {
			return
		}

		var cancel context.CancelFunc
		_ctx := ctxutil.ContextWithXTraceID(context.Background(), ctxutil.GetTraceIDFromContext(ctx))
		_ctx, cancel = context.WithTimeout(_ctx, 5*time.Second)
		defer helper.Recover(_ctx)
		defer cancel()

		if opCtx == nil {
			return
		}
		if opCtx.OperationName == "IntrospectionQuery" {
			return
		}

		if opCtx.Operation == nil {
			return
		}

		// 只記錄 mutation
		if opCtx.Operation.Operation != ast.Mutation {
			return
		}

		query := []string{}
		var operationName []string
		if opCtx.Operation != nil {
			operationName = make([]string, 0, len(opCtx.Operation.SelectionSet))

			for i := range opCtx.Operation.SelectionSet {
				var (
					count            int
					isStart          bool = false
					end              int
					operationNameEnd int
				)

				position := opCtx.Operation.SelectionSet[i].GetPosition()
				for j := position.Start; j < len(opCtx.RawQuery); j++ {
					if opCtx.RawQuery[j] == '(' || opCtx.RawQuery[j] == '{' {
						if !isStart {
							operationNameEnd = j
						}
						count++
						isStart = true
					} else if opCtx.RawQuery[j] == '}' {
						count--
					}
					if isStart && count == 0 {
						end = j
						break
					}
				}
				if position.Start < operationNameEnd {
					name := strings.TrimSpace(opCtx.RawQuery[position.Start:operationNameEnd])
					if _, ok := ignoreMethod[name]; ok {
						continue
					}
					if strings.ContainsAny(name, "list") ||
						strings.ContainsAny(name, "get") ||
						strings.ContainsAny(name, "post") {
						continue
					}
					operationName = append(operationName, name)
				}

				if end > position.Start {
					query = append(query, strings.ReplaceAll(opCtx.RawQuery[position.Start:end+1], "\n", ""))
				}

			}
		}

		if len(query) > 0 {
			tmp := make(map[string]interface{})
			tmp["variables"] = opCtx.Variables
			tmp["query"] = query
			b, _ := json.Marshal(tmp)

			if err := s.CreateAuditLog(_ctx, &dto.AuditLog{
				UserID:       u.Id,
				Method:       strings.Join(operationName, ", "),
				RequestInput: string(b),
			}); err != nil {
				log.Ctx(_ctx).Error().Msgf("fail to create audit log: %+v", err.Error())
				return
			}
		}
	}(ctx)
	return next(ctx)
}
