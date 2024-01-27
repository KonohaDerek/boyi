package zlog

import (
	"io"
	"strconv"
	"time"
	"unsafe"

	"github.com/buger/jsonparser"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

var (
	_   = io.WriteCloser(new(Writer))
	now = time.Now
)

type Writer struct {
}

func (w *Writer) Write(data []byte) (int, error) {
	event, ok := w.parseLogEvent(data)
	if ok {
		sentry.CaptureEvent(event)
	}

	return len(data), nil
}

func (w *Writer) Close() error {
	return nil
}

func (w *Writer) parseLogEvent(data []byte) (*sentry.Event, bool) {
	lvlStr, err := jsonparser.GetUnsafeString(data, zerolog.LevelFieldName)
	if err != nil {
		return nil, false
	}

	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		return nil, false
	}

	sentryLvl, ok := levelsMapping[lvl]
	if !ok {
		return nil, false
	}

	event := sentry.Event{
		Timestamp: now(),
		Level:     sentryLvl,
		Extra:     map[string]interface{}{},
	}

	err = jsonparser.ObjectEach(data, func(key, value []byte, vt jsonparser.ValueType, offset int) error {
		switch string(key) {
		// case zerolog.LevelFieldName, zerolog.TimestampFieldName:
		case zerolog.MessageFieldName:
			event.Message = bytesToStrUnsafe(value)
		case zerolog.ErrorFieldName:
			event.Exception = append(event.Exception, sentry.Exception{
				Value:      bytesToStrUnsafe(value),
				Stacktrace: newStacktrace(),
			})
		default:
			switch vt {
			case jsonparser.Boolean:
				event.Extra[string(key)], _ = strconv.ParseBool(string(value))
			case jsonparser.String:
				event.Extra[string(key)] = string(value)
			default:
				event.Extra[string(key)] = bytesToStrUnsafe(value)
			}
		}

		return nil
	})

	if err != nil {
		return nil, false
	}

	return &event, true
}

func newStacktrace() *sentry.Stacktrace {
	const (
		currentModule = "github.com/archdx/zerolog-sentry"
		zerologModule = "github.com/rs/zerolog"
	)

	st := sentry.NewStacktrace()

	threshold := len(st.Frames) - 1
	// drop current module frames
	for ; threshold > 0 && st.Frames[threshold].Module == currentModule; threshold-- {
	}

outer:
	// try to drop zerolog module frames after logger call point
	for i := threshold; i > 0; i-- {
		if st.Frames[i].Module == zerologModule {
			for j := i - 1; j >= 0; j-- {
				if st.Frames[j].Module != zerologModule {
					threshold = j
					break outer
				}
			}

			break
		}
	}

	st.Frames = st.Frames[:threshold+1]

	return st
}

func bytesToStrUnsafe(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
