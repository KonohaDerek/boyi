package mongodb

import (
	"context"
	"time"

	"boyi/pkg/Infra/errors"

	"github.com/rs/zerolog/log"

	"github.com/cenkalti/backoff/v4"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

type Config struct {
	URI                    string `mapstructure:"uri"`
	Database               string `mapstructure:"database"`
	ConnectTimeoutSec      uint64 `mapstructure:"connect_timeout_sec"`
	MaxConnIdleTimeSec     uint64 `mapstructure:"max_conn_idle_time_sec"`
	MinPoolSize            uint64 `mapstructure:"min_pool_size"`
	MaxPoolSize            uint64 `mapstructure:"max_pool_size"`
	ServerSelectionTimeout uint64 `mapstructure:"server_selection_timeout_sec"`
	HeartbeatIntervalSec   uint64 `mapstructure:"heartbeat_interval_sec"`
	Debug                  bool   `mapstructure:"debug"`
}

func New(cfg *Config, lc fx.Lifecycle) (*mongo.Database, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	var (
		ctx        context.Context = context.Background()
		connclient *mongo.Client
		db         *mongo.Database
	)

	if err := backoff.Retry(func() error {
		var err error
		clientOptions := newClientOption(cfg)

		connclient, err = mongo.NewClient(clientOptions)
		if err != nil {
			log.Err(err).Msg("fail to conn mongo")
			return errors.Wrap(errors.ErrInternalError, "fail to conn to mongo")
		}

		if err := connclient.Connect(ctx); err != nil {
			log.Err(err).Msg("fail to connect mongo")
			return err
		}

		if err := connclient.Ping(ctx, readpref.Primary()); err != nil {
			log.Err(err).Msg("fail to ping mongo")
			return errors.Wrap(errors.ErrInternalError, "fail to ping to mongo")
		}

		log.Info().Msgf("ping mongo success")

		db = connclient.Database(cfg.Database)

		return nil
	}, bo); err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			return connclient.Disconnect(ctx)
		},
	})

	return db, nil
}

func newClientOption(cfg *Config) *options.ClientOptions {
	// 設定客戶端連線配置
	clientOptions := options.Client().ApplyURI(cfg.URI)

	clientOptions.SetReadPreference(readpref.Primary())

	if cfg.MaxConnIdleTimeSec > 0 {
		clientOptions.SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTimeSec) * time.Second)
	}

	if cfg.ConnectTimeoutSec > 0 {
		clientOptions.SetConnectTimeout(time.Duration(cfg.ConnectTimeoutSec) * time.Second)
	}

	if cfg.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(cfg.MinPoolSize)
	}

	if cfg.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(cfg.MaxPoolSize)
	}

	if cfg.ServerSelectionTimeout > 0 {
		clientOptions.SetServerSelectionTimeout(time.Duration(cfg.ServerSelectionTimeout) * time.Second)
	}

	if cfg.HeartbeatIntervalSec != 0 {
		clientOptions.SetHeartbeatInterval(time.Duration(cfg.HeartbeatIntervalSec) * time.Second)
	} else {
		clientOptions.SetHeartbeatInterval(240 * time.Second)
	}

	if cfg.Debug {
		cmdMonitor := &event.CommandMonitor{
			Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
				log.Ctx(ctx).Debug().Msgf("%+v", evt)
			},
		}
		clientOptions.SetMonitor(cmdMonitor)
	}

	return clientOptions
}
