package user

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// GetUserWhitelist 取得 UserWhitelist 的資訊
func (s *service) GetUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption) (dto.UserWhitelist, error) {
	var (
		user dto.UserWhitelist
	)

	if err := s.repo.Get(ctx, nil, &user, opt); err != nil {
		return user, err
	}
	return user, nil
}

// CreateUserWhitelist 建立UserWhitelist
func (s *service) CreateUserWhitelist(ctx context.Context, data *dto.UserWhitelist) error {
	return s.repo.Create(ctx, nil, data)
}

// ListUserWhitelists 列出UserWhitelist
func (s *service) ListUserWhitelists(ctx context.Context, opt *option.UserWhitelistWhereOption) ([]dto.UserWhitelist, int64, error) {
	var (
		users []dto.UserWhitelist
	)
	total, err := s.repo.List(ctx, nil, &users, opt)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// DeleteUserWhitelist 刪除UserWhitelist
func (s *service) DeleteUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.UserWhitelist{}, opt)
}

// UpdateUserWhitelist 更新UserWhitelist
func (s *service) UpdateUserWhitelist(ctx context.Context, opt *option.UserWhitelistWhereOption, col *option.UserWhitelistUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}
