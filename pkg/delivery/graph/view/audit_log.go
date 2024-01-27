package view

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
)

func (filter AuditLogFilterInput) ConvertToOption() option.AuditLogWhereOption {
	var result option.AuditLogWhereOption
	if filter.BaseFilter != nil {
		result.BaseWhere = filter.BaseFilter.ConvertToBaseWhere()
		result.Sorting = filter.BaseFilter.ConvertToSorting()
	}

	if filter.AuditLog != nil {
		result.AuditLog = filter.AuditLog.ConvertToDTO()
	}

	return result
}

func (auditLog AuditLogInput) ConvertToDTO() (result dto.AuditLog) {
	if auditLog.ID != nil {
		result.ID = uint64(*auditLog.ID)
	}
	if auditLog.Method != nil {
		result.Method = *auditLog.Method
	}
	if auditLog.UserID != nil {
		result.UserID = *auditLog.UserID
	}

	return
}

func (u *AuditLog) FromDTO(in dto.AuditLog) *AuditLog {
	u.ID = in.ID
	u.UserID = in.UserID
	u.Method = in.Method
	u.RequestInput = in.RequestInput
	u.CreatedAt = in.CreatedAt
	return u
}
