package middleware

import (
	"boyi/pkg/iface"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type HostsDenyFunc func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error

func (f HostsDenyFunc) ExtensionName() string {
	return "HostsDenyFunc"
}

func (f HostsDenyFunc) Validate(schema graphql.ExecutableSchema) error {
	if f == nil {
		return fmt.Errorf("HostsDenyFunc can not be nil")
	}
	return nil
}

func (f HostsDenyFunc) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	return f(ctx, rc)
}

func HostDenyMiddleware(authSvc iface.IAuthService) HostsDenyFunc {
	return HostsDenyFunc(func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
		if err := authSvc.ValidateHostDeny(ctx); err != nil {
			rc.DisableIntrospection = true
			return gqlerror.WrapPath(nil, err)
		}
		return nil
	})
}
