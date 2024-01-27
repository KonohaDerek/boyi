package view

import (
	"boyi/pkg/model/option/common"
	"time"
)

func (base *BaseFilterInput) ConvertToBaseWhere() (result common.BaseWhere) {
	if base == nil {
		return
	}

	result.Ids = make([]int64, len(base.IDs))
	for i := range base.IDs {
		result.Ids[i] = int64(base.IDs[i])
	}

	if base.CreatedAtGt != nil && *base.CreatedAtGt != 0 {
		result.CreatedAtGt = time.Unix(int64(*base.CreatedAtGt), 0)
	}

	if base.CreatedAtGte != nil && *base.CreatedAtGte != 0 {
		result.CreatedAtGte = time.Unix(int64(*base.CreatedAtGte), 0)
	}

	if base.CreatedAtLt != nil && *base.CreatedAtLt != 0 {
		result.CreatedAtLt = time.Unix(int64(*base.CreatedAtLt), 0)
	}

	if base.CreatedAtLte != nil && *base.CreatedAtLte != 0 {
		result.CreatedAtLte = time.Unix(int64(*base.CreatedAtLte), 0)
	}

	if base.SearchIn != nil && *base.SearchIn != "" {
		result.SearchIn = *base.SearchIn
	}

	if base.Keyword != nil && *base.Keyword != "" {
		result.Keyword = *base.Keyword
	}

	if base.RangeField != nil {
		result.RangeField = *base.RangeField
	}

	if base.RangeType != nil {
		result.RangeType = RangeTypeToDTO[*base.RangeType]
	}

	if base.LessThan != nil {
		result.LessThan = *base.LessThan
	}
	if base.LessThanEqual != nil {
		result.LessThanEqual = *base.LessThanEqual
	}
	if base.GreaterThan != nil {
		result.GreaterThan = *base.GreaterThan
	}
	if base.GreaterThanEqual != nil {
		result.GreaterThanEqual = *base.GreaterThanEqual
	}

	return
}

func (base *PaginationInput) ConvertToPagination() (result common.Pagination) {
	if base == nil {
		return
	}

	if base.Page != nil && *base.Page != 0 {
		result.Page = int64(*base.Page)
	}

	if base.PerPage != nil && *base.PerPage != 0 {
		result.PerPage = int64(*base.PerPage)
	}

	if base.Limit != nil && *base.Limit != 0 {
		result.Limit = int64(*base.Limit)
	}

	if base.OffsetID != nil && *base.OffsetID != 0 {
		result.OffsetId = int64(*base.OffsetID)
	}

	if base.OffsetType != nil && *base.OffsetType != "" {
		result.OffsetType = OffsetTypeToDTO[*base.OffsetType]
	}

	return
}

func (base *BaseFilterInput) ConvertToSorting() (result common.Sorting) {
	if base == nil {
		return
	}

	if base.SortType != nil {
		result.Type = SortTypeToDTO[*base.SortType]
	}

	if base.SortField != nil && *base.SortField != "" {
		result.SortField = *base.SortField
	}

	return
}

func (yn *YesNo) ConvertToCommon() (result common.YesNo) {
	if yn == nil {
		return
	}
	switch *yn {
	case YesNoNo:
		result = common.YesNo__NO
	case YesNoYes:
		result = common.YesNo__YES
	}
	return
}

var (
	YesNoFromDTO = map[common.YesNo]YesNo{
		common.YesNo__YES: YesNoYes,
		common.YesNo__NO:  YesNoNo,
	}
	YesNoToDTO = map[YesNo]common.YesNo{
		YesNoYes: common.YesNo__YES,
		YesNoNo:  common.YesNo__NO,
	}

	SortTypeToDTO = map[SortType]common.SortingOrderType{
		SortTypeAsc:  common.SortingOrderType__ASC,
		SortTypeDesc: common.SortingOrderType__DESC,
	}
	OffsetTypeToDTO = map[OffsetType]common.OffsetType{
		OffsetTypeOld: common.OffsetType__Old,
		OffsetTypeNew: common.OffsetType__New,
	}
	RangeTypeToDTO = map[RangeType]common.RangeType{
		RangeTypeByDateTime: common.RangeType__ByDateTime,
		RangeTypeByNumber:   common.RangeType__ByNumber,
	}
)
