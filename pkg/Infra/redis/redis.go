package redis

import (
	"context"

	"boyi/pkg/Infra/ctxutil"
	"boyi/pkg/Infra/helper"

	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

// Config ...
type Config struct {
	ClusterMode         bool     `yaml:"cluster_mode" mapstructure:"cluster_mode"`
	Addresses           []string `yaml:"addresses" mapstructure:"addresses"`
	Password            string   `yaml:"password" mapstructure:"password"`
	MaxRetries          int      `yaml:"max_retries" mapstructure:"max_retries"`
	PoolSizePerNode     int      `yaml:"pool_size_per_node" mapstructure:"pool_size_per_node"`
	DB                  int      `yaml:"db" mapstructure:"db"`
	SubscriptNamePrefix string   `yaml:"subscript_name_prefix" mapstructure:"subscript_name_prefix"`
}

// NewInjection ...
func (c Config) NewInjection() *Config {
	return &c
}

// Redis 提供操作 redis 的介面
type Redis interface {
	redis.Cmdable
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	GetClient() *redis.Client
	GetConfig() *Config
}

// InitRedisClient init redis client
func InitRedisClient(redisCfg *Config) (Redis, error) {
	if redisCfg.ClusterMode {
		return newClusterClient(redisCfg)
	}
	return newClient(redisCfg)
}

type SubscriptName string

type SubscriptionFunc func(ctx context.Context, msg string) error

// InitSubscriptHandler ...
func InitSubscriptHandler(r Redis, redisCfg *Config, handlerMap map[SubscriptName]SubscriptionFunc) {
	logger := log.With().Str("handler", "redis_worker").Logger()
	ctx := logger.WithContext(context.Background())
	channels := make([]string, 0, len(handlerMap))
	for key := range handlerMap {
		channels = append(channels, string(key))
	}
	sub := r.Subscribe(ctx, channels...)
	go func() {
		defer helper.Recover(ctx)
		ch := sub.Channel()
		for {
			msg, ok := <-ch
			if !ok {
				break
			}

			ctx := ctxutil.TraceIDWithContext(context.Background())
			_logger := logger.With().
				Str("trace_id", ctxutil.GetTraceIDFromContext(ctx)).
				Str("channel", msg.Channel).
				Logger()
			ctx = _logger.WithContext(ctx)
			if f, ok := handlerMap[SubscriptName(msg.Channel)]; ok {
				if err := f(ctx, msg.Payload); err != nil {
					_logger.Error().Err(err).Msgf("redis worker channel error")
				}
			} else {
				_logger.Error().Msgf("Not handler channel, %s", msg.Channel)
			}
		}
		logger.Warn().Msgf("redis sub channel is close")
	}()
}
