package errors

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogError(ctx context.Context, err error, msg string, v ...interface{}) {
	logger := log.Ctx(ctx).With().Logger()

	statusCode := GetStatusWithErrors(err)

	var lEvent *zerolog.Event
	switch {
	case statusCode >= 500:
		lEvent = logger.Error()
	case statusCode >= 400:
		lEvent = logger.Warn()
	default:
		lEvent = logger.Error()
	}

	stackStr := fmt.Sprintf("%+v", err)
	stack := strings.Split(stackStr, "\n")

	l := 10
	if len(stack) < l {
		l = len(stack) - 1
	}
	lEvent = lEvent.Err(err).Strs("stack", stack[:l])
	if len(v) != 0 {
		lEvent.Msgf(msg, v)
	} else {
		lEvent.Msg(msg)
	}
}
