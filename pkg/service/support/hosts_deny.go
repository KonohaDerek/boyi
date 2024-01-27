package support

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
	"time"
)

// GetHostsDenyList 取得 HostsDeny 的資訊
func (s *service) GetHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) (dto.HostsDeny, error) {
	var (
		HostsDeny dto.HostsDeny
	)
	if err := s.repo.Get(ctx, nil, &HostsDeny, opt); err != nil {
		return HostsDeny, err
	}
	return HostsDeny, nil
}

func (s *service) ListHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) ([]dto.HostsDeny, int64, error) {
	var (
		results []dto.HostsDeny
	)
	total, err := s.repo.List(ctx, nil, &results, opt)
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}
func (s *service) CreateHostsDeny(ctx context.Context, data *dto.HostsDeny) error {
	return s.repo.Create(ctx, nil, data)
}

func (s *service) UpdateHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption, col *option.HostsDenyUpdateColumn) (dto.HostsDeny, error) {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return dto.HostsDeny{}, err
	}
	return s.GetHostsDeny(ctx, opt)
}

func (s *service) DeleteHostsDeny(ctx context.Context, opt *option.HostsDenyWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.HostsDeny{}, opt)
}

func (s *service) AutoDenyHostWithRule(ctx context.Context, t time.Time, d time.Duration) error {
	// 取得高風險IP
	// result, err := s.GetHightRiskIPs(ctx, t, d)
	// if err != nil {
	// 	return err
	// }
	// for _, item := range result {
	// 	if err := s.repo.CreateIfNotExists(ctx,
	// 		nil,
	// 		&dto.HostsDeny{
	// 			IPAddress: item.IPAddress,
	// 			IsEnabled: common.YesNo__YES,
	// 			Remark:    "異常訪問，系統自動封禁",
	// 		},
	// 		&option.HostsDenyWhereOption{
	// 			HostsDeny: dto.HostsDeny{
	// 				IPAddress: item.IPAddress,
	// 			},
	// 		}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
	// 		return err
	// 	}
	// }

	return nil
}
