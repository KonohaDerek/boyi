package user

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// CreateUserLoginHistory 建立UserLoginHistory
func (s *service) CreateUserLoginHistory(ctx context.Context, data *dto.UserLoginHistory) error {
	return s.repo.Create(ctx, nil, data)
}

// ListUserLoginHistories 列出UserLoginHistory
func (s *service) ListUserLoginHistories(ctx context.Context, opt *option.UserLoginHistoryWhereOption) ([]dto.UserLoginHistory, int64, error) {
	var (
		userTags []dto.UserLoginHistory
	)
	total, err := s.repo.List(ctx, nil, &userTags, opt)
	if err != nil {
		return nil, 0, err
	}
	return userTags, total, nil
}

// 更新UserLoginHistory
func (s *service) UpdateUserLoginHistory(ctx context.Context, opt *option.UserLoginHistoryWhereOption, col *option.UserLoginHistoryUpdateColumn) error {
	return s.repo.Update(ctx, nil, opt, col)
}

func (s *service) GetLastUserLoginHistories(ctx context.Context, opt *option.UserLoginHistoryWhereOption, col *dto.UserLoginHistory) (dto.UserLoginHistory, error) {
	err := s.repo.GetLast(ctx, nil, col, opt)
	if err != nil {
		return dto.UserLoginHistory{}, err
	}
	return *col, nil
}
