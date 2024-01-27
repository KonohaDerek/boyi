package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type HostsDenyWhereOption struct {
	HostsDeny  dto.HostsDeny     `json:"hosts_deny"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

type HostsDenyUpdateColumn struct {
	IPAddress    string
	IsEnabled    common.YesNo
	Remark       string
	UpdateUserID uint64
}

func (col *HostsDenyUpdateColumn) Columns() interface{} {
	return col
}

func (where *HostsDenyWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *HostsDenyWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *HostsDenyWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.HostsDeny)

	return db
}
func (where *HostsDenyWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.HostsDeny, dto.HostsDeny{})
}

func (where *HostsDenyWhereOption) TableName() string {
	return where.HostsDeny.TableName()
}

func (where *HostsDenyWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *HostsDenyWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
