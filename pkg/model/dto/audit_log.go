package dto

import "time"

type AuditLog struct {
	ID           uint64    `gorm:"autoIncrement;primary_key;comment:'id'"` // id
	UserID       uint64    `gorm:"NOT NULL;comment:'使用者ID'"`               // 使用者ID
	UserName     string    `gorm:"NOT NULL;comment:'使用者名稱'"`               // 使用者名稱
	Method       string    `gorm:"NOT NULL;comment:'方法'"`                  // 方法
	RequestInput string    `gorm:"NOT NULL;comment:'請求輸入'"`                // 請求輸入
	CreatedAt    time.Time `gorm:"type:TIMESTAMP;comment:'創建時間'"`          // 創建時間
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
