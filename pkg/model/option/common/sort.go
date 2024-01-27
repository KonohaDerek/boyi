package common

import (
	"fmt"

	"gorm.io/gorm"
)

type SortingOrderType int32

const (
	SortingOrderType__UNKNOW SortingOrderType = iota
	SortingOrderType__ASC
	SortingOrderType__DESC
)

type Sorting struct {
	SortField string           `json:"sort_field,omitempty"`
	Type      SortingOrderType `json:"type,omitempty"`
}

// Sort 依單一欄位單一方向排序
func (s *Sorting) Sort(db *gorm.DB) *gorm.DB {
	if len(s.SortField) != 0 && s.Type != 0 {
		var sortOrder string = "DESC"
		if s.Type == SortingOrderType__ASC {
			sortOrder = "ASC"
		}
		db = db.Order(fmt.Sprintf("%s %s", s.SortField, sortOrder))
	}
	return db
}
