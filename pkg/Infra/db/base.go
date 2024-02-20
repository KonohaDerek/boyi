package db

import (
	"context"
	"database/sql"

	"boyi/pkg/infra/errors"

	"gorm.io/gorm"
)

type WhereOption interface {
	Model
	Where(db *gorm.DB) *gorm.DB
	Page(db *gorm.DB) *gorm.DB
	Sort(db *gorm.DB) *gorm.DB
	Preload(db *gorm.DB) *gorm.DB
	IsEmptyWhereOpt() bool
	WithoutCount() bool
}

type Model interface {
	TableName() string
}

type UpdateColumns interface {
	Columns() interface{}
}

type IBaseRepository interface {
	GetDB() *gorm.DB
	Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	Get(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, data Model, scopes ...func(*gorm.DB) *gorm.DB) error
	List(ctx context.Context, tx *gorm.DB, data interface{}, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
	Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error
	Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error

	BatchInsert(ctx context.Context, tx *gorm.DB, data interface{}) error

	CreateIfNotExists(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	CreateOrUpdate(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, updateCol UpdateColumns) error
}

// Get 取得的資訊
func (conn *Connections) Get(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.ReadDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	tx = opt.Preload(tx)
	err := tx.Table(model.TableName()).Scopes(opt.Where).First(model).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}
	return nil
}

// Get 取得的資訊
func (conn *Connections) GetLast(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.ReadDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	tx = opt.Preload(tx)
	err := tx.Table(model.TableName()).Scopes(opt.Where).Last(model).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}
	return nil
}

// Create 建立
func (conn *Connections) Create(ctx context.Context, tx *gorm.DB, data Model, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertMySQLError(err)
}

// List 列出
func (conn *Connections) List(ctx context.Context, tx *gorm.DB, data interface{}, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	if tx == nil {
		tx = conn.ReadDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	var total int64

	db := tx.Table(opt.TableName()).Model(data).Scopes(opt.Where)
	if !opt.WithoutCount() {
		err := db.Count(&total).Error
		if err != nil {
			return total, errors.ConvertMySQLError(err)
		}
	}

	db = opt.Preload(db)
	err := db.Scopes(opt.Page, opt.Sort).Find(data).Error
	if err != nil {
		return total, errors.ConvertMySQLError(err)
	}
	return total, nil
}

// Update 更新
func (conn *Connections) Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	if opt.IsEmptyWhereOpt() {
		return errors.Wrap(errors.ErrInternalError, "database: Update err: where condition can't empty")
	}
	err := tx.Table(opt.TableName()).Scopes(opt.Where).Updates(col.Columns()).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

// Delete 刪除
func (conn *Connections) Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Scopes(opt.Where).Delete(model).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}
	return nil
}

// CreateIfNotExists 存在就不 Create
func (conn *Connections) CreateIfNotExists(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}
	err := tx.Table(data.TableName()).Scopes(opt.Where).FirstOrCreate(data).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (conn *Connections) CreateOrUpdate(ctx context.Context, tx *gorm.DB, data Model, opt WhereOption, updateCol UpdateColumns) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}
	err := tx.Table(data.TableName()).Scopes(opt.Where).Assign(updateCol.Columns()).FirstOrCreate(data).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (conn *Connections) BatchInsert(ctx context.Context, tx *gorm.DB, data interface{}) error {
	if tx == nil {
		tx = conn.WriteDB.WithContext(ctx)
	}

	if err := tx.Create(data).Error; err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (conn *Connections) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return conn.WriteDB.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		return fc(tx)
	}, opts...)
}
