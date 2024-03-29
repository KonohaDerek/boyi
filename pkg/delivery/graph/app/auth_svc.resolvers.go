package app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"boyi/pkg/delivery/graph/view"
	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/model/vo"
	"context"
	"fmt"
)

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, in view.RegisterReqInput) (*view.LoginResp, error) {
	var resp view.LoginResp
	// 註冊使用者
	user, err := r.authSvc.Register(ctx, in.ConvertToVo(ctx))
	if err != nil {
		return nil, err
	}

	// 登入使用者
	result, err := r.authSvc.Login(ctx, vo.LoginReq{
		Username: user.Username,
		Password: in.Password,
		Origin:   ctxutil.GetOriginFromContext(ctx),
	})
	if err != nil {
		return nil, err
	}
	resp = view.LoginResp{
		Token: result.Token,
	}
	return resp.FromClaims(result), nil
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, in view.LoginReqInput) (*view.LoginResp, error) {
	panic(fmt.Errorf("not implemented: Login - Login"))
}

// Logout is the resolver for the Logout field.
func (r *mutationResolver) Logout(ctx context.Context) (uint64, error) {
	panic(fmt.Errorf("not implemented: Logout - Logout"))
}

// RefreshToken is the resolver for the RefreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context) (*view.RefreshTokenResp, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - RefreshToken"))
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*view.Claims, error) {
	panic(fmt.Errorf("not implemented: Me - Me"))
}
