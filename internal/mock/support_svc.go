package mock

import (
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	context "context"
	"testing"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// NewSupportSvc ...
func NewSupportSvc(t *testing.T) iface.ISupportService {
	m := gomock.NewController(t)
	mock := NewMockISupportService(m)

	mock.EXPECT().GetHostsDeny(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, opt *option.HostsDenyWhereOption) (dto.HostsDeny, error) {
		var (
			data map[string]dto.HostsDeny = map[string]dto.HostsDeny{
				"1.1.1.1": {
					ID:           1,
					IPAddress:    "1.1.1.1",
					IsEnabled:    common.YesNo__YES,
					Remark:       "",
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					CreateUserID: 0,
					UpdateUserID: 0,
				},
			}
		)

		if out, ok := data[opt.HostsDeny.IPAddress]; ok && out.IsEnabled == common.YesNo__YES {
			return out, nil
		}
		return dto.HostsDeny{}, nil
	})
	return mock
}
