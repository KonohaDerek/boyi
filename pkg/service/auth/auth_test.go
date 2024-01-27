package auth

import (
	"boyi/configuration"
	"boyi/internal/claims"
	"boyi/internal/mock"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/vo"
	"context"
	"strings"
	"testing"
	"time"

	"boyi/pkg/Infra/ctxutil"
	"boyi/pkg/Infra/errors"
	"boyi/pkg/Infra/utils/rand"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {

	type args struct {
		ctx context.Context
		req vo.RegisterReq
	}

	tests := []struct {
		name   string
		args   args
		fields *service
	}{
		{
			name: "normal test",
			args: args{
				ctx: ctxutil.ContextWithXUserAgent(ctxutil.ContextWithXDeviceID(context.Background(), "abcdefgTestDeviceID"), "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Mobile/15E148 Safari/604.1"),
				req: vo.RegisterReq{
					Username: strings.Replace(rand.RandomUUID(), "-", "", -1)[0:16],
					Password: "12345678",
				},
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
		},
		{
			name: "normal with empty ctx test",
			args: args{
				ctx: context.Background(),
				req: vo.RegisterReq{
					Username: strings.Replace(rand.RandomUUID(), "-", "", -1)[0:16],
					Password: "12345678",
				},
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			gotUser, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) && !errors.Is(err, errors.ErrResourceAlreadyExists) {
				t.Errorf("service.InitUser() error = %v", err)
				return
			}
			if errors.Is(err, errors.ErrResourceAlreadyExists) {

				assert.Nil(t, err, "get user error")
			}

			assert.NotNil(t, gotUser, "user is nil")
			assert.Equal(t, gotUser.Username, tt.args.req.Username, "username not equal")
			assert.Equal(t, gotUser.AccountType, types.AccountType__Member, "user account type is not Member")
		})
	}

}

func Test_Login(t *testing.T) {
	ctx := log.Logger.WithContext(context.Background())
	type args struct {
		ctx                  context.Context
		req                  vo.LoginReq
		isManagementPlatform bool
	}

	tests := []struct {
		name   string
		args   args
		fields *service
	}{
		{
			name: "login test",
			args: args{
				ctx: ctxutil.ContextWithXUserAgent(ctxutil.ContextWithXDeviceID(ctx, "abcdefgTestDeviceID"), "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Mobile/15E148 Safari/604.1"),
				req: vo.LoginReq{
					Username: "Test1",
					Password: "test1",
				},
				isManagementPlatform: false,
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields

			result, err := s.Login(tt.args.ctx, tt.args.req)

			if err != nil {
				t.Errorf("service.Login() error = %v", err)
				return
			}

			assert.NotNil(t, result, "result is nil")
			assert.NotNil(t, result.Token, "token is nil")
		})
	}
}

func Test_RefreshToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		claims claims.Claims
		ttl    int64
	}
	const refreshTokenTTLSeconds = 3600 * 24 * 30

	// test 結構
	tests := []struct {
		name   string
		args   args
		fields *service
	}{
		{
			name: "refresh token test",
			args: args{
				ctx: ctxutil.ContextWithXUserAgent(ctxutil.ContextWithXDeviceID(context.Background(), "abcdefgTestDeviceID"), "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Mobile/15E148 Safari/604.1"),
				claims: claims.Claims{
					Id:          1,
					Username:    "Test",
					AccountType: uint64(types.AccountType__Member),
					Competences: nil,
					Token:       "eyJhbGciOiJIUzUxMiIsImlhdCI6MTYyOTk0NDQ3OCwiZXhwIjoxNjMyNTM2NDc4fQ.eyJpZCI6MzI2fQ.hmrJOhudCv8ByxOy0njjuMtgP9hz8CuG-rn9ZHt5XhEoJiPsyqjwWyoFVUni0DogRjNOLObnJ7dZgobMKVLiOQ",
					DeviceUid:   "abcdefgTestDeviceID",
					ExpiredAt:   time.Now().AddDate(0, 0, 1).Unix(),
				},
				ttl: refreshTokenTTLSeconds,
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			err := s.RefreshToken(tt.args.ctx, &tt.args.claims)
			if err != nil {
				t.Errorf("service.RefreshToken() error = %v", err)
				return
			}
		})
	}
}

func Test_service_ValidateHostDeny(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  *service
		args    args
		wantErr bool
	}{
		{
			name: "validate host deny test",
			args: args{
				ctx: ctxutil.ContextWithXRealIP(context.Background(), "1.1.1.1"),
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
			wantErr: true,
		},
		{
			name: "validate host deny test",
			args: args{
				ctx: ctxutil.ContextWithXRealIP(context.Background(), "2.2.2.2"),
			},
			fields: &service{
				jwtConfig: &configuration.Jwt{
					Issure:         "localhost",
					Audience:       "localhost",
					SignKey:        "1234567890",
					ExpiresMiubtes: 60,
				},
				userSvc:    mock.NewUserSvc(t),
				supportSvc: mock.NewSupportSvc(t),
				cacheRepo:  mock.NewCacheRepo(t),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			if err := s.ValidateHostDeny(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("service.ValidateHostDeny() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				assert.EqualError(t, err, "IP非法: [401005] IP非法.")
			}

		})
	}
}
