package cache

import (
	"context"
)

func (r *repository) FlushAllCache(ctx context.Context) error {
	if err := r.redis.FlushAll(ctx).Err(); err != nil {
		return err
	}
	return nil
}
