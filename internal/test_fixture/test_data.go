package test_fixture

import (
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"context"
	"time"

	"boyi/pkg/infra/db"
	"boyi/pkg/infra/errors"

	"github.com/rs/zerolog/log"
)

func MigrationTestData(repo iface.IRepository) error {
	ctx := log.Logger.WithContext(context.Background())
	data := []iface.Model{}

	setUserSeed(&data)
	setUserWhitelistSeed(&data)
	// setRolesSeed(&data)
	setMerchantSeed(&data)
	setMerchantOriginSeed(&data)
	createfunc := func(in []iface.Model) error {
		for i := range in {
			if err := repo.Create(ctx, nil, in[i]); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
				log.Err(err).Msgf("fail to create")
				return err
			}
		}
		return nil
	}

	_ = createfunc(data)

	return nil
}

// 設定使用者測試資料
func setUserSeed(in *[]iface.Model) {
	*in = append(*in,
		&dto.User{
			Username:    "test data",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test data 2",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test grpc friend invitation user",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test friend invitation not allow auto friend user",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test grpc remove friend user",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test  grpc remove friend invitation user",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test cs1",
			AccountType: types.AccountType__CustomerService,
			Status:      1,
		},
		&dto.User{
			Username:    "test cs2",
			AccountType: types.AccountType__CustomerService,
			Status:      1,
		},
		&dto.User{
			Username:    "test tourist",
			AccountType: types.AccountType__Tourist,
			Status:      1,
		},
		&dto.User{
			Username:    "testLoginMangerUser",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "testLoginboyiManagerUser",
			AccountType: types.AccountType__Manager,
			Status:      1,
		},
		&dto.User{
			Username:    "test boardcast common user",
			AccountType: types.AccountType__Member,
			Status:      1,
		},
		&dto.User{
			Username:    "testUpdatePasswordboyiCustomerServiceUser",
			AccountType: types.AccountType__CustomerService,
			Status:      1,
		},
	)
}

// 設定白名單測試資料
func setUserWhitelistSeed(in *[]iface.Model) {
	*in = append(*in,
		&dto.UserWhitelist{
			ID:        1,
			UserID:    1,
			IPAddress: "127.0.0.1",
			IsBind:    common.YesNo__YES,
			CreatedAt: time.Now(),
		},
		&dto.UserWhitelist{
			ID:        2,
			UserID:    1,
			IPAddress: "192.168.0.1",
			IsBind:    common.YesNo__YES,
			CreatedAt: time.Now(),
		},
		&dto.UserWhitelist{
			ID:        3,
			UserID:    1,
			IPAddress: "192.168.0.2", // 用於 grpc 更新
			IsBind:    common.YesNo__YES,
			CreatedAt: time.Now(),
		},
		&dto.UserWhitelist{
			ID:        4,
			UserID:    1,
			IPAddress: "192.168.0.3", // 用於 grpc 刪除
			IsBind:    common.YesNo__YES,
			CreatedAt: time.Now(),
		})
}

func setMerchantSeed(in *[]iface.Model) {
	// 因為Migrate 有預設資料，因此序號要從3開始
	*in = append(*in,
		&dto.Merchant{
			ID:           3,
			Name:         "test merchant",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "merchant3.dsn",
		},
		&dto.Merchant{
			ID:           4,
			Name:         "test merchant 2",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "merchant4.dsn",
		},
		&dto.Merchant{
			ID:           5,
			Name:         "test merchant update",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "merchant5.dsn",
		},
		&dto.Merchant{
			ID:           6,
			Name:         "test merchant delete",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "merchant6.dsn",
		},
	)
}

func setMerchantOriginSeed(in *[]iface.Model) {
	*in = append(*in,
		&dto.MerchantOrigin{
			ID:           3,
			Origin:       "merchant_3.host",
			MerchantID:   3,
			MerchantName: "test merchant",
			IsEnable:     common.YesNo__YES,
		},
		&dto.MerchantOrigin{
			ID:           4,
			Origin:       "merchant_4.host",
			MerchantID:   4,
			MerchantName: "test merchant 2",
			IsEnable:     common.YesNo__YES,
		},
		&dto.MerchantOrigin{
			ID:           5,
			Origin:       "merchant_5.host",
			MerchantID:   5,
			MerchantName: "test merchant update",
			IsEnable:     common.YesNo__YES,
		},
		&dto.MerchantOrigin{
			ID:           6,
			Origin:       "merchant_6.host",
			MerchantID:   6,
			MerchantName: "test merchant delete",
			IsEnable:     common.YesNo__YES,
		},
	)
}
