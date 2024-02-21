package middleware

import (
	"boyi/pkg/iface"
	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type JWTFunc func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error

func (f JWTFunc) ExtensionName() string {
	return "JWTFunc"
}

func (f JWTFunc) Validate(schema graphql.ExecutableSchema) error {
	if f == nil {
		return fmt.Errorf("HostsDenyFunc can not be nil")
	}
	return nil
}

func (f JWTFunc) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	return f(ctx, rc)
}

func JWTMiddleware(authSvc iface.IAuthService) JWTFunc {
	return JWTFunc(func(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
		auth := ctxutil.GetTokenFromContext(ctx)
		if auth == "" {
			return nil
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := authSvc.JwtValidate(context.Background(), auth)
		if err != nil || !validate.Valid {
			return gqlerror.WrapPath(nil, errors.Wrapf(errors.ErrUnauthorized, "Invalid token"))
		}
		return nil
	})
}
