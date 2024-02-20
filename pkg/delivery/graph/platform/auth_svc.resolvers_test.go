package platform_test

import (
	"boyi/internal/claims"
	"boyi/pkg/delivery/graph/view"
	"context"
	"fmt"
	"testing"

	"boyi/pkg/infra/ctxutil"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func Test_mutationResolver_Login(t *testing.T) {
	var (
		senderUserID uint64 = 1

		ctxClaims = claims.SetClaimsToContext(suite.ctx, claims.Claims{
			Id: senderUserID,
		})
	)

	ctx := log.Logger.WithContext(context.Background())
	type args struct {
		ctx    context.Context
		in     view.LoginReqInput
		origin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx: ctxutil.ContextWithXOrigin(ctx, "test.com"),
				in: view.LoginReqInput{
					Username: "testLoginMangerUser",
					Password: "testLoginMangerUser",
				},
				origin: "test.com",
			},
		},
		{
			name: "boyi user login with origin test",
			args: args{
				ctx: ctxutil.ContextWithXOrigin(ctx, "boyi.com"),
				in: view.LoginReqInput{
					Username: "testLoginboyiManagerUser",
					Password: "testLoginboyiManagerUser",
				},
				origin: "boyi.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := suite.mutationResolver
			got, err := r.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%+v\n", got)

			user, err := suite.queryResolver.GetUser(ctxClaims, view.UserFilterInput{User: &view.UserInput{
				Username: &tt.args.in.Username,
			}})
			if err != nil {
				assert.Nil(t, err)
				return
			}

			assert.Equal(t, tt.args.in.Username, user.Username)
		})
	}
}
