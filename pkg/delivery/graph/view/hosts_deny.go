package view

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (in HostsDenyCreateInput) ConvertToDTO(claims claims.Claims) (result dto.HostsDeny) {
	// regex ip address

	result.IPAddress = in.IPAddress
	result.IsEnabled = YesNoToDTO[in.IsEnabled]
	if in.Remark != nil {
		result.Remark = *in.Remark
	}
	result.CreateUserID = claims.Id
	return
}

func (result *HostsDeny) FromDTO(in dto.HostsDeny) *HostsDeny {
	if result == nil {
		result = &HostsDeny{}
	}
	result.ID = in.ID
	result.IPAddress = in.IPAddress
	result.IsEnabled = YesNoFromDTO[in.IsEnabled]
	result.Remark = in.Remark
	result.CreateUserID = in.CreateUserID
	result.UpdateUserID = in.UpdateUserID
	result.CreatedAt = in.CreatedAt
	result.UpdatedAt = in.UpdatedAt
	return result
}

func (filter *HostsDenyFilterInput) ConvertToOption() (result option.HostsDenyWhereOption) {
	if filter == nil {
		return result
	}

	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.HostsDeny != nil {
		result.HostsDeny = filter.HostsDeny.ConvertToDTO()
	}
	return result
}

func (in HostsDenyInput) ConvertToDTO() (result dto.HostsDeny) {

	if in.ID != nil {
		result.ID = *in.ID
	}
	if in.IPAddress != nil {
		result.IPAddress = *in.IPAddress
	}
	if in.IsEnabled != nil {
		result.IsEnabled = YesNoToDTO[*in.IsEnabled]
	}

	return
}

func (cols HostsDenyUpdateInput) ConvertToOption(claims claims.Claims) (result option.HostsDenyUpdateColumn) {
	if cols.IPAddress != nil {
		result.IPAddress = *cols.IPAddress
	}

	if cols.IsEnabled != nil {
		result.IsEnabled = YesNoToDTO[*cols.IsEnabled]
	}

	if cols.Remark != nil {
		result.Remark = *cols.Remark
	}

	result.UpdateUserID = claims.Id

	return
}
