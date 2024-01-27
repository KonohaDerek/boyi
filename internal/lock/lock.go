package lock

import (
	"boyi/pkg/Infra/redis"

	"github.com/bsm/redislock"
)

func NewRedisLocker(redis redis.Redis) *redislock.Client {
	return redislock.New(redis)
}
