package merchant

import (
	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// 商戶清單
func (s *service) ListMerchants(ctx context.Context, opt *option.MerchantWhereOption) ([]dto.Merchant, int64, error) {
	var (
		merchants []dto.Merchant
	)
	total, err := s.repo.List(ctx, nil, &merchants, opt)
	if err != nil {
		return nil, 0, err
	}
	return merchants, total, nil
}

// GetMerchant 取得特定商戶的資訊
func (s *service) GetMerchant(ctx context.Context, opt *option.MerchantWhereOption) (dto.Merchant, error) {
	var (
		merchant dto.Merchant
	)

	if err := s.repo.Get(ctx, nil, &merchant, opt); err != nil {
		return merchant, err
	}
	return merchant, nil
}

// 創建商戶
func (s *service) CreateMerchant(ctx context.Context, data *dto.Merchant) error {
	return s.repo.Create(ctx, nil, data)
}

// 更新商戶
func (s *service) UpdateMerchant(ctx context.Context, opt *option.MerchantWhereOption, col *option.MerchantUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// 刪除商戶
func (s *service) DeleteMerchant(ctx context.Context, opt *option.MerchantWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.Merchant{}, opt)
}

// 商戶來源清單
func (s *service) ListMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) ([]dto.MerchantOrigin, int64, error) {
	var (
		merchantOrigins []dto.MerchantOrigin
	)
	total, err := s.repo.List(ctx, nil, &merchantOrigins, opt)
	if err != nil {
		return nil, 0, err
	}
	return merchantOrigins, total, nil
}

// 取得特定商戶來源的資訊
func (s *service) GetMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) (dto.MerchantOrigin, error) {
	var (
		merchantOrigin dto.MerchantOrigin
	)

	if err := s.repo.Get(ctx, nil, &merchantOrigin, opt); err != nil {
		return merchantOrigin, err
	}
	return merchantOrigin, nil
}

// 創建商戶來源
func (s *service) CreateMerchantOrigin(ctx context.Context, data *dto.MerchantOrigin) error {
	return s.repo.Create(ctx, nil, data)
}

// 更新商戶來源
func (s *service) UpdateMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption, col *option.MerchantOriginUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// 刪除商戶來源
func (s *service) DeleteMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.MerchantOrigin{}, opt)
}

// GetMerchantOriginFromCtx 從 context 取得商戶來源資訊
func (s *service) GetMerchantOriginFromCtx(ctx context.Context) (dto.MerchantOrigin, error) {
	merchantOrigin := dto.MerchantOrigin{}
	origin := ctxutil.GetOriginFromContext(ctx)
	if len(origin) <= 0 {
		return merchantOrigin, errors.ErrResourceNotFound
	}

	err := s.repo.Get(ctx, nil, &merchantOrigin, &option.MerchantOriginWhereOption{
		MerchantOrigin: dto.MerchantOrigin{
			Origin: origin,
		},
	})
	if err != nil {
		return merchantOrigin, errors.ErrResourceNotFound
	}
	return merchantOrigin, nil
}
