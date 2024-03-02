package repository

import (
	"boyi/pkg/infra/db"
	"boyi/pkg/infra/errors"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"

	"gorm.io/gorm"
)

func (repo *repository) GetMerchantDB(ctx context.Context, merchantId uint64) (*gorm.DB, error) {
	if repo.merchantDBs == nil {
		if _, err := repo.GetALLMerchantDB(ctx); err != nil {
			return nil, err
		}
	}

	db, isExist := repo.merchantDBs[merchantId]
	if !isExist {
		return nil, errors.ErrResourceNotFound
	}

	return db, nil
}

func (repo *repository) SetMerchantDB(ctx context.Context, merchantId uint64, dsn string, databaseType db.DatabaseType) error {
	// 建立或更新 DB 資料
	conn, err := db.SetupDatabaseConnectionString(dsn, databaseType)
	if err != nil {
		return err
	}
	// 更新 repo 暫存的商戶資料庫連線
	repo.merchantDBs[merchantId] = conn
	return nil
}

func (repo *repository) DeleteMerchantDB(ctx context.Context, merchantId uint64) error {
	// 刪除DB 資料

	delete(repo.merchantDBs, merchantId)
	return nil
}

func (repo *repository) GetALLMerchantDB(ctx context.Context) (map[uint64]*gorm.DB, error) {
	if repo.merchantDBs == nil {
		repo.merchantDBs = make(map[uint64]*gorm.DB)
		var merchants []dto.Merchant
		// 從主要資料庫取用商戶資料庫連線
		if _, err := repo.List(ctx, nil, &merchants, &option.MerchantWhereOption{}); err != nil {
			return nil, err
		}
		for _, merchant := range merchants {
			if err := repo.SetMerchantDB(ctx, merchant.ID, merchant.DatabaseDSN, merchant.DatabaseType); err != nil {
				return nil, err
			}
		}

	}
	return repo.merchantDBs, nil
}
