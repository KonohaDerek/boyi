package zlog

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type SentryConfig struct {
	Dsn         string `mapstructure:"dsn"`
	Debug       bool   `mapstructure:"debug"`
	AppName     string `mapstructure:"app_name"`
	Environment string `mapstructure:"environment"`

	IgnoreErrors []string
}

func InitSentry(lc fx.Lifecycle, cfg *SentryConfig) error {
	// sentry ignoreErrors 處理
	defaultIgnoreErrors := []string{
		http.ErrServerClosed.Error(),
		net.ErrClosed.Error(),
		context.Canceled.Error(),
	}
	cfg.IgnoreErrors = append(cfg.IgnoreErrors, defaultIgnoreErrors...)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:          cfg.Dsn,
		ServerName:   cfg.AppName,
		Debug:        false,
		Environment:  cfg.Environment,
		IgnoreErrors: cfg.IgnoreErrors,
	})
	if err != nil {
		return err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			sentry.Flush(2 * time.Second)
			return nil
		},
	})

	return nil
}

var (
	levelsMapping = map[zerolog.Level]sentry.Level{
		zerolog.ErrorLevel: sentry.LevelError,
		zerolog.FatalLevel: sentry.LevelFatal,
		zerolog.PanicLevel: sentry.LevelFatal,
	}
)
