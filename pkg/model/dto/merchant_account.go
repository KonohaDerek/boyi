package dto

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/option/common"
	"time"

	"gorm.io/plugin/soft_delete"
)

type MerchantAccount struct {
	ID           uint64                `gorm:"autoIncrement;primary_key"`                                               // id
	Username     string                `gorm:"type:nvarchar(100);DEFAULT:'';NOT NULL;uniqueIndex:udx_merchant_account"` // 用户名(帳號)
	Password     db.Crypto             `gorm:"type:nvarchar(255);NOT NULL"`                                             // 密碼
	AliasName    string                `gorm:"type:nvarchar(100);DEFAULT:''"`                                           // 別名（顯示用)
	IsEnable     common.YesNo          `gorm:"NOT NULL;DEFAULT:1"`                                                      // 是否開啟
	Extra        db.JSON               `gorm:"type:json"`                                                               // 額外項目
	CreatedAt    time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                                // 创建时间
	UpdatedAt    time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                                // 更新时间
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:udx_merchant_account"`                                        // 刪除時間
	UpdateUserID uint64                `gorm:"type:int(11)"`                                                            // 更新人
	CreateUserID uint64                `gorm:"type:int(11)"`                                                            // 建立者ID
}

// TableName return database table name
func (s MerchantAccount) TableName() string {
	return "merchant_account"
}
