package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter UserRoleFilterInput) ConvertToOption() option.UserRoleWhereOption {
	var result option.UserRoleWhereOption
	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.UserRole != nil {
		result.UserRole = filter.UserRole.ConvertToDTO()
	}

	return result
}

func (userRole UserRoleUpdateInput) ConvertToOption() option.UserRoleUpdateColumn {
	var result option.UserRoleUpdateColumn

	if userRole.RoleID != nil {
		result.RoleID = uint64(*userRole.RoleID)
	}

	return result
}

func (userRole UserRoleInput) ConvertToDTO() (result dto.UserRole) {
	if userRole.ID != nil {
		result.ID = uint64(*userRole.ID)
	}
	if userRole.UserID != nil {
		result.UserID = uint64(*userRole.UserID)
	}
	if userRole.RoleID != nil {
		result.RoleID = uint64(*userRole.RoleID)
	}
	if userRole.IsAdmin != nil {
		result.IsAdmin = userRole.IsAdmin.ConvertToCommon()
	}

	return
}

func (u *UserRole) FromDTO(in dto.UserRole) *UserRole {
	if u == nil {
		u = &UserRole{}
	}

	u.ID = in.ID
	u.IsAdmin = YesNoFromDTO[in.IsAdmin]
	u.RoleID = in.UserID
	u.UserID = in.UserID
	u.CreatedAt = in.CreatedAt
	u.UpdatedAt = in.UpdatedAt
	u.CreateUserID = in.CreateUserID
	u.UpdateUserID = in.UpdateUserID

	u.Role = &Role{}
	u.Role = u.Role.FromDTO(in.Role)

	return u
}

func (u UserRoleCreateInput) ConvertToDTO() (result dto.UserRole) {

	result.IsAdmin = YesNoToDTO[u.IsAdmin]
	result.RoleID = uint64(u.RoleID)
	result.UserID = uint64(u.UserID)

	return
}
