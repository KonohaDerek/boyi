package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter UserTagFilterInput) ConvertToOption() option.UserTagWhereOption {
	var result option.UserTagWhereOption
	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.UserTag != nil {
		result.UserTag = filter.UserTag.ConvertToDTO()
	}

	return result
}

func (userTag UserTagInput) ConvertToDTO() (result dto.UserTag) {
	if userTag.ID != nil {
		result.ID = uint64(*userTag.ID)
	}
	if userTag.UserID != nil {
		result.UserID = uint64(*userTag.UserID)
	}
	if userTag.TagID != nil {
		result.TagID = uint64(*userTag.TagID)
	}

	return
}

func (u *UserTag) FromDTO(in dto.UserTag) *UserTag {
	if u == nil {
		u = &UserTag{}
	}

	u.ID = in.ID
	u.UserID = in.UserID
	u.TagID = in.TagID
	u.CreatedAt = in.CreatedAt
	u.UpdatedAt = in.UpdatedAt
	u.CreateUserID = in.CreateUserID
	u.UpdateUserID = in.UpdateUserID

	u.Tag = &Tag{}
	u.Tag = u.Tag.FromDTO(in.Tag)

	return u
}

func (cols UserTagUpdateInput) ConvertToOption() (result option.UserTagUpdateColumn) {
	if cols.TagID != nil {
		result.TagID = uint64(*cols.TagID)
	}

	return
}

func (u UserTagCreateInput) ConvertToDTO() (result dto.UserTag) {
	result.TagID = uint64(u.TagID)
	result.UserID = uint64(u.UserID)

	return
}
