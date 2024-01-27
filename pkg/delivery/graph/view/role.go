package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter RoleFilterInput) ConvertToOption() option.RoleWhereOption {
	var result option.RoleWhereOption
	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.Role != nil {
		result.Role = filter.Role.ConvertToDTO()
	}

	return result
}

func (role RoleInput) ConvertToDTO() (result dto.Role) {
	if role.ID != nil {
		result.ID = uint64(*role.ID)
	}
	if role.Name != nil {
		result.Name = *role.Name
	}

	if role.SupportAccountType != nil {
		result.SupportAccountType = AccountTypeToDTO[*role.SupportAccountType]
	}

	return
}

func (u *Role) FromDTO(in dto.Role) *Role {
	if u == nil {
		u = &Role{}
	}

	u.ID = in.ID
	u.Name = in.Name
	u.CreatedAt = in.CreatedAt
	u.UpdatedAt = in.UpdatedAt
	u.CreateUserID = in.CreateUserID
	u.UpdateUserID = in.UpdateUserID
	u.SupportAccountType = AccountTypeFromDTO[in.SupportAccountType]

	menus := in.Authority.GetMenus()

	u.Authority = make([]*Menu, len(menus))
	for i := range menus {
		u.Authority[i] = u.Authority[i].FromDTO(menus[i])
	}

	return u
}

func (cols RoleUpdateInput) ConvertToOption() (result option.RoleUpdateColumn) {
	if cols.Name != nil {
		result.Name = *cols.Name
	}

	if cols.SupportAccountType != nil {
		result.SupportAccountType = AccountTypeToDTO[*cols.SupportAccountType]
	}

	result.Authority = make(dto.Authority)
	if len(cols.Authority) != 0 {
		for i := range cols.Authority {
			dfsMenuInput(result.Authority, cols.Authority[i])
		}
	}

	return
}

func dfsMenuInput(result dto.Authority, input *MenuInput) {
	result[dto.ManagerMenuKey(input.Key)] = struct{}{}
	for i := range input.Next {
		dfsMenuInput(result, input.Next[i])
	}
}

func (input RoleCreateInput) ConvertToDTO() (result dto.Role) {
	result.Name = input.Name
	result.SupportAccountType = AccountTypeToDTO[input.SupportAccountType]
	result.Authority = make(dto.Authority)
	for i := range input.Authority {
		dfsMenuInput(result.Authority, input.Authority[i])
	}
	return
}
