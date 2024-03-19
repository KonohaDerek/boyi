package dto

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/option/common"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
	"gorm.io/plugin/soft_delete"
)

type MerchantUser struct {
	ID           uint64                `gorm:"autoIncrement;primary_key"`                                            // id
	Username     string                `gorm:"type:nvarchar(100);DEFAULT:'';NOT NULL;uniqueIndex:udx_merchant_user"` // 用户名(帳號)
	Password     string                `gorm:"type:nvarchar(255);NOT NULL"`                                          // 密碼
	AliasName    string                `gorm:"type:nvarchar(100);DEFAULT:''"`                                        // 別名（顯示用)
	IsEnable     common.YesNo          `gorm:"NOT NULL;DEFAULT:1"`                                                   // 是否開啟
	Extra        db.JSON               `gorm:"type:json"`                                                            // 額外項目
	CreatedAt    time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                             // 创建时间
	UpdatedAt    time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                             // 更新时间
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:udx_merchant_user"`                                        // 刪除時間
	CreateUserID uint64                `gorm:"type:int(11)"`                                                         // 建立者ID
	UpdateUserID uint64                `gorm:"type:int(11)"`                                                         // 更新人
}

// TableName return database table name
func (s MerchantUser) TableName() string {
	return "merchant_user"
}

func (c MerchantUser) Marshal() []byte {
	b, _ := msgpack.Marshal(c)
	return b
}

func (c *MerchantUser) Unmarshal(s string) error {
	if err := msgpack.Unmarshal([]byte(s), &c); err != nil {
		return err
	}
	return nil
}
