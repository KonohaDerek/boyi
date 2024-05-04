package merchant

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

func (s *service) GetUserRole(ctx context.Context, opt *option.MerchantUserRoleWhereOption) (dto.MerchantUserRole, error) {
	var (
		role dto.MerchantUserRole
	)

	opt.LoadRole = true
	if err := s.repo.Get(ctx, nil, &role, opt); err != nil {
		return role, err
	}
	return role, nil
}

func (s *service) ListUserRoles(ctx context.Context, opt *option.MerchantUserRoleWhereOption) ([]dto.MerchantUserRole, int64, error) {
	var (
		roles []dto.MerchantUserRole
	)

	opt.LoadRole = true
	total, err := s.repo.List(ctx, nil, &roles, opt)
	if err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}

func (s *service) CreateUserRole(ctx context.Context, data *dto.MerchantUserRole) error {
	return s.repo.Create(ctx, nil, data)
}

func (s *service) UpdateUserRole(ctx context.Context, opt *option.MerchantUserRoleWhereOption, col *option.MerchantUserRoleUpdateColumn) (dto.MerchantUserRole, error) {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return dto.MerchantUserRole{}, err
	}
	return s.GetUserRole(ctx, opt)
}

func (s *service) DeleteUserRole(ctx context.Context, opt *option.MerchantUserRoleWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.MerchantUserRole{}, opt)
}
