package auth

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"boyi/pkg/model/vo"
	"context"
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
	"time"

	"boyi/pkg/Infra/ctxutil"
	"boyi/pkg/Infra/errors"
	internalGin "boyi/pkg/Infra/gin"
	"boyi/pkg/Infra/utils/hash"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint64 `json:"userId"`
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
		claims    claims.Claims
		err       error
	)

	queryUser.User.Username = in.Username
	queryUser.LoadRoles = true
	queryUser.LoadRolesMenu = true
	if user, err = s.userSvc.GetUser(ctx, &queryUser); err != nil && !errors.Is(err, errors.ErrResourceNotFound) {
		return claims, err
	}

	if errors.Is(err, errors.ErrResourceNotFound) {
		return claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	if err := hash.CheckPasswordHash([]byte(user.Password), []byte(in.Password)); err != nil {
		return claims, errors.WithStack(errors.ErrUsernameOrPasswordUnavailable)
	}

	err = s.userSvc.UpsertUserLoginInfo(ctx, user.ID)
	if err != nil {
		return claims, err
	}

	// 產生 jwt
	if err := s.RefreshToken(ctx, &claims); err != nil {
		return claims, err
	}
	return claims, nil
}

func (s *service) CreateClaimsCache(ctx context.Context, cert *claims.Claims, token string) error {
	// set jwt to cache
	if err := s.cacheRepo.SetEX(ctx, token, cert, time.Duration(s.jwtConfig.ExpiresMiubtes)*time.Minute); err != nil {
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

		resp, err := s.GetToken(ctx, tmp[1])
		if err != nil {
			c.Next()
			return
		}

		var _claims claims.Claims

		if err := json.Unmarshal([]byte(resp), &c); err != nil {
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
	resp, err := s.GetToken(ctx, token)
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal([]byte(resp), &c); err != nil {
		return c, errors.WithStack(errors.ErrInternalError)
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
	expiresAt := time.Now().Add(time.Duration(s.jwtConfig.ExpiresMiubtes) * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   claims.Username,
			ExpiresAt: expiresAt,
		},
		UserID: claims.Id,
	})
	tokenString, err := token.SignedString(s.jwtConfig.SignKey)
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

func (s *service) GetToken(ctx context.Context, token string) (string, error) {
	res, err := s.cacheRepo.Get(ctx, token)
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", errors.ConvertRedisError(err)
	}
	if errors.Is(err, redis.Nil) {
		return "", errors.Wrapf(errors.ErrTokenUnavailable, "not found key")
	}

	return res, nil
}
