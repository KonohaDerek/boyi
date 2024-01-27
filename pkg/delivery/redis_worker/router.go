package redis_worker

import (
	"fmt"

	"boyi/pkg/Infra/redis"
)

func RegisterHandler(srv *handler, config *redis.Config) map[redis.SubscriptName]redis.SubscriptionFunc {
	fmt.Printf("key : __keyevent@%d__:expired \n", config.DB)
	return map[redis.SubscriptName]redis.SubscriptionFunc{
		redis.SubscriptName(fmt.Sprintf("__keyevent@%d__:expired", config.DB)): srv.HandlerExpired, // 監聽過期事件
	}
}
