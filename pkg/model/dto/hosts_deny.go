package dto

import (
	"boyi/pkg/model/option/common"
	"time"
)

type HostsDeny struct {
	ID           uint64       `gorm:"primaryKey;autoIncrement;"`                               // id
	IPAddress    string       `gorm:"uniqueIndex:unique_ip_address;type:varchar(40);NOT NULL"` // ip address
	IsEnabled    common.YesNo `gorm:"type:tinyint;NOT NULL;DEFAULT:1"`                         // 是否啟用
	Remark       string       `gorm:"type:varchar(100)"`                                       // 备注
	CreatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL"`                                 // 创建时间
	UpdatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL"`                                 // 更新时间
	CreateUserID uint64       `gorm:""`
	UpdateUserID uint64       `gorm:""`
}

func (HostsDeny) TableName() string {
	return "hosts_deny"
}
