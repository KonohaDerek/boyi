package user

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	"boyi/pkg/model/vo"
	"context"
	"strings"
	"time"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"

	"boyi/pkg/model/option/common"

	"github.com/bsm/redislock"
	"gorm.io/gorm"
)

// GetUser 取得User的資訊
func (s *service) GetUser(ctx context.Context, opt *option.UserWhereOption) (dto.User, error) {
	var (
		user dto.User
	)

	if err := s.repo.Get(ctx, nil, &user, opt); err != nil {
		return user, err
	}

	return user, nil
}

// GetUser 取得User的資訊
func (s *service) GetUserByID(ctx context.Context, userID uint64) (dto.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

// GetUser 取得User的資訊
func (s *service) GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error) {
	return s.repo.GetUserIDs(ctx, opt)
}

// ListUsers 列出User
func (s *service) ListUsers(ctx context.Context, opt *option.UserWhereOption) ([]dto.User, int64, error) {
	var (
		users []dto.User
	)

	total, err := s.repo.List(ctx, nil, &users, opt)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUser 更新User
// 圖片存在時 刪除圖片
func (s *service) UpdateUser(ctx context.Context, opt *option.UserWhereOption, col *option.UserUpdateColumn) error {
	var (
		user dto.User
		err  error
	)
	if opt.User.ID != 0 {
		user, err = s.repo.GetUserByID(ctx, opt.User.ID)
	} else {
		err = s.repo.Get(ctx, nil, &user, opt)
	}
	if err != nil {
		return err
	}
	//	驗證權限
	claims, err := claims.GetClaims(ctx)
	if err != nil {
		return err
	}

	if err := user.VerifyAllowUpdate(claims); err != nil {
		return err
	}

	if col.AvatarContent != nil && col.AvatarContent.Size != 0 {
		col.AvatarKey = col.AvatarKey.GenerateFileKey(col.AvatarContent.Filename)
		if err := s.s3Svc.UploadFileByReader(ctx, col.AvatarKey.String(), col.AvatarContent.File, nil); err != nil {
			return err
		}

		if user.AvatarKey != "" {
			if err := s.s3Svc.DeleteFile(ctx, user.AvatarKey.String()); err != nil {
				return err
			}
		}
	}

	if txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.repo.Update(ctx, tx, opt, col); err != nil {
			return err
		}

		if col.AccountType != 0 && user.AccountType != col.AccountType {
			var userRole dto.UserRole
			userRole.UserID = opt.User.ID

			if col.AccountType == types.AccountType__CustomerService {
				userRole.RoleID = dto.DefaultCSRoleID
			} else if col.AccountType == types.AccountType__Manager {
				userRole.RoleID = dto.DefaultManagerRoleID
			}

			if err := s.repo.CreateIfNotExists(ctx, tx, &userRole, &option.UserRoleWhereOption{
				UserRole: userRole,
			}); err != nil {
				return err
			}
		}

		return nil
	}); txErr != nil {
		return err
	}

	return nil
}

// DeleteUser 刪除User
func (s *service) DeleteUser(ctx context.Context, opt *option.UserWhereOption) error {
	user, err := s.GetUser(ctx, opt)
	if err != nil {
		return err
	}

	//	驗證權限
	claims, err := claims.GetClaims(ctx)
	if err != nil {
		return err
	}

	if err := user.VerifyAllowDelete(claims); err != nil {
		return err
	}

	if err := s.repo.Transaction(ctx, func(db *gorm.DB) error {
		if err := s.repo.Delete(ctx, nil, &user, &option.UserWhereOption{
			User: dto.User{
				ID: user.ID,
			},
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *service) InitUser(ctx context.Context, resp vo.CertificationResp) (dto.User, error) {

	user := dto.User{
		LastLoginAt: time.Now().UTC(),
		LastLoginIP: ctxutil.GetRealIPFromContext(ctx),
	}
	user.AccountType = resp.User.AccountType
	user.Username = resp.User.Username
	user.Email = resp.User.Email
	user.AliasName = resp.User.AliasName
	user.AvatarKey = resp.User.AvatarKey
	user.Area = resp.User.Area
	user.Notes = resp.User.Notes

	whitelist := dto.UserWhitelist{
		IPAddress: ctxutil.GetRealIPFromContext(ctx),
		IsBind:    common.YesNo__NO,
	}

	history := dto.UserLoginHistory{
		IPAddress: ctxutil.GetRealIPFromContext(ctx),
		Token:     resp.Token,
	}

	if history.IPAddress != "" {
		parser := s.ipRepo.Get(history.IPAddress)
		areaTemp := strings.Split(parser, "|")
		if len(areaTemp) >= 3 {
			history.Country = areaTemp[1]
			history.AdministrativeArea = areaTemp[2]
		}
	}

	user.Whitelists = make([]dto.UserWhitelist, 0, 1)
	user.Whitelists = append(user.Whitelists, whitelist)

	txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.repo.Create(ctx, tx, &user); err != nil {
			return err
		}

		if history.IPAddress != "" {
			history.UserID = user.ID
			if err := s.repo.Create(ctx, tx, &history); err != nil {
				return err
			}
		}

		if user.AccountType != types.AccountType__Member {
			if err := s.repo.Create(ctx, tx, &dto.UserRole{
				UserID: user.ID,
				RoleID: dto.DefaultManagerRoleID,
			}); err != nil {
				return err
			}
		}

		return nil
	})
	if txErr != nil {
		return user, txErr
	}

	err := s.repo.Get(ctx, nil, &user, &option.UserWhereOption{
		User: dto.User{
			ID: user.ID,
		},
		LoadWhitelists: true,
		LoadRoles:      true,
		LoadRolesMenu:  true,
	})
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpsertUserLoginInfo(ctx context.Context, userID uint64) error {
	var (
		userUpdateCol option.UserUpdateColumn = option.UserUpdateColumn{
			LastLoginAt: time.Now().UTC(),
			LastLoginIP: ctxutil.GetRealIPFromContext(ctx),
		}
	)

	whitelist := dto.UserWhitelist{
		IPAddress: ctxutil.GetRealIPFromContext(ctx),
		UserID:    userID,
	}

	history := dto.UserLoginHistory{
		UserID:    userID,
		IPAddress: ctxutil.GetRealIPFromContext(ctx),
	}
	if history.IPAddress != "" {
		parser := s.ipRepo.Get(history.IPAddress)
		areaTemp := strings.Split(parser, "|")
		if len(areaTemp) >= 3 {
			history.Country = areaTemp[1]
			history.AdministrativeArea = areaTemp[2]
		}
	}

	if txErr := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.repo.Update(ctx, tx, &option.UserWhereOption{
			User: dto.User{
				ID: userID,
			},
		}, &userUpdateCol); err != nil {
			return err
		}

		if err := s.repo.CreateIfNotExists(ctx, tx, &whitelist, &option.UserWhitelistWhereOption{
			UserWhitelist: whitelist,
		}); err != nil {
			return err
		}

		if err := s.repo.Create(ctx, tx, &history); err != nil {
			return err
		}

		return nil
	}); txErr != nil {
		return txErr
	}

	return nil
}

func (s *service) CreateTouristUser(ctx context.Context, deviceUID string) (dto.User, error) {
	var (
		tourist dto.User
		err     error
	)

	lock, err := s.redisLock.Obtain(ctx, dto.LockCreateTouristKey, time.Second*5, &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(time.Millisecond * 25),
	})
	if err != nil {
		return tourist, errors.Wrapf(errors.ErrInternalError, "redis lock obtain err: %+v", err)
	}
	defer func() {
		_ = lock.Release(ctx)
	}()

	var (
		lastTouristUser dto.User
	)
	if err := s.repo.GetLast(ctx, nil, &lastTouristUser, &option.UserWhereOption{
		User: dto.User{
			AccountType: types.AccountType__Tourist,
		},
		Sorting: common.Sorting{
			SortField: "id",
			Type:      common.SortingOrderType__ASC,
		},
		Pagination: common.Pagination{
			Page:         1,
			PerPage:      1,
			WithoutCount: true,
		},
	}); err != nil {
		return tourist, err
	}

	tourist.AccountType = types.AccountType__Tourist
	tourist.Whitelists = []dto.UserWhitelist{
		{
			IPAddress: ctxutil.GetRealIPFromContext(ctx),
		},
	}

	if err := s.repo.Create(ctx, nil, &tourist); err != nil {
		return tourist, err
	}

	return tourist, nil
}

// 建立登入使用者
func (s *service) CreateUser(ctx context.Context, in *dto.User) error {
	//	驗證登入
	claims, err := claims.GetClaims(ctx)
	if err != nil {
		return err
	}

	if err := in.VerifyAllowCreate(claims); err != nil {
		return err
	}

	if err := s.repo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := s.repo.CreateIfNotExists(ctx, tx.Unscoped(), in, &option.UserWhereOption{
			User: dto.User{
				Username: in.Username,
			},
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
