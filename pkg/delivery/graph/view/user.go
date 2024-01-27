package view

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"boyi/pkg/model/vo"
	"strings"
)

var (
	AccountTypeToDTO = map[AccountType]types.AccountType{
		AccountTypeAdmin:           types.AccountType__Admin,
		AccountTypeSystem:          types.AccountType__System,
		AccountTypeManager:         types.AccountType__Manager,
		AccountTypeCustomerService: types.AccountType__CustomerService,
		AccountTypeMember:          types.AccountType__Member,
		AccountTypeTourist:         types.AccountType__Tourist,
	}

	AccountTypeFromDTO = map[types.AccountType]AccountType{
		types.AccountType__Admin:           AccountTypeAdmin,
		types.AccountType__System:          AccountTypeSystem,
		types.AccountType__Manager:         AccountTypeManager,
		types.AccountType__CustomerService: AccountTypeCustomerService,
		types.AccountType__Member:          AccountTypeMember,
		types.AccountType__Tourist:         AccountTypeTourist,
	}

	UserStatusToDTO = map[UserStatus]types.UserStatus{
		UserStatusUnVerified: types.UserStatus__UnVerified,
		UserStatusActived:    types.UserStatus__Actived,
		UserStatusLocked:     types.UserStatus__Locked,
		UserStatusDisabled:   types.UserStatus__Disabled,
		UserStatusDeleted:    types.UserStatus__Deleted,
	}

	UserStatusFromDTO = map[types.UserStatus]UserStatus{
		types.UserStatus__UnVerified: UserStatusUnVerified,
		types.UserStatus__Actived:    UserStatusActived,
		types.UserStatus__Locked:     UserStatusLocked,
		types.UserStatus__Disabled:   UserStatusDisabled,
		types.UserStatus__Deleted:    UserStatusDeleted,
	}
)

func (filter UserFilterInput) ConvertToOption() option.UserWhereOption {
	var result option.UserWhereOption

	if filter.User != nil {
		result.User = filter.User.ConvertToDTO()
	}

	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if len(filter.TagIDs) != 0 {
		result.TagIDs = make([]uint64, len(filter.TagIDs))
		for i := range result.TagIDs {
			result.TagIDs[i] = uint64(filter.TagIDs[i])
		}
	}

	if len(filter.RoleIDs) != 0 {
		result.RoleIDs = make([]uint64, len(filter.RoleIDs))
		for i := range result.RoleIDs {
			result.RoleIDs[i] = uint64(filter.RoleIDs[i])
		}
	}

	return result
}

func (user UserInput) ConvertToDTO() (result dto.User) {
	if user.ID != nil {
		result.ID = uint64(*user.ID)
	}

	if user.AccountType != nil {
		result.AccountType = AccountTypeToDTO[*user.AccountType]
	}
	if user.Status != nil {
		result.Status = UserStatusToDTO[*user.Status]
	}
	if user.Username != nil {
		result.Username = *user.Username
	}
	if user.AliasName != nil {
		result.AliasName = *user.AliasName
	}
	if user.Email != nil {
		result.Email = *user.Email
	}
	if user.Area != nil {
		result.Area = *user.Area
	}
	if user.Notes != nil {
		result.Notes = *user.Notes
	}

	return
}

func (u *User) FromDTO(in *dto.User, fileURI string) *User {
	if u == nil {
		u = &User{}
	}

	if in == nil {
		return u
	}

	u.ID = in.ID

	u.AccountType = AccountTypeFromDTO[in.AccountType]
	u.Status = UserStatusFromDTO[in.Status]
	u.Username = in.Username
	u.AliasName = in.AliasName
	u.Email = in.Email
	u.Area = in.Area
	u.Notes = in.Notes
	if in.AvatarKey != "" {
		u.AvatarURL = in.AvatarKey.ToURL(fileURI)
	}
	u.LastLoginIP = in.LastLoginIP
	u.LastLoginAt = in.LastLoginAt
	u.CreatedAt = in.CreatedAt
	u.UpdatedAt = in.UpdatedAt
	u.UpdateUserID = in.UpdateUserID

	u.Roles = make([]*UserRole, len(in.Roles))
	u.Whitelists = make([]*UserWhitelist, len(in.Whitelists))
	u.Tags = make([]*UserTag, len(in.Tags))

	for i := range u.Roles {
		u.Roles[i] = &UserRole{}
		u.Roles[i] = u.Roles[i].FromDTO(in.Roles[i])
	}

	for i := range u.Whitelists {
		u.Whitelists[i] = &UserWhitelist{}
		u.Whitelists[i] = u.Whitelists[i].FromDTO(in.Whitelists[i])
	}

	for i := range u.Tags {
		u.Tags[i] = &UserTag{}
		u.Tags[i] = u.Tags[i].FromDTO(in.Tags[i])
	}

	return u
}

func (cols UserUpdateInput) ConvertToOption() (result option.UserUpdateColumn) {
	if cols.AliasName != nil {
		result.AliasName = *cols.AliasName
	}

	if cols.Area != nil {
		result.Area = *cols.Area
	}
	if cols.Notes != nil {
		result.Notes = *cols.Notes
	}
	if cols.AvatarContent != nil {
		result.AvatarContent = cols.AvatarContent
	}
	if cols.AccountType != nil {
		result.AccountType = AccountTypeToDTO[*cols.AccountType]
	}

	return
}

func (filter UserLoginHistoryFilterInput) ConvertToOption() option.UserLoginHistoryWhereOption {
	var result option.UserLoginHistoryWhereOption

	if filter.UserLoginHistory != nil {
		result.UserLoginHistory = filter.UserLoginHistory.ConvertToDTO()
	}

	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	return result
}

func (user UserLoginHistoryInput) ConvertToDTO() (result dto.UserLoginHistory) {
	if user.UserID != nil {
		result.UserID = *user.UserID
	}

	return
}

func (u *UserLoginHistory) FromDTO(in dto.UserLoginHistory) *UserLoginHistory {
	u.ID = in.ID
	u.UserID = in.UserID
	u.IPAddress = in.IPAddress
	u.Country = in.Country
	u.AdministrativeArea = in.AdministrativeArea
	u.CreatedAt = in.CreatedAt

	return u
}

func (in *CreateUserReqInput) ConvertToUser() dto.User {
	if in == nil {
		return dto.User{}
	}

	return dto.User{
		Username:    in.Username,
		AliasName:   in.AliasName,
		Password:    in.Password,
		AccountType: AccountTypeToDTO[in.AccountType],
	}
}

func (cols UserUpdatePasswordInput) ConvertToOption(c claims.Claims) (result option.UserUpdateColumn) {
	if cols.Password != "" {
		result.Password = cols.Password
	}

	result.UpdateUserID = c.Id

	return
}

func (in *CreateCommonUserReqInput) ConvertToDTO() vo.RegisterReq {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)
	return vo.RegisterReq{
		Username:    in.Username,
		Password:    in.Password,
		AccountType: AccountTypeToDTO[in.AccountType],
	}
}
