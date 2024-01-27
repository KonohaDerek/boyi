package dto

import (
	"boyi/pkg/model/option/common"
	"time"
)

type Tag struct {
	ID           uint64       `gorm:"autoIncrement;primary_key"`
	Name         string       `gorm:"NOT NULL;DEFAULT:''"`
	RGBHex       string       `gorm:"column:rgb_hex;DEFAULT:'';NOT NULL"`
	IsEnable     common.YesNo `gorm:"NOT NULL;DEFAULT:1"`
	CreatedAt    time.Time    `gorm:"NOT NULL;type:TIMESTAMP;"` // 创建时间
	UpdatedAt    time.Time    `gorm:"NOT NULL;type:TIMESTAMP;"` // 更新时间
	CreateUserID uint64       `gorm:"NOT NULL"`
	UpdateUserID uint64       `gorm:"NOT NULL"`
}

func (Tag) TableName() string {
	return "tags"
}
