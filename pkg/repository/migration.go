package repository

import (
	"boyi/configuration"
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"boyi/pkg/model/option/common"
	"context"
	"fmt"
	"time"

	"boyi/pkg/infra/db"
	"boyi/pkg/infra/errors"
	"boyi/pkg/infra/utils/hash"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func Migration(repo iface.IRepository) error {
	conn := repo.GetDB()
	conn.DisableForeignKeyConstraintWhenMigrating = true
	_conn := conn.Session(
		&gorm.Session{
			Logger: logger.Default.LogMode(logger.Warn),
		},
	)
	err := _conn.AutoMigrate(
		&dto.User{},
		&dto.Role{},
		&dto.Tag{},
		&dto.UserRole{},
		&dto.UserWhitelist{},
		&dto.UserTag{},
		&dto.UserLoginHistory{},
		&dto.AuditLog{},
		&dto.HostsDeny{},
		&dto.Merchant{},
		&dto.MerchantOrigin{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) SyncMenuTree() {
	menus := dto.GetMenu()

	repo.writeDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "menu_key"}},
		DoUpdates: clause.AssignmentColumns([]string{"menu_name", "menu_category", "super_key", "router"}),
	}).Create(&menus)

}

func InitAdmin(repo iface.IRepository, config *configuration.App) error {
	if config == nil {
		return nil
	}
	pwd, _ := hash.HashPassword([]byte("admin"))
	if err := repo.CreateIfNotExists(context.Background(),
		nil,
		&dto.User{
			AccountType:     types.AccountType__Admin,
			Username:        "admin",
			Password:        string(pwd),
			IsNeedChangePwd: common.YesNo__YES,
		},
		&option.UserWhereOption{
			User: dto.User{
				Username: "admin",
			},
		}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
		return err
	}
	return nil
}

func InitDefaultRole(repo iface.IRepository, app *configuration.App) error {
	defaultAdmin, err := dto.GetDefaultAuthority(app, types.AccountType__Manager)
	if err != nil {
		return err
	}

	defaultCS, err := dto.GetDefaultAuthority(app, types.AccountType__CustomerService)
	if err != nil {
		return nil
	}

	if err = repo.CreateIfNotExists(context.Background(),
		nil,
		&dto.Role{
			ID:                 dto.DefaultManagerRoleID,
			Name:               "预设管理员权限",
			IsEnable:           common.YesNo__YES,
			Authority:          defaultAdmin,
			SupportAccountType: types.AccountType__Admin,
		},
		&option.RoleWhereOption{
			Role: dto.Role{
				Name: "预设管理员权限",
			},
		}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
		return err
	}
	if err = repo.CreateIfNotExists(context.Background(),
		nil,
		&dto.Role{
			ID:                 dto.DefaultCSRoleID,
			Name:               "预设客服权限",
			IsEnable:           common.YesNo__YES,
			Authority:          defaultCS,
			SupportAccountType: types.AccountType__CustomerService,
		},
		&option.RoleWhereOption{
			Role: dto.Role{
				Name: "预设客服权限",
			},
		}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
		return err
	}
	return nil
}

func InitMerchant(repo iface.IRepository) error {
	seed := []dto.Merchant{
		{
			Name:         "demo_merchant_1",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "user:user@tcp(boyi-mysql:3306)/merchant_1",
			IsEnable:     common.YesNo__YES,
			Remark:       "",
			Extra:        db.JSON(""),
			CreatedAt:    time.Now(),
		},
		{
			Name:         "demo_merchant_2",
			DatabaseType: db.MySQL,
			DatabaseDSN:  "user:user@tcp(boyi-mysql:3306)/merchant_2",
			IsEnable:     common.YesNo__YES,
			Remark:       "",
			Extra:        db.JSON(""),
			CreatedAt:    time.Now(),
		},
	}
	for _, merchant := range seed {
		if err := repo.CreateIfNotExists(context.Background(),
			nil,
			&merchant,
			&option.MerchantWhereOption{
				Merchant: dto.Merchant{
					Name: merchant.Name,
				},
			}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
			return err
		}
		if err := InitMerchantOrigin(repo, merchant); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
			return err
		}
	}

	return nil
}

func InitMerchantOrigin(repo iface.IRepository, merchant dto.Merchant) error {
	if err := repo.CreateIfNotExists(context.Background(),
		nil,
		&dto.MerchantOrigin{
			Origin:       fmt.Sprintf("localhost.merchant.%d", merchant.ID),
			MerchantID:   merchant.ID,
			MerchantName: merchant.Name,
			IsEnable:     common.YesNo__YES,
			Extra:        db.JSON(""),
			Remark:       "",
			CreatedAt:    time.Now(),
		},
		&option.MerchantOriginWhereOption{
			MerchantOrigin: dto.MerchantOrigin{
				Origin: fmt.Sprintf("localhost.merchant.%d", merchant.ID),
			},
		}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
		return err
	}
	return nil
}

func checkAppOrigin(config *configuration.App) bool {
	if config.Origin.Name == "" || config.Origin.Host == "" {
		return false
	}
	return true
}

// # merchant
func MigrationMerchant(repo iface.IRepository) error {
	conns, err := repo.GetALLMerchantDB(context.Background())
	if err != nil {
		return err
	}
	for _, conn := range conns {
		conn.DisableForeignKeyConstraintWhenMigrating = true
		_conn := conn.Session(
			&gorm.Session{
				Logger: logger.Default.LogMode(logger.Warn),
			},
		)
		// Migrate 商戶資料庫
		err := _conn.AutoMigrate(
			&dto.MerchantAccount{},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitDefaultMerchantAccount(repo iface.IRepository) error {
	ctx := context.Background()
	conns, err := repo.GetALLMerchantDB(ctx)
	if err != nil {
		return err
	}
	for key, conn := range conns {
		conn = conn.WithContext(ctx)
		if err = repo.CreateIfNotExists(ctx,
			conn,
			&dto.MerchantAccount{
				Username:  "admin",
				Password:  db.Crypto("admin"),
				AliasName: fmt.Sprintf("admin_%d", key),
				IsEnable:  common.YesNo__YES,
				Extra:     db.JSON(""),
				CreatedAt: time.Now(),
			},
			&option.MerchantAccountWhereOption{
				MerchantAccount: dto.MerchantAccount{
					Username: "admin",
				},
			}); err != nil && !errors.Is(err, errors.ErrResourceAlreadyExists) {
			return err
		}
	}
	return nil
}
