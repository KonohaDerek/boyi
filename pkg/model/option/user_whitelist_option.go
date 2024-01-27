package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type UserWhitelistWhereOption struct {
	UserWhitelist dto.UserWhitelist `json:"user_whitelist"`
	Pagination    common.Pagination `json:"pagination"`
	BaseWhere     common.BaseWhere  `json:"base_where"`
	Sorting       common.Sorting    `json:"sorting"`
}

type UserWhitelistUpdateColumn struct {
	IPAddress    string
	UpdateUserID uint64
}

func (col *UserWhitelistUpdateColumn) Columns() interface{} {
	return col
}

func (where *UserWhitelistWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *UserWhitelistWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *UserWhitelistWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.UserWhitelist)

	return db
}
func (where *UserWhitelistWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.UserWhitelist, dto.UserWhitelist{})
}

func (where *UserWhitelistWhereOption) TableName() string {
	return where.UserWhitelist.TableName()
}

func (where *UserWhitelistWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *UserWhitelistWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
