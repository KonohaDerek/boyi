package view

import (
	"boyi/internal/claims"
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"time"
)

func (r *Merchant) FromDTO(merchant dto.Merchant) *Merchant {
	extra := ""
	if merchant.Extra != nil {
		extra = string(merchant.Extra)
	}
	return &Merchant{
		ID:           merchant.ID,
		Name:         merchant.Name,
		DatabaseType: string(merchant.DatabaseType),
		DatabaseDsn:  merchant.DatabaseDSN,
		IsEnabled:    YesNoFromDTO[merchant.IsEnable],
		Extra:        extra,
		CreatedAt:    merchant.CreatedAt,
		UpdatedAt:    merchant.UpdatedAt,
	}
}

func (filter *MerchantFilterInput) ConvertToOption() (result option.MerchantWhereOption) {
	if filter == nil {
		return result
	}

	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.Merchant != nil {
		result.Merchant = filter.Merchant.ConvertToDTO()
	}
	return result
}

func (in MerchantInput) ConvertToDTO() (result dto.Merchant) {
	if in.ID != nil {
		result.ID = *in.ID
	}

	if in.Name != nil {
		result.Name = *in.Name
	}

	return
}

func (in MerchantCreateInput) ConvertToDTO() (result dto.Merchant) {
	result = dto.Merchant{
		Name:         in.Name,
		DatabaseType: db.DatabaseType(in.DatabaseType),
		DatabaseDSN:  in.DatabaseDsn,
		IsEnable:     YesNoToDTO[in.IsEnabled],
		Remark:       "",
		Extra:        db.JSON(""),
		CreateUserID: 0,
	}

	if in.Extra != nil {
		result.Extra = db.JSON(*in.Extra)
	}
	if in.Remark != nil {
		result.Remark = *in.Remark
	}
	return
}

func (in MerchantUpdateInput) ConvertToOption(claims *claims.Claims) option.MerchantUpdateColumn {
	var result option.MerchantUpdateColumn
	if in.Name != nil {
		result.Name = *in.Name
	}

	if in.DatabaseType != nil {
		result.DatabaseType = db.DatabaseType(*in.DatabaseType)
	}

	if in.DatabaseDsn != nil {
		result.DatabaseDSN = *in.DatabaseDsn
	}

	if in.IsEnabled != nil {
		result.IsEnable = YesNoToDTO[*in.IsEnabled]
	}

	if in.Extra != nil {
		result.Extra = db.JSON(*in.Extra)
	}

	if in.Remark != nil {
		result.Remark = *in.Remark
	}
	result.UpdateUserID = claims.Id
	return result

}

func (in MerchantUserCreateInput) ConvertToDTO() (user dto.MerchantUser) {
	user = dto.MerchantUser{
		Username:  in.Username,
		Password:  in.Password,
		AliasName: *in.AliasName,
		IsEnable:  YesNoToDTO[in.IsEnabled],
		Extra:     db.JSON(""),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if in.AliasName != nil {
		user.AliasName = *in.AliasName
	}
	if in.Extra != nil {
		user.Extra = db.JSON(*in.Extra)
	}

	return
}

func (in *MerchantUserFilterInput) ConvertToOption() option.MerchantUserWhereOption {
	var result option.MerchantUserWhereOption
	if in == nil {
		return result
	}

	if in.BaseFilter != nil {
		result.BaseWhere = in.BaseFilter.ConvertToBaseWhere()
		result.Sorting = in.BaseFilter.ConvertToSorting()
	}

	if in.MerchantUser != nil {
		result.MerchantUser = in.MerchantUser.ConvertToDTO()
	}
	return result
}

func (in *MerchantUserInput) ConvertToDTO() (result dto.MerchantUser) {
	if in.ID != nil {
		result.ID = *in.ID
	}

	if in.Username != nil {
		result.Username = *in.Username
	}

	if in.AliasName != nil {
		result.AliasName = *in.AliasName
	}

	if in.IsEnabled != nil {
		result.IsEnable = YesNoToDTO[*in.IsEnabled]
	}
	return
}

func (r *MerchantUser) FromDTO(user dto.MerchantUser) *MerchantUser {
	extra := ""
	if user.Extra != nil {
		extra = string(user.Extra)
	}
	return &MerchantUser{
		ID:        user.ID,
		Username:  user.Username,
		AliasName: user.AliasName,
		IsEnabled: YesNoFromDTO[user.IsEnable],
		Extra:     extra,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (in MerchantUserUpdateInput) ConvertToOption(claims *claims.Claims) option.MerchantUserUpdateColumn {
	var result option.MerchantUserUpdateColumn
	if in.AliasName != nil {
		result.AliasName = *in.AliasName
	}
	result.IsEnable = YesNoToDTO[in.IsEnabled]
	if in.Extra != nil {
		result.Extra = db.JSON(*in.Extra)
	}

	result.UpdateUserID = claims.Id
	return result
}
