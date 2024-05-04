package dto

import (
	"boyi/pkg/model/option/common"
	"time"
)

type HostsDeny struct {
	ID           uint64       `gorm:"primaryKey;autoIncrement;comment:'id'"`                                   // id
	IPAddress    string       `gorm:"uniqueIndex:unique_ip_address;type:varchar(40);NOT NULL;comment:'ip 地址'"` // ip address
	IsEnabled    common.YesNo `gorm:"type:tinyint;NOT NULL;DEFAULT:1;comment:'是否啟用(預設啟用)'"`                    // 是否啟用
	Remark       string       `gorm:"type:varchar(100);comment:'備註'"`                                          // 备注
	CreatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL;comment:'創建時間'"`                                  // 创建时间
	UpdatedAt    time.Time    `gorm:"type:TIMESTAMP NOT NULL;comment:'更新時間'"`                                  // 更新时间
	CreateUserID uint64       `gorm:"comment:'建立者ID'"`                                                         // 建立者ID
	CreateUser   string       `gorm:"comment:'建立者'"`                                                           // 建立者
	UpdateUserID uint64       `gorm:"comment:'更新者ID'"`                                                         // 更新人ID
	UpdateUser   string       `gorm:"comment:'更新者'"`                                                           // 更新者
}

func (HostsDeny) TableName() string {
	return "hosts_deny"
}
