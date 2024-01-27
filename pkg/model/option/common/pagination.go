package common

import "gorm.io/gorm"

type OffsetType int32

const (
	OffsetType__UNKNOWN OffsetType = iota
	OffsetType__Old
	OffsetType__New
)

type Pagination struct {
	// 基礎分頁
	Page       int64 `json:"page,omitempty"`
	PerPage    int64 `json:"per_page,omitempty"`
	TotalCount int64 `json:"total_count,omitempty"`
	// 透過 id 做 offset
	OffsetId     int64      `json:"offset_id,omitempty"`
	Limit        int64      `json:"limit,omitempty"`
	OffsetType   OffsetType `json:"offset_type,omitempty"`
	WithoutCount bool       `json:"without_count,omitempty"`
}

const globalDefaultPerPage = 30

// CheckOrSetDefault 檢查Page值若未設置則設置預設值
func (p *Pagination) CheckOrSetDefault(params ...int64) *Pagination {
	var defaultPerPage int64
	if len(params) >= 1 {
		defaultPerPage = params[0]
	}

	if defaultPerPage <= 0 {
		defaultPerPage = globalDefaultPerPage
	}

	if p.Page == 0 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = defaultPerPage
	}
	return p
}

// LimitAndOffset return limit and offset
func (p *Pagination) LimitAndOffset(db *gorm.DB) *gorm.DB {
	if p.PerPage != 0 || p.Offset() != 0 {
		db = db.Limit(int(p.PerPage)).Offset(int(p.Offset()))
	}
	if p.Limit != 0 {
		db = db.Limit(int(p.Limit))
	}

	if p.OffsetId != 0 && p.OffsetType != 0 {
		switch p.OffsetType {
		case OffsetType__New:
			db = db.Where("id > ?", p.OffsetId)
		case OffsetType__Old:
			db = db.Where("id < ?", p.OffsetId)
		}
	}

	return db
}

// Offset 計算 offset 的值
func (p *Pagination) Offset() int64 {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.PerPage
}
