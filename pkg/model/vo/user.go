package vo

type CustomerServiceJoinCount struct {
	UserID uint64 `gorm:"column:user_id"`
	Count  int64  `gorm:"column:count"`
}
