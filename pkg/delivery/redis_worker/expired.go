package redis_worker

import (
	"context"
	"fmt"
	"time"

	"github.com/bsm/redislock"
	"github.com/rs/zerolog/log"
)

var (
	LockKeyHandlerExpired = "lock_handler_expired"
)

// 監聽過期通知事件
func (h *handler) HandlerExpired(ctx context.Context, msg string) error {
	log.Ctx(ctx).Debug().Msgf("HandlerExpired: %s", msg)
	// 設定分散式鎖
	lock, err := h.redisLock.Obtain(ctx, fmt.Sprintf("%s:%s", LockKeyHandlerExpired, msg), 20*time.Second, nil)
	if err == redislock.ErrNotObtained {
		return nil
	} else if err != nil {
		return err
	}

	defer func() {
		_ = lock.Release(ctx)
	}()

	return err
}
