package dto

import (
	"boyi/pkg/model/enums/common"
	"time"
)

type UserLoginHistory struct {
	ID                 uint64          `gorm:"primaryKey"`
	UserID             uint64          `gorm:"type:int(11)"`
	IPAddress          string          `gorm:"type:varchar(40)"`
	Country            string          `gorm:"column:country;type:varchar(30)"`
	AdministrativeArea string          `gorm:"column:administrative_area;type:varchar(30)"`
	DeviceOS           common.DeviceOS `gorm:"column:device_os;type:smallint"`
	CreatedAt          time.Time       `gorm:"type:TIMESTAMP"`
	LogoutAt           *time.Time      `gorm:"type:TIMESTAMP"`
	DeviceUID          string          `gorm:"type:varchar(100)"`
	Token              string          `gorm:"type:varchar(500)"`
	UserAgent          string          `gorm:"type:varchar(500)"`
}

func (UserLoginHistory) TableName() string {
	return "user_login_histories"
}
