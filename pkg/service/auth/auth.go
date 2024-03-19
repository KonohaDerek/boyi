package auth

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"boyi/pkg/model/vo"
	"context"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"
	internalGin "boyi/pkg/infra/gin"
	"boyi/pkg/infra/utils/hash"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint64            `json:"userId"`
	Extra  map[string]string `json:"extra"`
}

func (s *service) Register(ctx context.Context, in vo.RegisterReq) (dto.User, error) {
	var (
		user dto.User
		err  error
	)

	// 新增使用者至資料庫
	user.Username = in.Username
	hashByte, err := hash.HashPassword([]byte(in.Password))
	if err != nil {
		return user, err
	}
	user.Password = string(hashByte)
	user.AccountType = in.AccountType

	err = s.userSvc.CreateUser(ctx, &user)
	if err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
		return user, err
	}

	if errors.Is(err, errors.ErrResourceAlreadyExists) {
		return user, errors.WithStack(errors.ErrAccountAlreadyRegistered)
	}

	claims := claims.Claims{
		Id:          user.ID,
		AccountType: uint64(user.AccountType),
		Username:    user.Username,
		AliasName:   user.AliasName,
	}

	if err := s.RefreshToken(ctx, &claims); err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Login(ctx context.Context, in vo.LoginReq) (claims.Claims, error) {
	var (
		user      dto.User
		queryUser option.UserWhereOption
		_claims   claims.Claims
		err       error
	)

	queryUser.User.Username = in.Username
	queryUser.LoadRoles = true
	queryUser.LoadRolesMenu = true
	if user, err = s.userSvc.GetUser(ctx, &queryUser); err != nil && !errors.Is(err, errors.ErrResourceNotFound) {
		return _claims, err
	}

	if errors.Is(err, errors.ErrResourceNotFound) {
		return _claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	if err := hash.CheckPasswordHash([]byte(user.Password), []byte(in.Password)); err != nil {
		return _claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	err = s.userSvc.UpsertUserLoginInfo(ctx, user.ID)
	if err != nil {
		return _claims, err
	}

	_claims = claims.Claims{
		Id:          user.ID,
		AccountType: uint64(user.AccountType),
		Username:    user.Username,
		AliasName:   user.AliasName,
	}

	// 產生 jwt
	if err := s.RefreshToken(ctx, &_claims); err != nil {
		return _claims, err
	}
	return _claims, nil
}

func (s *service) CreateClaimsCache(ctx context.Context, cert *claims.Claims, token string) error {
	// set jwt to cache
	if err := s.cacheRepo.SetEX(ctx, token, cert.Marshal(), time.Duration(s.jwtConfig.Expire)*time.Minute); err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (s *service) SetClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		deviceID := c.GetHeader(internalGin.DeviceID)
		ctx := ctxutil.ContextWithXDeviceID(c.Request.Context(), deviceID)

		realIP := c.GetHeader(string(ctxutil.XRealIP))
		realIPSplit := strings.Split(realIP, ",")
		if len(realIP) != 0 {
			ctx = ctxutil.ContextWithXRealIP(ctx, realIPSplit[0])
		}

		url, err := url.Parse(c.Request.Header.Get("Origin"))
		if err != nil {
			ctx = ctxutil.ContextWithXOrigin(ctx, c.Request.Header.Get("Origin"))
		} else {
			ctx = ctxutil.ContextWithXOrigin(ctx, url.Host)
		}

		c.Request = c.Request.WithContext(ctx)
		if bearerToken == "" || deviceID == "" {
			c.Next()
			return
		}

		tmp := strings.Split(bearerToken, " ")
		if len(tmp) != 2 {
			c.Next()
			return
		}

		var _claims claims.Claims
		if err := s.GetToken(ctx, tmp[1], &_claims); err != nil {
			c.Next()
			return
		}

		ctx = claims.SetClaimsToContext(c.Request.Context(), _claims)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (s *service) Logout(ctx context.Context, c claims.Claims) error {
	// 刪除ClaimsCache

	// 紀錄登出時間
	if err := s.userSvc.UpdateUserLoginHistory(ctx, &option.UserLoginHistoryWhereOption{
		UserLoginHistory: dto.UserLoginHistory{
			DeviceUID: c.DeviceUid,
			Token:     c.Token,
		},
	}, &option.UserLoginHistoryUpdateColumn{
		LogoutAt: time.Now().UTC(),
	}); err != nil {
		log.Ctx(ctx).Err(err)
	}

	return nil
}

func (s *service) GetClaimsByToken(ctx context.Context, token string) (claims.Claims, error) {
	var (
		c claims.Claims
	)

	if err := s.GetToken(ctx, token, &c); err != nil {
		return c, err
	}
	return c, nil
}

func (s *service) FlushAllCache(ctx context.Context) error {
	return s.cacheRepo.FlushAllCache(ctx)
}

func (s *service) RefreshToken(ctx context.Context, claims *claims.Claims) error {
	// 刪除舊的 token
	if claims.Token != "" {
		if err := s.cacheRepo.Del(ctx, claims.Token); err != nil && !errors.Is(err, redis.Nil) {
			return err
		}
	}

	// 使用 golang-jwt 重新產生 token
	// 重新產生 jwt
	expiresAt := time.Now().Add(time.Duration(s.jwtConfig.Expire) * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   claims.Username,
			ExpiresAt: expiresAt,
		},
		UserID: claims.Id,
		Extra:  claims.Extra,
	})

	tokenString, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return err
	}
	claims.Token = tokenString
	claims.ExpiredAt = expiresAt
	// set jwt to cache
	if err := s.CreateClaimsCache(ctx, claims, tokenString); err != nil {
		return err
	}
	return nil
}

func (s *service) ValidateHostDeny(ctx context.Context) error {
	ip := ctxutil.GetRealIPFromContext(ctx)
	if ip == "" {
		return nil
	}
	if hostsDeny, err := s.supportSvc.GetHostsDeny(ctx, &option.HostsDenyWhereOption{
		HostsDeny: dto.HostsDeny{
			IPAddress: ctxutil.GetRealIPFromContext(ctx),
			IsEnabled: common.YesNo__YES,
		},
	}); err == nil && !reflect.DeepEqual(hostsDeny, dto.HostsDeny{}) {
		return errors.Wrapf(errors.ErrHostsDeny, "IP非法")
	}
	return nil
}

func (s *service) GetToken(ctx context.Context, token string, _claims *claims.Claims) error {
	res, err := s.cacheRepo.Get(ctx, token)
	if err != nil && !errors.Is(err, redis.Nil) {
		return errors.ConvertRedisError(err)
	}
	if errors.Is(err, redis.Nil) {
		return errors.Wrapf(errors.ErrTokenUnavailable, "not found key")
	}

	return _claims.Unmarshal(res)
}

func (s *service) JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &authClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return s.jwtConfig.Secret, nil
	})
}

func (s *service) MerchantLogin(ctx context.Context, in vo.LoginReq) (claims.Claims, error) {
	var (
		user        dto.MerchantUser
		_claims     claims.Claims
		merchant_id = ctxutil.GetMerchantIDFromContext(ctx)
		err         error
	)
	tx, err := s.repo.GetMerchantDB(ctx, merchant_id)
	if err != nil {
		return _claims, err
	}
	if err := s.repo.Get(ctx, tx, &user, &option.MerchantUserWhereOption{
		MerchantUser: dto.MerchantUser{
			Username: in.Username,
		},
	}); err != nil {
		return _claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	if err := hash.CheckPasswordHash([]byte(user.Password), []byte(in.Password)); err != nil {
		return _claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	extra := make(map[string]string, 0)
	extra["merchant_id"] = fmt.Sprintf("%d", merchant_id)
	_claims = claims.Claims{
		Id:          user.ID,
		AccountType: uint64(types.AccountType__Merchant),
		Username:    user.Username,
		AliasName:   user.AliasName,
		Extra:       extra,
	}

	// 產生 jwt
	if err := s.RefreshToken(ctx, &_claims); err != nil {
		return _claims, err
	}
	return _claims, nil
}
