package merchant

import (
	"boyi/internal/mock"
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_service_ListMerchants(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *option.MerchantWhereOption
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		total   int64
		wantErr bool
	}{
		{
			name: "normal test case",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				opt: &option.MerchantWhereOption{
					Merchant: dto.Merchant{},
				},
			},
			total:   6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:      tt.fields.repo,
				cacheRepo: tt.fields.cacheRepo,
			}
			_, total, err := s.ListMerchants(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListMerchants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if total != tt.total {
				t.Errorf("service.ListMerchants() total = %v, want %v", total, tt.total)
			}
		})
	}
}

func Test_service_CreateMerchant(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *dto.Merchant
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test case",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				data: &dto.Merchant{
					Name:         "New Merchant Unit Test",
					DatabaseType: db.MySQL,
					DatabaseDSN:  "127.0.0.1",
					IsEnable:     common.YesNo__YES,
					Remark:       "",
					Extra:        db.JSON("{\"key\":\"value\"}"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:      tt.fields.repo,
				cacheRepo: tt.fields.cacheRepo,
			}
			if err := s.CreateMerchant(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.True(t, tt.args.data.ID > 0)
			assert.Equal(t, tt.args.data.Name, "New Merchant Unit Test")
			merchant, err := s.GetMerchant(tt.args.ctx, &option.MerchantWhereOption{Merchant: dto.Merchant{ID: tt.args.data.ID}})
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListMerchants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, merchant.Name, tt.args.data.Name)
			assert.Equal(t, merchant.DatabaseType, tt.args.data.DatabaseType)
			assert.Equal(t, merchant.DatabaseDSN, tt.args.data.DatabaseDSN)
		})
	}
}

func Test_service_UpdateMerchant(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *option.MerchantWhereOption
		col *option.MerchantUpdateColumn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test case",
			args: args{
				ctx: context.Background(),
				opt: &option.MerchantWhereOption{
					Merchant: dto.Merchant{
						Name: "test merchant update",
					},
				},
				col: &option.MerchantUpdateColumn{
					Name:         "test merchant update complated",
					DatabaseType: db.Postgres,
					DatabaseDSN:  "localhost_merchat_update.host",
					IsEnable:     common.YesNo__NO,
					Remark:       "remark",
					Extra:        db.JSON("{\"key\":\"value2\"}"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := suite.svc
			if err := s.UpdateMerchant(tt.args.ctx, tt.args.opt, tt.args.col); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			merchant, err := s.GetMerchant(tt.args.ctx, &option.MerchantWhereOption{Merchant: dto.Merchant{Name: tt.args.col.Name}})
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, merchant.DatabaseType, tt.args.col.DatabaseType)
			assert.Equal(t, merchant.DatabaseDSN, tt.args.col.DatabaseDSN)
			assert.Equal(t, merchant.IsEnable, tt.args.col.IsEnable)
			assert.Equal(t, merchant.Remark, tt.args.col.Remark)
			assert.Equal(t, merchant.Extra, tt.args.col.Extra)
		})
	}
}

func Test_service_ListMerchantOrigins(t *testing.T) {
	type args struct {
		ctx context.Context
		opt *option.MerchantOriginWhereOption
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		total   int64
		wantErr bool
	}{
		{
			name: "normal test case",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				opt: &option.MerchantOriginWhereOption{
					MerchantOrigin: dto.MerchantOrigin{},
				},
			},
			total:   6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:      tt.fields.repo,
				cacheRepo: tt.fields.cacheRepo,
			}
			list, total, err := s.ListMerchantOrigin(tt.args.ctx, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListMerchantOrigin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if total != tt.total {
				t.Errorf("service.ListMerchantOrigin() total = %v, want %v", total, tt.total)
			}

			assert.True(t, len(list) > 0)
		})
	}
}

func Test_service_CreateMerchantOrigin(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *dto.MerchantOrigin
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "normal test case",
			fields: &service{
				repo:      suite.repo,
				cacheRepo: mock.NewCacheRepo(t),
			},
			args: args{
				ctx: context.Background(),
				data: &dto.MerchantOrigin{
					Origin:       "new_origin.com",
					MerchantID:   3,
					MerchantName: "test merchant",
					IsEnable:     common.YesNo__YES,
					Remark:       "",
					Extra:        db.JSON("{\"key\":\"value\"}"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo:      tt.fields.repo,
				cacheRepo: tt.fields.cacheRepo,
			}
			if err := s.CreateMerchantOrigin(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.CreateMerchantOrigin() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.True(t, tt.args.data.ID > 0)
			assert.Equal(t, tt.args.data.Origin, tt.args.data.Origin)
			merchant, err := s.GetMerchantOrigin(tt.args.ctx, &option.MerchantOriginWhereOption{MerchantOrigin: dto.MerchantOrigin{ID: tt.args.data.ID}})
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetMerchantOrigin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, merchant.Origin, tt.args.data.Origin)
			assert.Equal(t, merchant.MerchantID, tt.args.data.MerchantID)
			assert.Equal(t, merchant.MerchantName, tt.args.data.MerchantName)
			assert.Equal(t, merchant.IsEnable, tt.args.data.IsEnable)
			assert.Equal(t, merchant.Remark, tt.args.data.Remark)
			assert.Equal(t, merchant.Extra, tt.args.data.Extra)
		})
	}
}
