package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type UserWhereOption struct {
	User       dto.User          `json:"user"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`

	IDs           []uint64
	AccountTypeIn []types.AccountType
	TagIDs        []uint64 `json:"tag_ids"`
	RoleIDs       []uint64 `json:"role_ids"`

	LoadWhitelists bool `json:"load_whitelists"`
	LoadRoles      bool `json:"load_roles"`
	LoadRolesMenu  bool `json:"load_roles_menu"`
	LoadTag        bool `json:"load_tag"`
}

func (where *UserWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *UserWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *UserWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.User)

	if len(where.IDs) != 0 {
		db = db.Where("id IN (?)", where.IDs)
	}

	if len(where.AccountTypeIn) != 0 {
		db = db.Where("account_type IN (?)", where.AccountTypeIn)
	}

	return db
}

func (where *UserWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.User, dto.User{})
}

func (where *UserWhereOption) TableName() string {
	return where.User.TableName()
}

func (where *UserWhereOption) Preload(db *gorm.DB) *gorm.DB {
	if where.LoadWhitelists {
		db = db.Preload("Whitelists")
	}
	if where.LoadRoles {
		db = db.Preload("Roles")
	}
	if where.LoadRolesMenu {
		db = db.Preload("Roles.Role")
	}
	if where.LoadTag {
		db = db.Preload("Tags")
	}

	return db
}

func (where *UserWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type UserUpdateColumn struct {
	Username      string
	Password      string
	AliasName     string
	UpdateUserID  uint64
	Email         string
	Area          string
	Notes         string
	AvatarContent *graphql.Upload `gorm:"-"`
	AvatarKey     dto.FileKey
	AccountType   types.AccountType
	LastLoginAt   time.Time // 最後登陸時間
	LastLoginIP   string    // 最後登陸 IP
}

func (cols *UserUpdateColumn) Columns() interface{} {
	return cols
}
