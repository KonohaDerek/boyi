package option

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type MerchantAccountWhereOption struct {
	MerchantAccount dto.MerchantAccount `json:"merchant_account"`
	Pagination      common.Pagination   `json:"pagination"`
	BaseWhere       common.BaseWhere    `json:"base_where"`
	Sorting         common.Sorting      `json:"sorting"`
}

type MerchantAccountUpdateColumn struct {
	Password  db.Crypto    // 密碼
	AliasName string       // 別名（顯示用)
	IsEnable  common.YesNo // 是否開啟
	Extra     string       // 額外項目
}

func (col *MerchantAccountUpdateColumn) Columns() interface{} {
	return col
}

func (where *MerchantAccountWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantAccountWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantAccountWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.MerchantAccount)

	return db
}
func (where *MerchantAccountWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.MerchantAccount, dto.MerchantAccount{})
}

func (where *MerchantAccountWhereOption) TableName() string {
	return where.MerchantAccount.TableName()
}

func (where *MerchantAccountWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantAccountWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
