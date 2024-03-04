package cron

import (
	"github.com/rs/zerolog/log"
)

// cronLogger 實作 cron.Logger 的 interface
type cronLogger struct{}

// Info logs routine messages about cron's operation.
func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	format, v := l.toFormat(msg, keysAndValues...)
	log.Info().Msgf(format, v)
}

// Error logs an error condition.
func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	format, v := l.toFormat(msg, keysAndValues...)
	log.Err(err).Msgf(format, v)
}

func (l *cronLogger) toFormat(msg string, keysAndValues ...interface{}) (format string, v []interface{}) {
	format = "%+v"
	for i := 0; i < len(keysAndValues); i++ {
		format += " %+v"
	}
	v = []interface{}{msg}
	v = append(v, keysAndValues...)
	return format, v
}
