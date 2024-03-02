package view

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

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
