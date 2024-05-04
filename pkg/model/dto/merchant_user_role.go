package dto

import (
	"time"
)

// UserRole admin 管理员权限
type MerchantUserRole struct {
	ID           uint64    `gorm:"autoIncrement;primary_key"`
	UserID       uint64    `gorm:"uniqueIndex:unique_user_role_id;DEFAULT:0;NOT NULL"` // admin 管理员 id
	RoleID       uint64    `gorm:"uniqueIndex:unique_user_role_id;DEFAULT:0;NOT NULL"` // 角色 id
	CreateUserID uint64    `gorm:"NOT NULL;DEFAULT:0"`                                 // 建立 user id
	UpdateUserID uint64    `gorm:"NOT NULL;DEFAULT:0"`                                 // 更新 user id
	CreatedAt    time.Time `gorm:"NOT NULL;"`                                          // 创建时间
	UpdatedAt    time.Time `gorm:"NOT NULL;"`                                          // 更新时间

	Role Role `gorm:"PRELOAD:false;foreignKey:RoleID"`
}

// TableName return database table name
func (MerchantUserRole) TableName() string {
	return "merchant_user_roles"
}
