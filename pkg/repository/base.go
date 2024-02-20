package repository

import (
	"boyi/pkg/iface"
	"context"

	"boyi/pkg/infra/errors"

	"gorm.io/gorm"
)

// Get 取得的資訊
func (repo *repository) Get(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.readDB.WithContext(ctx)
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
func (repo *repository) GetLast(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.readDB.WithContext(ctx)
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
func (repo *repository) Create(ctx context.Context, tx *gorm.DB, data iface.Model, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Create(data).Error
	return errors.ConvertMySQLError(err)
}

// List 列出
func (repo *repository) List(ctx context.Context, tx *gorm.DB, data interface{}, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	if tx == nil {
		tx = repo.readDB.WithContext(ctx)
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
func (repo *repository) Update(ctx context.Context, tx *gorm.DB, opt iface.WhereOption, col iface.UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
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
func (repo *repository) Delete(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Scopes(opt.Where).Delete(model).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}
	return nil
}

// CreateIfNotExists 存在就不 Create
func (repo *repository) CreateIfNotExists(ctx context.Context, tx *gorm.DB, data iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
	}
	err := tx.Table(data.TableName()).Scopes(opt.Where).FirstOrCreate(data).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (repo *repository) CreateOrUpdate(ctx context.Context, tx *gorm.DB, data iface.Model, opt iface.WhereOption, updateCol iface.UpdateColumns) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
	}
	err := tx.Table(data.TableName()).Scopes(opt.Where).Assign(updateCol.Columns()).FirstOrCreate(data).Error
	if err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (repo *repository) BatchInsert(ctx context.Context, tx *gorm.DB, data interface{}) error {
	if tx == nil {
		tx = repo.writeDB.WithContext(ctx)
	}

	if err := tx.Create(data).Error; err != nil {
		return errors.ConvertMySQLError(err)
	}

	return nil
}

func (repo *repository) Count(ctx context.Context, tx *gorm.DB, opt iface.WhereOption) (int64, error) {
	if tx == nil {
		tx = repo.readDB.WithContext(ctx)
	}

	var total int64

	db := tx.Table(opt.TableName()).Scopes(opt.Where)
	err := db.Count(&total).Error
	if err != nil {
		return total, errors.ConvertMySQLError(err)
	}

	return total, nil
}
