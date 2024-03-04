package db

import (
	"database/sql"
	"fmt"
	"os"

	"boyi/pkg/infra/errors"

	"github.com/pressly/goose/v3"
)

func Migrate(cfg *Config) error {
	connStr, err := GetConnectionStr(cfg.Write)
	if err != nil {
		return err
	}

	db, err := sql.Open(string(cfg.Write.Type), connStr)
	if err != nil {
		return errors.Wrap(err, "fail to open sql")
	}
	if cfg.MigratePath == "" {
		dir, _ := os.Getwd()
		cfg.MigratePath = dir + "/deployment/database"
	}
	goose.SetDialect(string(cfg.Write.Type))

	fmt.Println(cfg.MigratePath)
	if err := goose.Up(db, cfg.MigratePath, goose.WithAllowMissing()); err != nil {
		return errors.Wrap(err, "fail to migrate up")
	}
	return nil
}
