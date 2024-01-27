package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type TagWhereOption struct {
	Tag        dto.Tag           `json:"user_tag"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *TagWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *TagWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *TagWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.Tag)

	return db
}
func (where *TagWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.Tag, dto.Tag{})
}

func (where *TagWhereOption) TableName() string {
	return where.Tag.TableName()
}

func (where *TagWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *TagWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type TagUpdateColumn struct {
	Name         string
	RGBHex       string
	UpdateUserID uint64
	IsEnable     common.YesNo
}

func (col *TagUpdateColumn) Columns() interface{} {
	return col
}
