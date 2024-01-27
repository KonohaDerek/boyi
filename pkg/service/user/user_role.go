package user

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"context"

	"gorm.io/gorm"
)

// GetUserRole 取得 UserRole 的資訊
func (s *service) GetUserRole(ctx context.Context, opt *option.UserRoleWhereOption) (dto.UserRole, error) {
	var (
		user dto.UserRole
	)

	opt.LoadRole = true
	if err := s.repo.Get(ctx, nil, &user, opt); err != nil {
		return user, err
	}
	return user, nil
}

// CreateUserRole 建立UserRole
func (s *service) CreateUserRole(ctx context.Context, data *dto.UserRole) error {
	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {

		if err := s.repo.Create(ctx, tx, data); err != nil {
			return err
		}

		if err := s.ClearUserRolesClaims(ctx, []dto.UserRole{*data}); err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

// ListUserRoles 列出UserRole
func (s *service) ListUserRoles(ctx context.Context, opt *option.UserRoleWhereOption) ([]dto.UserRole, int64, error) {
	var (
		users []dto.UserRole
	)

	opt.LoadRole = true
	total, err := s.repo.List(ctx, nil, &users, opt)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// DeleteUserRole 刪除UserRole
func (s *service) DeleteUserRole(ctx context.Context, opt *option.UserRoleWhereOption) error {

	opt.Pagination.WithoutCount = true
	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		var userRoles []dto.UserRole

		_, err := s.repo.List(ctx, tx, &userRoles, opt)
		if err != nil {
			return err
		}

		if err := s.repo.Delete(ctx, tx, &dto.UserRole{}, opt); err != nil {
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

// UpdateUserRole 更新UserRole
func (s *service) UpdateUserRole(ctx context.Context, opt *option.UserRoleWhereOption, col *option.UserRoleUpdateColumn) error {
	opt.Pagination.WithoutCount = true
	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		var userRoles []dto.UserRole

		_, err := s.repo.List(ctx, tx, &userRoles, opt)
		if err != nil {
			return err
		}

		if err := s.repo.Update(ctx, tx, opt, col); err != nil {
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

	users, _, err := s.ListUsers(ctx, &option.UserWhereOption{
		IDs: userIds,
		Pagination: common.Pagination{
			WithoutCount: true,
		},
	})
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return nil
	}

	return nil
}
