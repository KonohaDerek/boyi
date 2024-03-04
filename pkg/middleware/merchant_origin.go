package middleware

import (
	"boyi/pkg/iface"
	"boyi/pkg/infra/ctxutil"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type MerchantOriginFunc func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error

func (f MerchantOriginFunc) ExtensionName() string {
	return "MerchantOriginFunc"
}

func (f MerchantOriginFunc) Validate(schema graphql.ExecutableSchema) error {
	if f == nil {
		return fmt.Errorf("MerchantOriginFunc can not be nil")
	}
	return nil
}

func (f MerchantOriginFunc) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	return f(ctx, rc)
}

func MerchantOriginMiddleware(merchantSvc iface.IMercahntService) MerchantOriginFunc {
	return MerchantOriginFunc(func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
		origin, err := merchantSvc.GetMerchantOriginFromCtx(ctx)
		if err != nil {
			rc.DisableIntrospection = true
			return gqlerror.WrapPath(nil, err)
		}
		_ = ctxutil.ContextWithMerchantID(ctx, origin.MerchantID)
		return nil
	})
}
