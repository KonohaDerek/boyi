package role

import (
	"boyi/internal/mock"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_service_GetRole(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *option.RoleWhereOption
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		want    dto.Role
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx: context.Background(),
				opt: &option.RoleWhereOption{
					Role: dto.Role{},
				},
			},
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			got, err := s.GetRole(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.Equal(t, got.Name, "预设管理员权限")
		})
	}
}

func Test_service_CreateRole(t *testing.T) {

	type args struct {
		ctx  context.Context
		data *dto.Role
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				data: &dto.Role{
					Name:               "UnitTest",
					IsEnable:           common.YesNo__YES,
					Authority:          dto.GetAllAuthority(),
					SupportAccountType: types.AccountType__Manager,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.CreateRole(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.LessOrEqual(t, uint64(0), tt.args.data.ID)
		})
	}
}

func Test_service_ListRoles(t *testing.T) {

	type args struct {
		ctx context.Context
		opt *option.RoleWhereOption
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		want    []dto.Role
		want1   int64
		wantErr bool
	}{
		{
			name: "normal test",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				opt: &option.RoleWhereOption{
					Role: dto.Role{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			got, got1, err := s.ListRoles(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotNil(t, got)
			assert.LessOrEqual(t, int64(0), got1)
		})
	}
}

func Test_service_UpdateRole(t *testing.T) {

	type args struct {
		ctx  context.Context
		opt  *option.RoleWhereOption
		col  *option.RoleUpdateColumn
		data *dto.Role
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				opt: &option.RoleWhereOption{},
				col: &option.RoleUpdateColumn{
					Name:         "更新測試",
					IsEnable:     common.YesNo__NO,
					UpdateUserID: 1,
				},
				data: &dto.Role{
					Name:               "更新測試新增",
					IsEnable:           common.YesNo__YES,
					Authority:          dto.GetAllAuthority(),
					SupportAccountType: types.AccountType__Manager,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.CreateRole(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.args.opt = &option.RoleWhereOption{
				Role: dto.Role{
					ID: tt.args.data.ID,
				},
			}
			if err := s.UpdateRole(tt.args.ctx, tt.args.opt, tt.args.col); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_service_DeleteRole(t *testing.T) {
	type args struct {
		ctx  context.Context
		opt  *option.RoleWhereOption
		data *dto.Role
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				opt: &option.RoleWhereOption{},
				data: &dto.Role{
					Name:               "刪除測試新增",
					IsEnable:           common.YesNo__YES,
					Authority:          dto.GetAllAuthority(),
					SupportAccountType: types.AccountType__Manager,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.CreateRole(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.args.opt = &option.RoleWhereOption{
				Role: dto.Role{
					ID: tt.args.data.ID,
				},
			}

			if err := s.DeleteRole(tt.args.ctx, tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
