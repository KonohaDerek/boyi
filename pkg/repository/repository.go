package repository

import (
	"boyi/pkg/iface"
	"boyi/pkg/infra/db"
	"context"
	"database/sql"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type repository struct {
	readDB  *gorm.DB
	writeDB *gorm.DB

	cacheRepo iface.ICacheRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(
		Migration,
		InitAdmin,
		InitDefaultRole,
	),
)

type Params struct {
	fx.In

	DBConns   *db.Connections
	CacheRepo iface.ICacheRepository
}

func New(p Params) (iface.IRepository, error) {
	repo := &repository{
		readDB:    p.DBConns.ReadDB,
		writeDB:   p.DBConns.WriteDB,
		cacheRepo: p.CacheRepo,
	}
	return repo, nil
}

func (repo *repository) GetDB() *gorm.DB {
	return repo.writeDB
}

func (repo *repository) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return repo.writeDB.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		return fc(tx)
	}, opts...)
}
