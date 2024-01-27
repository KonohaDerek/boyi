package cron

import (
	"boyi/pkg/Infra/ctxutil"
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
)

// JobWrapper job wrapper
type JobWrapper func(job Job) Job

// LogWrapper 在執行的前後印出 log
func LogWrapper(key, spec string) JobWrapper {
	return func(job Job) Job {
		return func(ctx context.Context) error {
			startAt := time.Now()
			ctx = ctxutil.TraceIDWithContext(ctx)
			logger := log.Ctx(ctx).With().
				Time("start_time", startAt).
				Str("job_name", key).
				Str("spec", spec).
				Logger()
			// logger.Debug().Msgf("schedule %s now is working, and spec = %s", key, spec)
			ctx = ctxutil.ContextWithXTraceID(ctx, ctxutil.GetTraceIDFromContext(ctx))
			// 執行排程工作
			if err := job(ctx); err != nil {
				logger.Err(err).Msg("scheduler fail")
			}
			// spentTime := (time.Now().UnixNano() - startAt) / 1e6
			// logger.Debug().Msgf("schedule %s now has done, it spent %d (ms) and spec = %s", key, spentTime, spec)

			return nil
		}
	}
}

// LogTimeoutWrapper 當執行太久的時候發出通知跟 log
func LogTimeoutWrapper(key string, timeout time.Duration) JobWrapper {
	return func(job Job) Job {
		return func(ctx context.Context) error {
			// 超時未結束的話要跳通知
			Done := make(chan struct{})
			timeoutTimer := time.NewTicker(timeout)
			go func() {
				select {
				case <-timeoutTimer.C:
					log.Ctx(ctx).Warn().Msgf("scheduler %s timout!!", key)
				case <-Done:
					//
				}
			}()

			// 執行排程工作
			job(ctx)
			Done <- struct{}{}
			return nil
		}
	}
}

// SkipIfStillRunningWrapper 重新實作 cron.SkipIfStillRunning
// skips an invocation of the Job if a previous invocation is
// still running. It logs skips to the given logger at Info level.
func SkipIfStillRunningWrapper(key string) JobWrapper {
	return func(job Job) Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}

		return func(ctx context.Context) error {
			select {
			case v := <-ch:
				defer func() {
					ch <- v
				}()

				job(ctx)
			default:
				log.Warn().Msgf("scheduler %s skip!!", key)
			}
			return nil
		}
	}
}

// RecoverJobWrapper recovery job panic
//
//	印出 stack 資訊
func RecoverJobWrapper(key string) JobWrapper {
	return func(job Job) Job {

		return func(ctx context.Context) error {
			startTime := time.Now()

			defer func() {
				if err := recover(); err != nil {
					endTime := time.Now()
					trace := make([]byte, 4096)
					runtime.Stack(trace, true)
					var msg string
					for i := 2; ; i++ {
						_, file, line, ok := runtime.Caller(i)
						if !ok {
							break
						}
						msg += fmt.Sprintf("%s:%d\n", file, line)
					}

					log.Ctx(ctx).Err(err.(error)).
						Str("stack_error", string(trace)).
						Time("end_time", endTime).
						Float64("latency.sec", endTime.Sub(startTime).Seconds()).
						Msgf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥\n %s scheduler error", err, msg, key)
				}
			}()
			job(ctx)
			return nil
		}
	}
}
