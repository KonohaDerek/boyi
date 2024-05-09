package merchant

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"context"

	"gorm.io/gorm"
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
func (s *service) GetMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption) (dto.MerchantAcct, error) {
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
	return s.GetMerchantAcct(ctx, opt)
}

// DeleteMerchantAcct 刪除MerchantAcct
func (s *service) DeleteMerchantAcct(ctx context.Context, opt *option.MerchantAcctWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.MerchantAcct{}, opt)
}

// 商戶帳戶異動申請
func (s *service) MerchantAcctChanges(ctx context.Context, opt *option.MerchantAcctWhereOption, col *option.MerchantAcctChangeColumn) (dto.MerchantAcct, error) {
	acct, err := s.GetMerchantAcct(ctx, opt)
	if err != nil {
		return dto.MerchantAcct{}, err
	}
	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		// 更新商戶帳戶
		err := s.repo.Update(ctx, tx, opt, col)
		if err != nil {
			return err
		}

		operation := types.ConverterToMerchantAcctOperation[col.ChangeType]
		if err := s.repo.Create(ctx, tx, &dto.MerchantAcctLog{
			Currency:          acct.Currency,
			MerchantID:        acct.MerchantID,
			MerchantName:      acct.MerchantName,
			MerchantAcctID:    acct.ID,
			ChangeBalance:     col.Balance,
			BeforeBalance:     acct.Balance,
			AfterBalance:      acct.Balance.Add(col.Balance),
			ChangeBlockAmount: col.BlockAmount,
			BeforeBlockAmount: acct.BlockAmount,
			AfterBlockAmount:  acct.BlockAmount.Add(col.BlockAmount),
			ChangeProfit:      col.Profit,
			BeforeProfit:      acct.Profit,
			AfterProfit:       acct.Profit.Add(col.Profit),
			ChangeBlockProfit: col.BlockProfit,
			BeforeBlockProfit: acct.BlockProfit,
			AfterBlockProfit:  acct.BlockProfit.Add(col.BlockProfit),
			OperationType:     operation,
			Remark:            col.Remark,
		}); err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return dto.MerchantAcct{}, txErr
	}

	return s.GetMerchantAcct(ctx, opt)
}

// 商戶帳戶異動紀錄
func (s *service) ListMerchantAcctLogs(ctx context.Context, opt *option.MerchantAcctLogWhereOption) ([]dto.MerchantAcctLog, int64, error) {
	var (
		result []dto.MerchantAcctLog
	)
	total, err := s.repo.List(ctx, nil, &result, opt)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}
