package redis

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/cenkalti/backoff/v4"
	"github.com/redis/go-redis/v9"
)

// ClusterClient ...
type ClusterClient struct {
	*redis.ClusterClient
	Config *Config
}

func newClusterClient(redisCfg *Config) (Redis, error) {

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	var client *redis.ClusterClient
	err := backoff.Retry(func() error {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      redisCfg.Addresses,
			Password:   redisCfg.Password,
			MaxRetries: redisCfg.MaxRetries,
			PoolSize:   redisCfg.PoolSizePerNode,
		})
		err := client.Ping(context.Background()).Err()
		if err != nil {
			return fmt.Errorf("ping occurs error after connecting to redis: %s", err)
		}
		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return &ClusterClient{
		ClusterClient: client,
		Config:        redisCfg,
	}, nil
}

// RegisterCasbinPubSub 註冊redis pubsub功能, 用來同步casbin policy
func (c *ClusterClient) RegisterCasbinPubSub(ctx context.Context, externalFunc func() error) error {
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

func (c *ClusterClient) GetClient() *redis.Client {
	panic("not support")
}

func (c *ClusterClient) GetConfig() *Config {
	return c.Config
}
