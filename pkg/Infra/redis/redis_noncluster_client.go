package redis

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/cenkalti/backoff/v4"
	"github.com/redis/go-redis/v9"
	goredis "github.com/redis/go-redis/v9"
)

// NonClusterClient ...
type NonClusterClient struct {
	*goredis.Client
	Config *Config
}

func newClient(redisCfg *Config) (Redis, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if len(redisCfg.Addresses) == 0 {
		return nil, fmt.Errorf("redis config address is empty")
	}

	var client *redis.Client
	err := backoff.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr:       redisCfg.Addresses[0],
			Password:   redisCfg.Password,
			MaxRetries: redisCfg.MaxRetries,
			PoolSize:   redisCfg.PoolSizePerNode,
			DB:         redisCfg.DB,
		})
		err := client.Ping(context.Background()).Err()
		if err != nil {
			log.Error().Msgf("ping occurs error after connecting to redis: %s", err)
			return fmt.Errorf("ping occurs error after connecting to redis: %s", err)
		}
		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return &NonClusterClient{Client: client, Config: redisCfg}, nil
}

// RegisterCasbinPubSub 註冊redis pubsub功能, 用來同步casbin policy
func (c *NonClusterClient) RegisterCasbinPubSub(ctx context.Context, externalFunc func() error) error {
	pubsub := c.Subscribe(ctx, "Casbin")

	_, err := pubsub.Receive(ctx)
	if err != nil {
		return err
	}
	go func(externalFunc func() error) {
		defer pubsub.Close()
		for {
			_, err := pubsub.Receive(ctx)
			if err != nil {
				return
			}
			delay := rand.Intn(5)
			time.Sleep(time.Millisecond * 100 * time.Duration(delay))
			err = externalFunc()
			if err != nil {
				return
			}
			log.Info().Msgf("reload policy: ")
		}
	}(externalFunc)

	return nil
}

func (c *NonClusterClient) GetClient() *redis.Client {
	return c.Client
}

func (c *NonClusterClient) GetConfig() *Config {
	return c.Config
}
