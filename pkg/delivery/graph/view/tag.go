package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter TagFilterInput) ConvertToOption() option.TagWhereOption {
	var result option.TagWhereOption
	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.Tag != nil {
		result.Tag = filter.Tag.ConvertToDTO()
	}

	return result
}

func (role TagInput) ConvertToDTO() (result dto.Tag) {
	if role.ID != nil {
		result.ID = uint64(*role.ID)
	}
	if role.Name != nil {
		result.Name = *role.Name
	}
	if role.IsEnable != nil {
		result.IsEnable = YesNoToDTO[*role.IsEnable]
	}

	return
}

func (u *Tag) FromDTO(in dto.Tag) *Tag {
	if u == nil {
		u = &Tag{}
	}

	u.ID = in.ID
	u.Name = in.Name
	u.CreatedAt = in.CreatedAt
	u.UpdatedAt = in.UpdatedAt
	u.CreateUserID = in.CreateUserID
	u.UpdateUserID = in.UpdateUserID
	u.RGBHex = in.RGBHex
	u.IsEnable = YesNoFromDTO[in.IsEnable]

	return u
}

func (cols TagUpdateInput) ConvertToOption() (result option.TagUpdateColumn) {
	if cols.Name != nil {
		result.Name = *cols.Name
	}

	if cols.RGBHex != nil {
		result.RGBHex = *cols.RGBHex
	}

	if cols.IsEnable != nil {
		result.IsEnable = YesNoToDTO[*cols.IsEnable]
	}

	return
}

func (input TagCreateInput) ConvertToDTO() (result dto.Tag) {
	result.Name = input.Name
	result.RGBHex = input.RGBHex

	return
}
