package merchant

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// 商戶帳戶清單
func (s *service) ListUsers(ctx context.Context, opt *option.MerchantUserWhereOption) ([]dto.MerchantUser, int64, error) {
	var (
		users []dto.MerchantUser
	)
	total, err := s.repo.List(ctx, nil, &users, opt)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// GetUser 取得特定商戶使用者的資訊
func (s *service) GetUser(ctx context.Context, opt *option.MerchantUserWhereOption) (dto.MerchantUser, error) {
	var (
		user dto.MerchantUser
	)

	if err := s.repo.Get(ctx, nil, &user, opt); err != nil {
		return user, err
	}
	return user, nil
}

// 創建商戶帳戶
func (s *service) CreateUser(ctx context.Context, data *dto.MerchantUser) error {
	return s.repo.Create(ctx, nil, data)
}

// 更新商戶
func (s *service) UpdateUser(ctx context.Context, opt *option.MerchantUserWhereOption, col *option.MerchantUserUpdateColumn) (dto.MerchantUser, error) {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return dto.MerchantUser{}, err
	}
	return s.GetUser(ctx, opt)
}

// 刪除商戶
func (s *service) DeleteUser(ctx context.Context, opt *option.MerchantUserWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.MerchantUser{}, opt)
}
