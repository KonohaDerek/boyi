package option

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// MerchantAcct 商戶帳戶
type MerchantAcct struct {
	ID           uint64          `gorm:"autoIncrement;primary_key"`
	Currency     types.Currency  `db:"currency" json:"currency"`           // 货币
	MerchantID   uint64          `db:"merchant_id" json:"merchant_id"`     // 商户 ID
	MerchantName string          `db:"merchant_name" json:"merchant_name"` // 商户名称
	Balance      decimal.Decimal `db:"balance" json:"balance"`             // 余额
	BlockAmount  decimal.Decimal `db:"block_amount" json:"block_amount"`   // 圈存金额
	Profit       decimal.Decimal `db:"profit" json:"profit"`               // 盈利
	BlockProfit  decimal.Decimal `db:"block_profit" json:"block_profit"`   // 圈存盈利
	Extra        string          `db:"extra" json:"extra"`                 // 附加信息
	CreateUserID uint64          `gorm:"NOT NULL;DEFAULT:0"`               // 建立 user id
	UpdateUserID uint64          `gorm:"NOT NULL;DEFAULT:0"`               // 更新 user id
	CreatedAt    time.Time       `gorm:"NOT NULL;"`                        // 创建时间
	UpdatedAt    time.Time       `gorm:"NOT NULL;"`                        // 更新时间
}

type MerchantAcctWhereOption struct {
	MerchantAcct dto.MerchantAcct  `json:"merchant_acct"`
	Pagination   common.Pagination `json:"pagination"`
	BaseWhere    common.BaseWhere  `json:"base_where"`
	Sorting      common.Sorting    `json:"sorting"`
}

func (where *MerchantAcctWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantAcctWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantAcctWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.MerchantAcct)

	return db
}
func (where *MerchantAcctWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.MerchantAcct, dto.MerchantAcct{})
}

func (where *MerchantAcctWhereOption) TableName() string {
	return where.MerchantAcct.TableName()
}

func (where *MerchantAcctWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantAcctWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}

type MerchantAcctUpdateColumn struct {
	Extra        db.JSON // 額外項目
	UpdateUserID uint64  // 更新人
}

func (col *MerchantAcctUpdateColumn) Columns() interface{} {
	return col
}

type MerchantAcctChangeColumn struct {
	Currnecy    types.Currency               // 货币
	Balance     decimal.Decimal              // 余额
	BlockAmount decimal.Decimal              // 圈存金额
	Profit      decimal.Decimal              // 盈利
	BlockProfit decimal.Decimal              // 圈存盈利
	ChangeType  types.MerchantAcctChangeType // 異動類型
	Remark      string                       // 備註
}

func (col *MerchantAcctChangeColumn) Columns() interface{} {
	return col
}

// 商戶帳戶紀錄
type MerchantAcctLogWhereOption struct {
	MerchantAcctLog dto.MerchantAcctLog `json:"merchant_acct_log"`
	Pagination      common.Pagination   `json:"pagination"`
	BaseWhere       common.BaseWhere    `json:"base_where"`
	Sorting         common.Sorting      `json:"sorting"`
}

func (where *MerchantAcctLogWhereOption) Page(db *gorm.DB) *gorm.DB {
	return where.Pagination.LimitAndOffset(db)
}

func (where *MerchantAcctLogWhereOption) Sort(db *gorm.DB) *gorm.DB {
	return where.Sorting.Sort(db)
}

func (where *MerchantAcctLogWhereOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Scopes(where.BaseWhere.Where)
	db = db.Where(where.MerchantAcctLog)

	return db
}
func (where *MerchantAcctLogWhereOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.MerchantAcctLog, dto.MerchantAcct{})
}

func (where *MerchantAcctLogWhereOption) TableName() string {
	return where.MerchantAcctLog.TableName()
}

func (where *MerchantAcctLogWhereOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *MerchantAcctLogWhereOption) WithoutCount() bool {
	return where.Pagination.WithoutCount
}
