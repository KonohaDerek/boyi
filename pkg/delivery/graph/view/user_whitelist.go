package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter UserWhitelistFilterInput) ConvertToOption() option.UserWhitelistWhereOption {
	var result option.UserWhitelistWhereOption

	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.UserWhitelist != nil {
		result.UserWhitelist = filter.UserWhitelist.ConvertToDTO()
	}

	return result
}

func (whitelist UserWhitelistInput) ConvertToDTO() (result dto.UserWhitelist) {
	if whitelist.ID != nil {
		result.ID = uint64(*whitelist.ID)
	}
	if whitelist.IPAddress != nil {
		result.IPAddress = *whitelist.IPAddress
	}
	if whitelist.UserID != nil {
		result.UserID = uint64(*whitelist.UserID)
	}

	return
}

func (u *UserWhitelist) FromDTO(in dto.UserWhitelist) *UserWhitelist {
	if u == nil {
		u = &UserWhitelist{}
	}

	u.ID = in.ID
	u.IPAddress = in.IPAddress
	u.UserID = in.UserID
	u.CreatedAt = in.CreatedAt

	return u
}

func (u UserWhitelistUpdateInput) ConvertToOption() option.UserWhitelistUpdateColumn {
	var result option.UserWhitelistUpdateColumn

	if u.IPAddress != nil {
		result.IPAddress = *u.IPAddress
	}

	return result
}

func (u UserWhitelistCreateInput) ConvertToDTO() (result dto.UserWhitelist) {
	result.IPAddress = u.IPAddress
	result.UserID = uint64(u.UserID)

	return
}
