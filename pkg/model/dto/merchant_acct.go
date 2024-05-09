package dto

import (
	"boyi/pkg/model/enums/types"
	"time"

	"github.com/shopspring/decimal"
)

// MerchantAcct 商戶帳戶
type MerchantAcct struct {
	ID           uint64          `gorm:"autoIncrement;primary_key"`
	Name         string          `db:"name" json:"name"`                   // 名称
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

// TableName return database table name
func (MerchantAcct) TableName() string {
	return "merchant_accts"
}

// MerchantAcctLog 商戶帳戶紀錄
type MerchantAcctLog struct {
	ID                uint64                      `gorm:"autoIncrement;primary_key"`
	Currency          types.Currency              `db:"currency" json:"currency"`                       // 货币
	MerchantID        uint64                      `db:"merchant_id" json:"merchant_id"`                 // 商户 ID
	MerchantName      string                      `db:"merchant_name" json:"merchant_name"`             // 商户名称
	MerchantAcctID    uint64                      `db:"merchant_acct_id" json:"merchant_acct_id"`       // 商户帐户 ID
	ChangeBalance     decimal.Decimal             `db:"change_balance" json:"change_balance"`           // 变动余额
	BeforeBalance     decimal.Decimal             `db:"before_balance" json:"before_balance"`           // 之前余额
	AfterBalance      decimal.Decimal             `db:"after_balance" json:"after_balance"`             // 之后余额
	ChangeBlockAmount decimal.Decimal             `db:"change_block_amount" json:"change_block_amount"` // 变动圈存金额
	BeforeBlockAmount decimal.Decimal             `db:"before_block_amount" json:"before_block_amount"` // 之前圈存金额
	AfterBlockAmount  decimal.Decimal             `db:"after_block_amount" json:"after_block_amount"`   // 之后圈存金额
	ChangeProfit      decimal.Decimal             `db:"change_profit" json:"change_profit"`             // 变动盈利
	BeforeProfit      decimal.Decimal             `db:"before_profit" json:"before_profit"`             // 之前盈利
	AfterProfit       decimal.Decimal             `db:"after_profit" json:"after_profit"`               // 之后盈利
	ChangeBlockProfit decimal.Decimal             `db:"change_block_profit" json:"change_block_profit"` // 变动圈存盈利
	BeforeBlockProfit decimal.Decimal             `db:"before_block_profit" json:"before_block_profit"` // 之前圈存盈利
	AfterBlockProfit  decimal.Decimal             `db:"after_block_profit" json:"after_block_profit"`   // 之后圈存盈利
	OperationType     types.MerchantAcctOperation `db:"operation_type" json:"operation_type"`           // 操作类型
	Remark            string                      `db:"remark" json:"remark"`                           // 备注
	Extra             string                      `db:"extra" json:"extra"`                             // 附加信息
	CreateUserID      uint64                      `gorm:"NOT NULL;DEFAULT:0"`                           // 建立 user id
	CreatedAt         time.Time                   `gorm:"NOT NULL;"`                                    // 创建时间
}

// TableName return database table name
func (MerchantAcctLog) TableName() string {
	return "merchant_acct_logs"
}
