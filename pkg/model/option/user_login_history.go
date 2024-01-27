package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type UserLoginHistoryWhereOption struct {
	UserLoginHistory dto.UserLoginHistory `json:"tag"`
	Pagination       common.Pagination    `json:"pagination"`
	BaseWhere        common.BaseWhere     `json:"base_where"`
	Sorting          common.Sorting       `json:"sorting"`
}

func (where *UserLoginHistoryWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *UserLoginHistoryWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *UserLoginHistoryWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.UserLoginHistory)

	return db
}

func (where *UserLoginHistoryWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.UserLoginHistory, dto.UserLoginHistory{})
}

func (where *UserLoginHistoryWhereOption) TableName() string {
	return where.UserLoginHistory.TableName()
}

func (where *UserLoginHistoryWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *UserLoginHistoryWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type UserLoginHistoryUpdateColumn struct {
	LogoutAt time.Time
}

func (cols *UserLoginHistoryUpdateColumn) Columns() interface{} {
	return cols
}
