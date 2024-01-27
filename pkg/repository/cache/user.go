package cache

import (
	"boyi/pkg/model/dto"
	"context"
	"strings"

	"boyi/pkg/Infra/errors"

	"github.com/rs/zerolog/log"
)

const (
	ListUserOnlineKey = "user:online"
)

func (repo *repository) UserOnlineMap(ctx context.Context) (map[string]bool, error) {
	list, err := repo.redis.SMembers(ctx, ListUserOnlineKey).Result()
	if err != nil {
		return nil, errors.ConvertRedisError(err)
	}
	if len(list) == 0 {
		old, err := repo.UserOnlineMapWithKey(ctx)
		if err != nil {
			return nil, err
		}
		for key := range old {
			list = append(list, key)
		}
		_, err = repo.redis.SAdd(ctx, ListUserOnlineKey, list).Result()
		if err != nil {
			log.Ctx(ctx).Err(err)
		}
	}
	maps := make(map[string]bool, 0)
	for _, item := range list {
		maps[item] = true
	}
	return maps, nil
}

func (repo *repository) AddUserOnline(ctx context.Context, user *dto.User) error {
	_, err := repo.redis.SAdd(ctx, ListUserOnlineKey, user.GenerateOnlineKey()).Result()
	if err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (repo *repository) RemoveUserOnline(ctx context.Context, user *dto.User) error {
	_, err := repo.redis.SRem(ctx, ListUserOnlineKey, user.GenerateOnlineKey()).Result()
	if err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (repo *repository) UserOnlineMapWithKey(ctx context.Context) (map[string]bool, error) {
	list, err := repo.redis.Keys(ctx, strings.Replace(dto.UserOnlineKey, "%d", "*", -1)).Result()
	if err != nil {
		return nil, errors.ConvertRedisError(err)
	}
	maps := make(map[string]bool, 0)
	for _, item := range list {
		maps[item] = true
	}
	return maps, nil
}
