package option

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option/common"
	"reflect"

	"gorm.io/gorm"
)

type MerchantUserWhereOption struct {
	MerchantUser dto.MerchantUser  `json:"merchant_account"`
	Pagination   common.Pagination `json:"pagination"`
	BaseWhere    common.BaseWhere  `json:"base_where"`
	Sorting      common.Sorting    `json:"sorting"`
}

type MerchantUserUpdateColumn struct {
	AliasName    string       // 別名（顯示用)
	IsEnable     common.YesNo // 是否開啟
	Extra        db.JSON      // 額外項目
	UpdateUserID uint64       // 更新人
}

func (col *MerchantUserUpdateColumn) Columns() interface{} {
	return col
}

func (where *MerchantUserWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantUserWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantUserWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.MerchantUser)

	return db
}
func (where *MerchantUserWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.MerchantUser, dto.MerchantUser{})
}

func (where *MerchantUserWhereOption) TableName() string {
	return where.MerchantUser.TableName()
}

func (where *MerchantUserWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantUserWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
