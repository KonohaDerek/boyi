package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type RoleWhereOption struct {
	Role       dto.Role          `json:"role"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *RoleWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *RoleWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *RoleWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.Role)

	return db
}
func (where *RoleWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.Role, dto.Role{})
}

func (where *RoleWhereOption) TableName() string {
	return where.Role.TableName()
}

func (where *RoleWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *RoleWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type RoleUpdateColumn struct {
	Name               string
	SupportAccountType types.AccountType
	Authority          dto.Authority
	IsEnable           common.YesNo
	UpdateUserID       uint64
}

func (col *RoleUpdateColumn) Columns() interface{} {
	return col
}
