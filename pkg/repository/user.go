package repository

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
	"fmt"
	"time"

	"boyi/pkg/Infra/errors"
)

func (repo *repository) GetUserByID(ctx context.Context, userID uint64) (dto.User, error) {
	var out dto.User
	key := fmt.Sprintf(dto.UserCacheKey, userID)

	data, err := repo.cacheRepo.Get(ctx, key)
	if err != nil && !errors.Is(err, errors.ErrResourceNotFound) {
		return dto.User{}, err
	} else if err == nil {
		if err := out.Unmarshal(data); err != nil {
			return dto.User{}, err
		}
		return out, nil
	}

	if err := repo.Get(ctx, nil, &out, &option.UserWhereOption{
		User: dto.User{
			ID: userID,
		},
	}); err != nil {
		return dto.User{}, err
	}

	if err := repo.cacheRepo.SetEX(ctx, key, out.Marshal(), time.Minute*30); err != nil {
		return dto.User{}, err
	}

	return out, nil
}

func (repo *repository) GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error) {
	var result []uint64
	tx := repo.readDB.WithContext(ctx)
	err := tx.
		Table(dto.User{}.TableName()).
		Select("id").Scopes(opt.Where).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
