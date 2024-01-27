package dto

import "time"

type UserTag struct {
	ID           uint64    `gorm:"autoIncrement;primary_key"`
	UserID       uint64    `gorm:"uniqueIndex:unique_user_tag_id;NOT NULL"`
	TagID        uint64    `gorm:"uniqueIndex:unique_user_tag_id;NOT NULL"`
	CreateUserID uint64    `gorm:"NOT NULL;DEFAULT:0"` // 建立 user id
	UpdateUserID uint64    `gorm:"NOT NULL;DEFAULT:0"` // 更新 user id
	CreatedAt    time.Time `gorm:"NOT NULL;"`          // 创建时间
	UpdatedAt    time.Time `gorm:"NOT NULL;"`          // 更新时间

	Tag Tag ``
}

//TableName return database table name
func (UserTag) TableName() string {
	return "user_tags"
}
