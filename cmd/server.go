package cmd

import (
	"boyi/configuration"
	"boyi/internal/lock"
	"boyi/pkg/Infra/db"
	"boyi/pkg/Infra/gin"
	"boyi/pkg/Infra/mongodb"
	"boyi/pkg/Infra/qqzeng_ip"
	"boyi/pkg/Infra/redis"
	"boyi/pkg/Infra/storage"
	graphPlatform "boyi/pkg/delivery/graph/platform"
	"boyi/pkg/delivery/redis_worker"
	"boyi/pkg/delivery/restful"
	restfulPlatform "boyi/pkg/delivery/restful/backend"
	"boyi/pkg/hub"
	"boyi/pkg/model/dto"
	"boyi/pkg/repository"
	"boyi/pkg/repository/cache"
	"boyi/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"boyi/pkg/Infra/helper"
	"boyi/pkg/Infra/zlog"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// ServerCmd 是此程式的Service入口點
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "server",
}

var Module = fx.Options(
	fx.Provide(
		configuration.Init,
		gin.StartGin,
		db.InitDatabases,
		redis.InitRedisClient,
		lock.NewRedisLocker,
		mongodb.New,
		storage.New,
		qqzeng_ip.New,
	),
	fx.Invoke(
		dto.SetMenu,
		zlog.Init,
		zlog.InitSentry,
		restful.RegisterAPIRouter,
	),
)

var RepositoryModule = fx.Options(
	repository.Module,
	cache.Module,
)

var (
	migrateSQL bool
)

func init() {
	ServerCmd.PersistentFlags().BoolVar(&migrateSQL, "migrate_sql", false, "migrate with sql file")
}

func run(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	logger := log.Level(zerolog.InfoLevel)
	fxOption := []fx.Option{
		fx.Logger(&logger),
	}
	if migrateSQL {
		fxOption = append(fxOption, fx.Invoke(db.Migrate))
	}

	fxOption = append(fxOption,
		Module,
		RepositoryModule,
		service.Module,
		redis_worker.Module,
		hub.Module,
	)

	fxOption = append(fxOption, restfulPlatform.Module, graphPlatform.Module)

	app := fx.New(
		fxOption...,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Err(err).Msg("app start err")
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	log.Info().Msgf("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Err(err).Msg("app stop err")
	}

	os.Exit(exitCode)
}
