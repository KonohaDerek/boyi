package common

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BaseWhere struct {
	SearchIn         string    `json:"search_in,omitempty"`
	Keyword          string    `json:"keyword,omitempty"`
	CreatedAtLt      time.Time `json:"created_at_lt,omitempty"`
	CreatedAtLte     time.Time `json:"created_at_lte,omitempty"`
	CreatedAtGt      time.Time `json:"created_at_gt,omitempty"`
	CreatedAtGte     time.Time `json:"created_at_gte,omitempty"`
	UpdatedAtLt      time.Time `json:"updated_at_lt,omitempty"`
	UpdatedAtLte     time.Time `json:"updated_at_lte,omitempty"`
	UpdatedAtGt      time.Time `json:"updated_at_gt,omitempty"`
	UpdatedAtGte     time.Time `json:"updated_at_gte,omitempty"`
	Ids              []int64   `json:"ids,omitempty"`
	DeletedAtLt      time.Time `json:"deleted_at_lt,omitempty"`
	DeletedAtLte     time.Time `json:"deleted_at_lte,omitempty"`
	DeletedAtGt      time.Time `json:"deleted_at_gt,omitempty"`
	DeletedAtGte     time.Time `json:"deleted_at_gte,omitempty"`
	RangeField       string    `json:"range_field,omitempty"`
	RangeType        RangeType `json:"range_type,omitempty"`
	LessThan         int64     `json:"less_than,omitempty"`
	LessThanEqual    int64     `json:"less_than_equal,omitempty"`
	GreaterThan      int64     `json:"greater_than,omitempty"`
	GreaterThanEqual int64     `son:"greater_than_equal,omitempty"`
}

type RangeType int32

const (
	RangeType__Unknown RangeType = iota
	RangeType__ByDateTime
	RangeType__ByNumber
)

var AllRangeType = []RangeType{
	RangeType__ByDateTime,
	RangeType__ByNumber,
}

func (e RangeType) IsValid() bool {
	switch e {
	case RangeType__ByDateTime, RangeType__ByNumber:
		return true
	}
	return false
}

func (e RangeType) String() string {
	return string(e)
}

func (where *BaseWhere) Where(db *gorm.DB) *gorm.DB {
	if where.Ids != nil && len(where.Ids) != 0 {
		db = db.Where("id IN (?)", where.Ids)
	}
	if where.SearchIn != "" && where.Keyword != "" {
		fields := strings.Split(where.SearchIn, ",")
		for _, field := range fields {
			field := field
			db = db.Where(fmt.Sprintf("%s like ?", field), "%"+where.Keyword+"%")
		}
	}

	if !where.CreatedAtLt.IsZero() {
		db = db.Where("created_at < ?", where.CreatedAtLt)
	}
	if !where.CreatedAtLte.IsZero() {
		db = db.Where("created_at <= ?", where.CreatedAtLte)
	}
	if !where.CreatedAtGt.IsZero() {
		db = db.Where("created_at > ?", where.CreatedAtGt)
	}
	if !where.CreatedAtGte.IsZero() {
		db = db.Where("created_at >= ?", where.CreatedAtGte)
	}
	if !where.UpdatedAtLt.IsZero() {
		db = db.Where("updated_at < ?", where.UpdatedAtLt)
	}
	if !where.UpdatedAtLte.IsZero() {
		db = db.Where("updated_at <= ?", where.UpdatedAtLte)
	}
	if !where.UpdatedAtGt.IsZero() {
		db = db.Where("updated_at > ?", where.UpdatedAtGt)
	}
	if !where.UpdatedAtGte.IsZero() {
		db = db.Where("updated_at >= ?", where.UpdatedAtGte)
	}
	if !where.DeletedAtLt.IsZero() {
		db = db.Where("deleted_at < ?", where.DeletedAtLt)
	}
	if !where.DeletedAtLte.IsZero() {
		db = db.Where("deleted_at <= ?", where.DeletedAtLte)
	}
	if !where.DeletedAtGt.IsZero() {
		db = db.Where("deleted_at > ?", where.DeletedAtGt)
	}
	if !where.DeletedAtGte.IsZero() {
		db = db.Where("deleted_at >= ?", where.DeletedAtGte)
	}

	if where.RangeField != "" && where.RangeType != 0 {
		switch where.RangeType {
		case RangeType__ByDateTime:
			if where.LessThan != 0 {
				_tmp := time.Unix(where.LessThan, 0)
				db = db.Where(fmt.Sprintf("%s < ?", where.RangeField), _tmp)
			}
			if where.LessThanEqual != 0 {
				_tmp := time.Unix(where.LessThanEqual, 0)
				db = db.Where(fmt.Sprintf("%s <= ?", where.RangeField), _tmp)
			}
			if where.GreaterThan != 0 {
				_tmp := time.Unix(where.GreaterThan, 0)
				db = db.Where(fmt.Sprintf("%s > ?", where.RangeField), _tmp)
			}
			if where.GreaterThanEqual != 0 {
				_tmp := time.Unix(where.GreaterThanEqual, 0)
				db = db.Where(fmt.Sprintf("%s >= ?", where.RangeField), _tmp)
			}
		case RangeType__ByNumber:
			if where.LessThan != 0 {
				db = db.Where(fmt.Sprintf("%s < ?", where.RangeField), where.LessThan)
			}
			if where.LessThanEqual != 0 {
				db = db.Where(fmt.Sprintf("%s <= ?", where.RangeField), where.LessThanEqual)
			}
			if where.GreaterThan != 0 {
				db = db.Where(fmt.Sprintf("%s > ?", where.RangeField), where.GreaterThan)
			}
			if where.GreaterThanEqual != 0 {
				db = db.Where(fmt.Sprintf("%s >= ?", where.RangeField), where.GreaterThanEqual)
			}
		}

	}
	return db
}
