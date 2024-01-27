package dto

import (
	"boyi/pkg/model/option/common"
	"time"
)

type UserWhitelist struct {
	ID           uint64       `gorm:"primaryKey;autoIncrement;"`                            // id
	UserID       uint64       `gorm:"uniqueIndex:unique_user_ip;NOT NULL"`                  // admin 管理员 id
	IPAddress    string       `gorm:"uniqueIndex:unique_user_ip;type:varchar(40);NOT NULL"` // ip address
	IsBind       common.YesNo `gorm:"NOT NULL;DEFAULT:2"`
	CreatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL"` // 创建时间
	UpdatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL"` // 更新时间
	CreateUserID uint64       `gorm:""`
	UpdateUserID uint64       `gorm:""`
}

func (UserWhitelist) TableName() string {
	return "user_whitelists"
}
