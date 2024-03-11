package option

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type MerchantWhereOption struct {
	Merchant   dto.Merchant      `json:"merchant"`
	Pagination common.Pagination `json:"pagination"`
	BaseWhere  common.BaseWhere  `json:"base_where"`
	Sorting    common.Sorting    `json:"sorting"`
}

type MerchantUpdateColumn struct {
	Name         string          // 商戶名稱
	DatabaseType db.DatabaseType // 資料庫類型
	DatabaseDSN  string          // 資料庫連線資訊
	IsEnable     common.YesNo    // 是否開啟
	Remark       string          // 備註
	Extra        db.JSON         // 額外項目
}

func (col *MerchantUpdateColumn) Columns() interface{} {
	return col
}

func (where *MerchantWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.Merchant)

	return db
}
func (where *MerchantWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.Merchant, dto.Merchant{})
}

func (where *MerchantWhereOption) TableName() string {
	return where.Merchant.TableName()
}

func (where *MerchantWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type MerchantOriginWhereOption struct {
	MerchantOrigin dto.MerchantOrigin `json:"merchant_origin"`
	Pagination     common.Pagination  `json:"pagination"`
	BaseWhere      common.BaseWhere   `json:"base_where"`
	Sorting        common.Sorting     `json:"sorting"`
}

type MerchantOriginUpdateColumn struct {
	Origin       string       // 域名
	MerchantID   uint64       // 商戶ID
	MerchantName string       // 商戶名稱
	IsEnable     common.YesNo // 是否開啟
	Remark       string       // 備註
	Extra        db.JSON      // 額外項目
}

func (col *MerchantOriginUpdateColumn) Columns() interface{} {
	return col
}

func (where *MerchantOriginWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantOriginWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantOriginWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.MerchantOrigin)

	return db
}
func (where *MerchantOriginWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.MerchantOrigin, dto.Merchant{})
}

func (where *MerchantOriginWhereOption) TableName() string {
	return where.MerchantOrigin.TableName()
}

func (where *MerchantOriginWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantOriginWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
