package dto

import "time"

type AuditLog struct {
	ID           uint64    `gorm:"autoIncrement;primary_key"`
	UserID       uint64    `gorm:"NOT NULL;"`
	Method       string    `gorm:"NOT NULL;"`
	RequestInput string    `gorm:"NOT NULL;"`
	CreatedAt    time.Time `gorm:"type:TIMESTAMP"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
