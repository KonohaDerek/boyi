// Package zlog 負責初始化 zerolog 的格式和等級
package zlog

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// Teal ...
	Teal = Color("\033[1;36m%s\033[0m")
	// Yellow ...
	Yellow = Color("\033[35m%s\033[0m")
)

// Color ...
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// Config ...
type Config struct {
	AppName     string        `mapstructure:"app_name"`
	Environment string        `mapstructure:"environment"` // local, dev, sitable, prod
	Level       zerolog.Level `mapstructure:"level"`       // Debug: 0, info: 1 , warn: 2, error: 3, fatal: 4, panic: 5
	Local       bool
}

// NewInjection ...
func (c Config) NewInjection() *Config {
	return &c
}

// Init ...
func Init(c *Config, sentryConfig *SentryConfig) {
	zerolog.DisableSampling(true)
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.9999Z07:00"
	zerolog.ErrorStackMarshaler = zerolog.ErrorMarshalFunc
	hostname, _ := os.Hostname()
	lvl := zerolog.DebugLevel
	if c.Level != 0 {
		lvl = c.Level
	}

	var (
		z zerolog.Logger
	)

	if c.Local {
		output := zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("[ %s ]", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", Teal(i))
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatTimestamp = func(i interface{}) string {
			t := fmt.Sprintf("%s", i)
			millisecond, err := strconv.ParseInt(fmt.Sprintf("%s", i), 10, 64)
			if err == nil {
				t = time.Unix(int64(millisecond/1000), 0).Local().Format("2006/01/02 15:04:05")
			}
			return Yellow(t)
		}
		z = zerolog.New(output)
	} else {
		multiWriter := zerolog.MultiLevelWriter(new(Writer), os.Stdout)
		z = zerolog.New(multiWriter)
	}

	log.Logger = z.With().
		Fields(map[string]interface{}{
			"app": c.AppName,
			"env": c.Environment,
		}).
		Str("host", hostname).
		Timestamp().
		Caller().
		Logger().
		Level(lvl)
}

// Ctx wrap zerolog Ctx func, if ctx not setting logger, return a default prevent for panic
func Ctx(ctx context.Context) *zerolog.Logger {
	defaultLogger := log.Logger
	if ctx == nil {
		defaultLogger.Warn().Msg("zlog func Ctx() not set context.Context in right way.")
		return &defaultLogger
	}

	return log.Ctx(ctx) // if ctx is not null and not set logger yet. A disabled logger is returned.
}
