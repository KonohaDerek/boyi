package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type UserTagWhereOption struct {
	UserTag    dto.UserTag       `json:"tag"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *UserTagWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *UserTagWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *UserTagWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.UserTag)

	return db
}
func (where *UserTagWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.UserTag, dto.UserTag{})
}

func (where *UserTagWhereOption) TableName() string {
	return where.UserTag.TableName()
}

func (where *UserTagWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *UserTagWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type UserTagUpdateColumn struct {
	TagID        uint64
	UpdateUserID uint64
}

func (col *UserTagUpdateColumn) Columns() interface{} {
	return col
}
