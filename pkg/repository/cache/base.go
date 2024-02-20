package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"boyi/pkg/infra/errors"

	"github.com/redis/go-redis/v9"
)

func (r *repository) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error) {
	ok, err := r.redis.SetNX(ctx, key, data, expireAt).Result()
	if err != nil {
		return false, errors.ConvertRedisError(err)
	}
	return ok, err
}

func (r *repository) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	if err := r.redis.SetEx(ctx, key, data, expireAt).Err(); err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (r *repository) SetEXWithJson(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return errors.Wrapf(errors.ErrInvalidInput, "not implemented json marshal, err: %+v", err)
	}
	if err := r.redis.SetEx(ctx, key, b, expireAt).Err(); err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (r *repository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, errors.ConvertRedisError(err)
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (r *repository) Get(ctx context.Context, key string) (string, error) {
	result, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", errors.ConvertRedisError(err)
	}
	return result, nil
}

func (r *repository) SetTTL(ctx context.Context, key string, expireAt time.Duration) error {
	err := r.redis.Expire(ctx, key, expireAt).Err()
	if err != nil {
		return errors.ConvertRedisError(err)
	}

	return nil
}

func (r *repository) Scan(ctx context.Context, pattern string) (keys []string, err error) {

	var (
		count  int64 = 100
		cursur uint64
	)

	var out []string
	for {
		keys, scanCursur, err := r.redis.Scan(ctx, cursur, pattern, count).Result()
		if err != nil {
			return nil, errors.ConvertRedisError(err)
		}
		cursur = scanCursur

		out = append(out, keys...)
		if cursur == 0 {
			break
		}
	}

	return out, nil
}

func (r *repository) Del(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return errors.ConvertRedisError(err)
	}

	return nil
}

func (r *repository) RPush(ctx context.Context, key string, data interface{}) (total int64, err error) {
	total, err = r.redis.RPush(ctx, key, data).Result()
	if err != nil {
		return 0, errors.ConvertRedisError(err)
	}
	return total, nil
}

func (r *repository) LLen(ctx context.Context, key string) (total int64, err error) {
	total, err = r.redis.LLen(ctx, key).Result()
	if err != nil {
		return 0, errors.ConvertRedisError(err)
	}
	return total, nil
}

func (r *repository) ZAddNX(ctx context.Context, key string, members ...redis.Z) error {
	if err := r.redis.ZAddNX(ctx, key, members...).Err(); err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}

func (r *repository) ZCard(ctx context.Context, key string) (int64, error) {
	total, err := r.redis.ZCard(ctx, key).Result()
	if err != nil {
		return 0, errors.ConvertRedisError(err)
	}

	return total, nil
}

func (r *repository) ZPopMin(ctx context.Context, key string, popCount int64) ([]redis.Z, error) {
	zMember, err := r.redis.ZPopMin(ctx, key, popCount).Result()
	if err != nil {
		return nil, errors.ConvertRedisError(err)
	}

	return zMember, nil
}

func (r *repository) ZRangeWithScore(ctx context.Context, key string, start, end int64) ([]redis.Z, error) {
	zMember, err := r.redis.ZRangeWithScores(ctx, key, start, end).Result()
	if err != nil {
		return nil, errors.ConvertRedisError(err)
	}

	return zMember, nil
}

func (r *repository) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	zValues, err := r.redis.ZRangeByScore(ctx, key, opt).Result()
	if err != nil {
		return nil, errors.ConvertRedisError(err)
	}

	return zValues, nil
}

func (r *repository) ZRem(ctx context.Context, key string, member interface{}) error {
	if err := r.redis.ZRem(ctx, key, member).Err(); err != nil {
		return errors.ConvertRedisError(err)
	}

	return nil
}

func (r *repository) Publish(ctx context.Context, key string, message interface{}) error {
	name := fmt.Sprintf("%s:%s", r.redis.GetConfig().SubscriptNamePrefix, key)
	if err := r.redis.Publish(ctx, name, message).Err(); err != nil {
		return errors.ConvertRedisError(err)
	}
	return nil
}
