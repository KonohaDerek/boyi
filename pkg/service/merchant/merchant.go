package merchant

import (
	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

func (s *service) ListMerchants(ctx context.Context, opt *option.MerchantWhereOption) ([]dto.Merchant, int64, error) {
	return []dto.Merchant{}, 0, nil
}
func (s *service) CreateMerchant(ctx context.Context, data *dto.Merchant) error {
	return nil
}
func (s *service) UpdateMerchant(ctx context.Context, opt *option.MerchantWhereOption, col *option.MerchantUpdateColumn) error {
	return nil
}
func (s *service) DeleteMerchant(ctx context.Context, opt *option.MerchantWhereOption) error {
	return nil
}

func (s *service) ListMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) ([]dto.MerchantOrigin, int64, error) {
	return []dto.MerchantOrigin{}, 0, nil
}
func (s *service) CreateMerchantOrigin(ctx context.Context, data *dto.MerchantOrigin) error {
	return nil
}
func (s *service) UpdateMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption, col *option.MerchantOriginUpdateColumn) error {
	return nil
}
func (s *service) DeleteMerchantOrigin(ctx context.Context, opt *option.MerchantOriginWhereOption) error {
	return nil
}

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
