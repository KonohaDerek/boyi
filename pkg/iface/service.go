package iface

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/vo"
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IService interface {
	IMenu
}

type IMenu interface {
	GetParsedMenuTree(c claims.Claims) []*dto.Menu
}
type IAuthService interface {
	Register(ctx context.Context, in vo.RegisterReq) (dto.User, error)
	Login(ctx context.Context, in vo.LoginReq) (claims.Claims, error)
	Logout(ctx context.Context, c claims.Claims) error

	ValidateHostDeny(ctx context.Context) error

	SetClaims() gin.HandlerFunc
	GetClaimsByToken(ctx context.Context, token string) (claims.Claims, error)
	FlushAllCache(ctx context.Context) error

	// RefreshToken 刷新 Token
	RefreshToken(ctx context.Context, claims *claims.Claims) error
	// GetToken 取得 Token
	GetToken(ctx context.Context, token string, claims *claims.Claims) error
	// 驗證 Token
	JwtValidate(ctx context.Context, token string) (*jwt.Token, error)

	// Merchant
	MerchantLogin(ctx context.Context, in vo.LoginReq) (claims.Claims, error)
}

type IUserService interface {
	GetUser(ctx context.Context, opt *option.UserWhereOption) (dto.User, error)
	GetUserByID(ctx context.Context, id uint64) (dto.User, error)
	GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error)
	CreateUser(ctx context.Context, data *dto.User) error
	ListUsers(ctx context.Context, opt *option.UserWhereOption) ([]dto.User, int64, error)
	UpdateUser(ctx context.Context, opt *option.UserWhereOption, col *option.UserUpdateColumn) error
	DeleteUser(ctx context.Context, opt *option.UserWhereOption) error
	CreateTouristUser(ctx context.Context, deviceUID string) (dto.User, error)
	UpsertUserLoginInfo(ctx context.Context, userID uint64) error

	GetUserRole(ctx context.Context, opt *option.UserRoleWhereOption) (dto.UserRole, error)
	CreateUserRole(ctx context.Context, data *dto.UserRole) error
	ListUserRoles(ctx context.Context, opt *option.UserRoleWhereOption) ([]dto.UserRole, int64, error)
	UpdateUserRole(ctx context.Context, opt *option.UserRoleWhereOption, col *option.UserRoleUpdateColumn) error
	DeleteUserRole(ctx context.Context, opt *option.UserRoleWhereOption) error

	GetUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption) (dto.UserWhitelist, error)
	CreateUserWhitelist(ctx context.Context, data *dto.UserWhitelist) error
	ListUserWhitelists(ctx context.Context, opt *option.UserWhitelistWhereOption) ([]dto.UserWhitelist, int64, error)
	DeleteUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption) error
	UpdateUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption, col *option.UserWhitelistUpdateColumn) error

	GetUserTag(ctx context.Context, opt *option.UserTagWhereOption) (dto.UserTag, error)
	CreateUserTag(ctx context.Context, data *dto.UserTag) error
	ListUserTags(ctx context.Context, opt *option.UserTagWhereOption) ([]dto.UserTag, int64, error)
	DeleteUserTag(ctx context.Context, opt *option.UserTagWhereOption) error
	UpdateUserTag(ctx context.Context, opt *option.UserTagWhereOption, col *option.UserTagUpdateColumn) error

	CreateUserLoginHistory(ctx context.Context, data *dto.UserLoginHistory) error
	ListUserLoginHistories(ctx context.Context, opt *option.UserLoginHistoryWhereOption) ([]dto.UserLoginHistory, int64, error)
	UpdateUserLoginHistory(ctx context.Context, opt *option.UserLoginHistoryWhereOption, col *option.UserLoginHistoryUpdateColumn) error
	GetLastUserLoginHistories(ctx context.Context, opt *option.UserLoginHistoryWhereOption, col *dto.UserLoginHistory) (dto.UserLoginHistory, error)
}

type IRoleService interface {
	GetRole(ctx context.Context, opt *option.RoleWhereOption) (dto.Role, error)
	CreateRole(ctx context.Context, data *dto.Role) error
	ListRoles(ctx context.Context, opt *option.RoleWhereOption) ([]dto.Role, int64, error)
	UpdateRole(ctx context.Context, opt *option.RoleWhereOption, col *option.RoleUpdateColumn) error
	DeleteRole(ctx context.Context, opt *option.RoleWhereOption) error
}

type ITagService interface {
	GetTag(ctx context.Context, opt *option.TagWhereOption) (dto.Tag, error)
	CreateTag(ctx context.Context, data *dto.Tag) error
	ListTags(ctx context.Context, opt *option.TagWhereOption) ([]dto.Tag, int64, error)
	UpdateTag(ctx context.Context, opt *option.TagWhereOption, col *option.TagUpdateColumn) error
	DeleteTag(ctx context.Context, opt *option.TagWhereOption) error
}

type IAuditLogService interface {
	GetAuditLog(ctx context.Context, opt *option.AuditLogWhereOption) (dto.AuditLog, error)
	CreateAuditLog(ctx context.Context, data *dto.AuditLog) error
	ListAuditLogs(ctx context.Context, opt *option.AuditLogWhereOption) ([]dto.AuditLog, int64, error)
	UpdateAuditLog(ctx context.Context, opt *option.AuditLogWhereOption, col *option.AuditLogUpdateColumn) error
	DeleteAuditLog(ctx context.Context, opt *option.AuditLogWhereOption) error

	RecordAuditLogForGraphql(ctx context.Context, next graphql.ResponseHandler) *graphql.Response
}

// 商戶服務
type IMercahntService interface {
	// 商戶
	ListMerchants(ctx context.Context, opt *option.MerchantWhereOption) ([]dto.Merchant, int64, error)
	GetMerchant(ctx context.Context, opt *option.MerchantWhereOption) (dto.Merchant, error)
	CreateMerchant(ctx context.Context, data *dto.Merchant) error
	UpdateMerchant(ctx context.Context, opt *option.MerchantWhereOption, col *option.MerchantUpdateColumn) (dto.Merchant, error)
	DeleteMerchant(ctx context.Context, opt *option.MerchantWhereOption) error

	// 商戶域名
	ListMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) ([]dto.MerchantOrigin, int64, error)
	GetMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) (dto.MerchantOrigin, error)
	CreateMerchantOrigin(ctx context.Context, data *dto.MerchantOrigin) error
	UpdateMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption, col *option.MerchantOriginUpdateColumn) (dto.MerchantOrigin, error)
	DeleteMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) error

	// 商戶使用者
	ListUsers(ctx context.Context, opt *option.MerchantUserWhereOption) ([]dto.MerchantUser, int64, error)
	GetUser(ctx context.Context, opt *option.MerchantUserWhereOption) (dto.MerchantUser, error)
	CreateUser(ctx context.Context, data *dto.MerchantUser) error
	UpdateUser(ctx context.Context, opt *option.MerchantUserWhereOption, col *option.MerchantUserUpdateColumn) (dto.MerchantUser, error)
	DeleteUser(ctx context.Context, opt *option.MerchantUserWhereOption) error

	// 取得商戶域名設置
	GetMerchantOriginFromCtx(ctx context.Context) (dto.MerchantOrigin, error)

	// 商戶帳戶
	ListMerchantAccts(ctx context.Context, opt *option.MerchantAcctWhereOption) ([]dto.MerchantAcct, int64, error)
	GetMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption) (dto.MerchantAcct, error)
	CreateMerchantAcct(ctx context.Context, data *dto.MerchantAcct) error
	UpdateMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption, col *option.MerchantAcctUpdateColumn) (dto.MerchantAcct, error)
	DeleteMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption) error
	// 商戶帳戶異動申請
	MerchantAcctChanges(ctx context.Context, opt *option.MerchantAcctWhereOption, col *option.MerchantAcctChangeColumn) (dto.MerchantAcct, error)
	// 商戶帳戶異動紀錄
	ListMerchantAcctLogs(ctx context.Context, opt *option.MerchantAcctLogWhereOption) ([]dto.MerchantAcctLog, int64, error)
}

// ISupportService 一些雜項服務 FAQ ,Platform Setting
type ISupportService interface {
	// CreateUploadURL 預先產生上傳 URL
	CreateUploadURL(ctx context.Context, in []vo.FileInfo, expire time.Duration) ([]vo.FileInfo, error)

	// HostsDeny 建立阻擋網域
	GetHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) (dto.HostsDeny, error)
	ListHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) ([]dto.HostsDeny, int64, error)
	CreateHostsDeny(ctx context.Context, data *dto.HostsDeny) error
	UpdateHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption, col *option.HostsDenyUpdateColumn) (dto.HostsDeny, error)
	DeleteHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) error
	AutoDenyHostWithRule(ctx context.Context, t time.Time, duration time.Duration) error
}
