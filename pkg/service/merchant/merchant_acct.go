package merchant

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// 商戶帳戶

// ListMerchantAccts 列出MerchantAcct
func (s *service) ListMerchantAccts(ctx context.Context, opt *option.MerchantAcctWhereOption) ([]dto.MerchantAcct, int64, error) {
	var (
		accts []dto.MerchantAcct
	)
	total, err := s.repo.List(ctx, nil, &accts, opt)
	if err != nil {
		return nil, 0, err
	}
	return accts, total, nil
}

// GetMerchantAcct 取得MerchantAcct的資訊
func (s *service) GetAcct(ctx context.Context, opt *option.MerchantAcctWhereOption) (dto.MerchantAcct, error) {
	var (
		acct dto.MerchantAcct
	)
	if err := s.repo.Get(ctx, nil, &acct, opt); err != nil {
		return acct, err
	}
	return acct, nil
}

// CreateMerchantAcct 建立MerchantAcct
func (s *service) CreateMerchantAcct(ctx context.Context, data *dto.MerchantAcct) error {
	return s.repo.Create(ctx, nil, data)
}

// UpdateMerchantAcct 更新MerchantAcct
func (s *service) UpdateMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption, col *option.MerchantAcctUpdateColumn) (dto.MerchantAcct, error) {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return dto.MerchantAcct{}, err
	}
	return s.GetAcct(ctx, opt)
}
