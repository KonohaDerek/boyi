package dto

import (
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"time"
)

const (
	// DefaultManagerRoleID 預設 role id, 提供給第一次平台登錄的管理員
	DefaultManagerRoleID = 1
	// DefaultCSRoleID 預設 role id, 提供沒設定角色的客服
	DefaultCSRoleID = 2
)

// Role 角色
type Role struct {
	ID                 uint64            `gorm:"autoIncrement;primary_key"`  // id
	Name               string            `gorm:"type:varchar(100);NOT NULL"` // 角色名稱
	IsEnable           common.YesNo      `gorm:"NOT NULL;DEFAULT:1"`         // 是否開啟
	Authority          Authority         `gorm:"type:json"`                  // 權限內文
	SupportAccountType types.AccountType `gorm:"type:tinyint(4)"`            // 支援帳號類型
	CreatedAt          time.Time         `gorm:"NOT NULL"`                   // 創建時間
	CreateUserID       uint64            `gorm:"type:int(11) NOT NULL"`      // 創建人
	UpdatedAt          time.Time         `gorm:"NOT NULL"`                   // 更新时间
	UpdateUserID       uint64            `gorm:"type:int(11) NOT NULL"`      // 更新人
}

// TableName return database table name
func (s Role) TableName() string {
	return "roles"
}

func dfsToMap(result map[string]bool, m Authority) {
	for k := range m {
		if _, ok := result[k.String()]; ok {
			continue
		}
		result[k.String()] = true
	}
}
