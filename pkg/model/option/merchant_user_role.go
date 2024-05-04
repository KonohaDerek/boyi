package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type MerchantUserRoleWhereOption struct {
	UserRole   dto.MerchantUserRole `json:"user_role"`
	Pagination common.Pagination    `json:"pagination"`
	BaseWhere  common.BaseWhere     `json:"base_where"`
	Sorting    common.Sorting       `json:"sorting"`

	RoleIDs      []uint64 `json:"role_ids"`
	AuthorityKey string   // 搜尋 role 裡面是否有 menu key

	LoadRole bool
}

type MerchantUserRoleUpdateColumn struct {
	RoleID       uint64
	UpdateUserID uint64 `gorm:"NOT NULL"` // 更新 user id
}

func (col *MerchantUserRoleUpdateColumn) Columns() interface{} {
	return col
}

func (where *MerchantUserRoleWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantUserRoleWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantUserRoleWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.UserRole)

	if len(where.RoleIDs) != 0 {
		db = db.Where("role_id IN (?)", where.RoleIDs)
	}

	if where.AuthorityKey != "" {
		db = db.Where(`role_id IN (
			SELECT id FROM roles WHERE authority LIKE ?
		)`, "%\""+where.AuthorityKey+"\"%")
	}

	return db
}
func (where *MerchantUserRoleWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.UserRole, dto.MerchantUserRole{})
}

func (where *MerchantUserRoleWhereOption) TableName() string {
	return where.UserRole.TableName()
}

func (where *MerchantUserRoleWhereOption) Preload(db *gorm.DB) *gorm.DB {
	db = db.Preload("Role")
	return db
}

func (where *MerchantUserRoleWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
