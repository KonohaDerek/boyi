package test_fixture

import (
	"boyi/configuration"
	"boyi/pkg/model/dto"

	"boyi/pkg/Infra/db"
	"boyi/pkg/Infra/zlog"

	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Fixture struct {
	app *fx.App
}

var fixture Fixture

// Initialize 初始化 suite
func Initialize(fxOption ...fx.Option) error {
	viper.AutomaticEnv()
	if os.Getenv("CONFIG_NAME") == "" {
		_ = os.Setenv("CONFIG_NAME", "app_test")
	}

	if os.Getenv("CONFIG_PATH") == "" {
		if os.Getenv("PROJ_DIR") == "" {
			return errors.New("PROJ_DIR is required")
		}
		_ = os.Setenv("CONFIG_PATH", os.Getenv("PROJ_DIR")+"/deployment/config")
	}

	fmt.Println(os.Getenv("CONFIG_PATH"), os.Getenv("CONFIG_NAME"))

	var config = configuration.AppConfig{
		Log: &zlog.Config{
			Local: true,
		},
		Databases: &db.Config{},
		App:       &configuration.App{},
	}

	var err error
	isLocalTest := os.Getenv("LOCAL_TEST")
	if isLocalTest == "1" {
		dbCfg := &db.Database{
			Type:         db.SQLite,
			Debug:        true,
			WithColor:    false,
			Host:         "file::memory:?cache=shared",
			MaxOpenConns: 1,
			MaxIdleConns: 0,
		}
		config.Databases.Read = dbCfg
		config.Databases.Write = dbCfg
	} else {
		config, err = configuration.Init()
		if err != nil {
			return err
		}
	}
	config.App.MenuFilePath = os.Getenv("PROJ_DIR") + "/internal/test_fixture/menu.csv"
	config.App.MenuDefaultAdminFilePath = os.Getenv("PROJ_DIR") + "/internal/test_fixture/menu_default_admin.csv"
	config.App.MenuDefaultCSFilePath = os.Getenv("PROJ_DIR") + "/internal/test_fixture/menu_default_cs.csv"

	fmt.Printf("%+v", config)

	conns := &db.Connections{}
	t := &testing.T{}
	base := []fx.Option{
		fx.Supply(config),
		fx.Supply(t),
		fx.Provide(
			db.InitDatabases,
		),
		fx.Invoke(dto.SetMenu),
		fx.Supply(&conns),
	}

	conns.WriteDB = conns.ReadDB // sqlite 好像會有lock table 的問題

	// config.Databases.MigratePath = path.Join(os.Getenv("PROJ_DIR"),"deployment/database")
	// if err := db.Migrate(config.Databases); err != nil {
	// 	return err
	// }

	base = append(base, fxOption...)

	app := fx.New(
		base...,
	)

	fixture.app = app
	return app.Start(context.Background())
}

// Close 停止 container
func Close() {
	log.Info().Msgf("close app")
	isLocalTest := viper.GetString("LOCAL_TEST")
	if isLocalTest == "1" {
		os.Remove(fmt.Sprintf("%s/test/.data/sqlite.db", os.Getenv("PROJ_DIR")))
	}

}
