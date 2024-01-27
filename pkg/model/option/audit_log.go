package option

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type AuditLogWhereOption struct {
	AuditLog   dto.AuditLog      `json:"user_tag"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

func (where *AuditLogWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *AuditLogWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *AuditLogWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.AuditLog)

	return db
}
func (where *AuditLogWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.AuditLog, dto.AuditLog{})
}

func (where *AuditLogWhereOption) TableName() string {
	return where.AuditLog.TableName()
}

func (where *AuditLogWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *AuditLogWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type AuditLogUpdateColumn struct {
	Name         string
	RGBHex       string
	UpdateUserID uint64
	IsEnable     common.YesNo
}

func (col *AuditLogUpdateColumn) Columns() interface{} {
	return col
}
