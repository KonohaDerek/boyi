package user

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// GetUserTag 取得 UserTag 的資訊
func (s *service) GetUserTag(ctx context.Context, opt *option.UserTagWhereOption) (dto.UserTag, error) {
	var (
		userTag dto.UserTag
	)

	if err := s.repo.Get(ctx, nil, &userTag, opt); err != nil {
		return userTag, err
	}
	return userTag, nil
}

// CreateUserTag 建立UserTag
func (s *service) CreateUserTag(ctx context.Context, data *dto.UserTag) error {
	return s.repo.Create(ctx, nil, data)
}

// ListUserTags 列出UserTag
func (s *service) ListUserTags(ctx context.Context, opt *option.UserTagWhereOption) ([]dto.UserTag, int64, error) {
	var (
		userTags []dto.UserTag
	)
	total, err := s.repo.List(ctx, nil, &userTags, opt)
	if err != nil {
		return nil, 0, err
	}
	return userTags, total, nil
}

// UpdateUserTag 更新UserTag
func (s *service) UpdateUserTag(ctx context.Context, opt *option.UserTagWhereOption, col *option.UserTagUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserTag 刪除UserTag
func (s *service) DeleteUserTag(ctx context.Context, opt *option.UserTagWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.UserTag{}, opt)
}
