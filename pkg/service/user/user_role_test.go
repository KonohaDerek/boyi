package user

import (
	"boyi/internal/claims"
	"boyi/internal/mock"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"context"
	"testing"
)

func Test_service_CreateUserRole(t *testing.T) {
	var (
		claimsUserID uint64 = 1
		ctxClaims           = claims.SetClaimsToContext(context.Background(), claims.Claims{Id: claimsUserID})
	)
	type args struct {
		ctx  context.Context
		data *dto.UserRole
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx: ctxClaims,
				data: &dto.UserRole{
					UserID:  1,
					RoleID:  2,
					IsAdmin: common.YesNo__YES,
				},
			},
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
				s3Svc:     mock.NewStorageSvc(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.CreateUserRole(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateUserRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
