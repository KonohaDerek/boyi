package cron

import (
	"boyi/pkg/infra/errors"
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/robfig/cron/v3"
)

// **** by vic ****

// Job job function with context
type Job func(ctx context.Context) error

func cmd(ctx context.Context, job Job) func() {
	return func() {
		job(ctx)
	}
}

// Cron 包裝別人的 Crontab 排程工具
type Cron struct {
	*cron.Cron

	sync.Mutex
	mapTable map[string]scheduleInfo
	ctx      context.Context
	cancel   context.CancelFunc
}

type scheduleInfo struct {
	EntryID cron.EntryID
	Spec    string
}

// NewCron service DI
func NewCron() (*Cron, error) {
	logger := &cronLogger{}

	c := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(logger),
	)
	ctx, cancel := context.WithCancel(context.Background())

	return &Cron{
		Cron:     c,
		mapTable: make(map[string]scheduleInfo),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Shutdown graceful shutdown cron scheduler with 5 min timeout
func (c *Cron) Shutdown() error {
	c.cancel()
	ctx := c.Cron.Stop()
	timeout := time.After(5 * time.Minute)
	select {
	case <-ctx.Done():
		return nil
	case <-timeout:
		return errors.New("shutdown timeout")
	}
}

// AddFunc 在增加前先記錄一下
func (c *Cron) AddFunc(key string, spec string, job Job) (cron.EntryID, error) {
	job = c.wrapJob(key, spec, job)

	entryID, err := c.Cron.AddFunc(spec, cmd(c.ctx, job))
	if err != nil {
		return 0, err
	}
	c.Lock()
	c.mapTable[key] = scheduleInfo{EntryID: entryID, Spec: spec}
	c.Unlock()

	log.Ctx(c.ctx).Debug().Msgf("scheduler add job %s with %s", key, spec)
	return entryID, nil
}

func (c *Cron) wrapJob(key, spec string, job Job) Job {
	job = RecoverJobWrapper(key)(job)
	job = LogWrapper(key, spec)(job)
	job = LogTimeoutWrapper(key, 2*time.Minute)(job)
	job = SkipIfStillRunningWrapper(key)(job)
	return job
}

// EditSpec 更改指定 Crontab
func (c *Cron) EditSpec(key string, spec string) (cron.EntryID, error) {
	sInfo, ok := c.mapTable[key]
	if !ok {
		return 0, errors.NewWithMessagef(errors.ErrResourceNotFound, "schedule %s not found", key)
	}
	oriEntry := c.Entry(sInfo.EntryID)
	if oriEntry.ID == 0 {
		return 0, errors.NewWithMessagef(errors.ErrResourceNotFound, "schedule %s not found", key)
	}

	// 存在的話先刪除
	c.Remove(sInfo.EntryID)

	// 嘗試新增上去，失敗的話把舊的復原
	newEntryID, err := c.Cron.AddFunc(spec, oriEntry.Job.Run)
	if err != nil {
		newEntryID, _ = c.Cron.AddFunc(sInfo.Spec, oriEntry.Job.Run)
		// 更新內建映射表
		c.Lock()
		c.mapTable[key] = scheduleInfo{EntryID: newEntryID, Spec: sInfo.Spec}
		c.Unlock()
		return newEntryID, errors.NewWithMessagef(errors.ErrInvalidInput, "failed to add schedule %s %s, err:%+v", key, spec, err)
	}

	// 更新內建映射表
	c.Lock()
	c.mapTable[key] = scheduleInfo{EntryID: newEntryID, Spec: spec}
	c.Unlock()

	return newEntryID, nil
}

// Trigger 觸發指定 Crontab
func (c *Cron) Trigger(key string) error {
	sInfo, ok := c.mapTable[key]
	if !ok {
		return errors.NewWithMessagef(errors.ErrResourceNotFound, "schedule %s not found", key)
	}
	entry := c.Entry(sInfo.EntryID)
	if entry.ID == 0 {
		return errors.NewWithMessagef(errors.ErrResourceNotFound, "schedule %s not found", key)
	}

	entry.Job.Run()

	return nil
}
