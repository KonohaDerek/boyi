package role

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"context"

	"boyi/pkg/infra/errors"

	"gorm.io/gorm"
)

// GetRole 取得Role的資訊
func (s *service) GetRole(ctx context.Context, opt *option.RoleWhereOption) (dto.Role, error) {
	var (
		role dto.Role
	)
	if err := s.repo.Get(ctx, nil, &role, opt); err != nil {
		return role, err
	}
	return role, nil
}

// CreateRole 建立Role
func (s *service) CreateRole(ctx context.Context, data *dto.Role) error {
	if data.SupportAccountType == 0 {
		return errors.NewWithMessagef(errors.ErrInvalidInput, "role support account type can't be 0")
	}
	return s.repo.Create(ctx, nil, data)
}

// ListRoles 列出Role
func (s *service) ListRoles(ctx context.Context, opt *option.RoleWhereOption) ([]dto.Role, int64, error) {
	var (
		roles []dto.Role
	)
	total, err := s.repo.List(ctx, nil, &roles, opt)
	if err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}

// UpdateRole 更新Role
func (s *service) UpdateRole(ctx context.Context, opt *option.RoleWhereOption, col *option.RoleUpdateColumn) error {
	var (
		userRoles []dto.UserRole
		roles     []dto.Role
	)

	count, err := s.repo.List(ctx, nil, &roles, opt)
	if err != nil {
		return err
	}

	tmp := make([]uint64, count)
	for i := range tmp {
		tmp[i] = roles[i].ID
	}

	_, err = s.repo.List(ctx, nil, &userRoles, &option.UserRoleWhereOption{
		RoleIDs: tmp,
		Pagination: common.Pagination{
			WithoutCount: true,
		},
	})
	if err != nil {
		return err
	}

	if col.IsEnable != 0 && len(userRoles) > 0 {
		return errors.NewWithMessagef(errors.ErrResultBeenBound, "The role has been bound to the user and must be unbound first")
	}

	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		err := s.repo.Update(ctx, tx, opt, col)
		if err != nil {
			return err
		}

		if err := s.ClearUserRolesClaims(ctx, userRoles); err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

// DeleteRole 刪除Role
func (s *service) DeleteRole(ctx context.Context, opt *option.RoleWhereOption) error {
	var (
		userRoles []dto.UserRole
		roles     []dto.Role
	)

	opt.Pagination.WithoutCount = true
	_, err := s.repo.List(ctx, nil, &roles, opt)
	if err != nil {
		return err
	}

	tmp := make([]uint64, len(roles))
	for i := range tmp {
		tmp[i] = roles[i].ID
	}

	_, err = s.repo.List(ctx, nil, &userRoles, &option.UserRoleWhereOption{
		RoleIDs: tmp,
		Pagination: common.Pagination{
			WithoutCount: true,
		},
	})
	if err != nil {
		return err
	}

	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.repo.Delete(ctx, tx, &dto.Role{}, opt); err != nil {
			return err
		}

		if err := s.ClearUserRolesClaims(ctx, userRoles); err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

func (s *service) ClearUserRolesClaims(ctx context.Context, userRoles []dto.UserRole) error {
	var (
		users []dto.User
	)

	if len(userRoles) == 0 {
		return nil
	}

	userIds := make([]uint64, len(userRoles))
	for i := range userRoles {
		userIds[i] = userRoles[i].UserID
	}

	if len(userIds) == 0 {
		return nil
	}

	_, err := s.repo.List(ctx, nil, &users, &option.UserWhereOption{
		IDs: userIds,
	})
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return nil
	}

	return nil
}
