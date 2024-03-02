package dto

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/option/common"
	"time"

	"gorm.io/plugin/soft_delete"
)

// Merchant 商戶
type Merchant struct {
	ID           uint64                `gorm:"autoIncrement;primary_key"`                                   // id
	Name         string                `gorm:"type:nvarchar(100);NOT NULL;uniqueIndex:udx_merchant_name"`   // 商戶名稱
	DatabaseType db.DatabaseType       `gorm:"type:nvarchar(100);NOT NULL"`                                 // 資料庫類型
	DatabaseDSN  string                `gorm:"type:nvarchar(500);NOT NULL;uniqueIndex:udx_merchant_dsn;"`   // 資料庫連線資訊
	IsEnable     common.YesNo          `gorm:"NOT NULL;DEFAULT:1"`                                          // 是否開啟
	Remark       string                `gorm:"type:nvarchar(500);"`                                         // 備註
	Extra        db.JSON               `gorm:"type:json"`                                                   // 額外項目
	CreatedAt    time.Time             `gorm:"NOT NULL"`                                                    // 創建時間
	CreateUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 創建人
	UpdatedAt    time.Time             `gorm:"NOT NULL"`                                                    // 更新时间
	UpdateUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 更新人
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:udx_merchant_name;uniqueIndex:udx_merchant_dsn;"` // 刪除時間
	DeleteUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 更新人
}

// TableName return database table name
func (s Merchant) TableName() string {
	return "merchants"
}

type MerchantOrigin struct {
	ID           uint64                `gorm:"autoIncrement;primary_key"`                                   // id
	Origin       string                `gorm:"type:nvarchar(100);NOT NULL;uniqueIndex:udx_merchant_origin"` // 域名
	MerchantID   uint64                `gorm:"type:int(11) NOT NULL"`                                       // 商戶ID
	MerchantName string                `gorm:"type:nvarchar(100);NOT NULL"`                                 // 商戶名稱
	IsEnable     common.YesNo          `gorm:"NOT NULL;DEFAULT:1"`                                          // 是否開啟
	Extra        db.JSON               `gorm:"type:json"`                                                   // 額外項目
	Remark       string                `gorm:"type:nvarchar(500);"`                                         // 備註
	CreatedAt    time.Time             `gorm:"NOT NULL"`                                                    // 創建時間
	CreateUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 創建人
	UpdatedAt    time.Time             `gorm:"NOT NULL"`                                                    // 更新时间
	UpdateUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 更新人
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:udx_merchant_origin"`                             // 刪除時間
	DeleteUserID uint64                `gorm:"type:int(11) NOT NULL"`                                       // 刪除人員
}

// TableName return database table name
func (s MerchantOrigin) TableName() string {
	return "merchant_origins"
}
