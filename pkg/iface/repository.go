package iface

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
	"database/sql"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IRepository interface {
	GetDB() *gorm.DB
	Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	Get(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	GetLast(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, data Model, scopes ...func(*gorm.DB) *gorm.DB) error
	List(ctx context.Context, tx *gorm.DB, data interface{}, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
	Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error
	Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Count(ctx context.Context, tx *gorm.DB, opt WhereOption) (int64, error)

	BatchInsert(ctx context.Context, tx *gorm.DB, data interface{}) error

	CreateIfNotExists(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	CreateOrUpdate(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, updateCol UpdateColumns) error

	IMerchantRepository

	IUserRepository
}

type IMerchantBaseRepository interface {
	
}

type IMerchantRepository interface {
	GetALLMerchantDB(ctx context.Context) (map[uint64]*gorm.DB, error)
	GetMerchantDB(ctx context.Context, merchantId uint64) (*gorm.DB, error)
	SetMerchantDB(ctx context.Context, merchantId uint64, connectStr string, databaseType db.DatabaseType) error
	DeleteMerchantDB(ctx context.Context, merchantId uint64) error
}

type IUserRepository interface {
	GetUserByID(ctx context.Context, userID uint64) (dto.User, error)
	GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error)
}

type ICacheRepository interface {
	FlushAllCache(ctx context.Context) error

	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) (bool, error)
	SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error)
	SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error
	SetTTL(ctx context.Context, key string, expireAt time.Duration) error
	Scan(ctx context.Context, pattern string) (keys []string, err error)
	Del(ctx context.Context, key string) error
	RPush(ctx context.Context, key string, data interface{}) (total int64, err error)
	LLen(ctx context.Context, key string) (total int64, err error)
	ZAddNX(ctx context.Context, key string, members ...redis.Z) error
	ZCard(ctx context.Context, key string) (int64, error)
	ZPopMin(ctx context.Context, key string, popCount int64) ([]redis.Z, error)
	ZRangeWithScore(ctx context.Context, key string, start, end int64) ([]redis.Z, error)
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error)
	ZRem(ctx context.Context, key string, member interface{}) error
	SetEXWithJson(ctx context.Context, key string, data interface{}, expireAt time.Duration) error
	Publish(ctx context.Context, key string, message interface{}) error

	// 使用者上線清單
	UserOnlineMap(ctx context.Context) (map[string]bool, error)
	// 加入使用者
	AddUserOnline(ctx context.Context, user *dto.User) error
	RemoveUserOnline(ctx context.Context, user *dto.User) error
	UserOnlineMapWithKey(ctx context.Context) (map[string]bool, error)
}

type WhereOption interface {
	Model
	Where(db *gorm.DB) *gorm.DB
	Page(db *gorm.DB) *gorm.DB
	Sort(db *gorm.DB) *gorm.DB
	Preload(db *gorm.DB) *gorm.DB
	IsEmptyWhereOpt() bool
	WithoutCount() bool
}

type Model interface {
	TableName() string
}

type UpdateColumns interface {
	Columns() interface{}
}
