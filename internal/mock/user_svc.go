package mock

import (
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option"
	vo "boyi/pkg/model/vo"
	context "context"
	"errors"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
)

// NewUserSvc ...
func NewUserSvc(t *testing.T) iface.IUserService {
	userMap := mockUserSeed()
	m := gomock.NewController(t)
	mock := NewMockIUserService(m)

	mock.EXPECT().GetUser(gomock.Any(), gomock.AssignableToTypeOf(&option.UserWhereOption{})).
		AnyTimes().
		DoAndReturn(func(_ context.Context, opt *option.UserWhereOption) (dto.User, error) {
			var userID uint64
			userID = opt.User.ID

			/*查看元素在集合中是否存在 */
			user, ok := userMap[userID]
			if ok {
				return user, nil
			} else {
				return dto.User{}, errors.New("GetUser : boyi user not found , userId :" + string(rune(userID)))
			}
		})

	return mock
}

func mockUser(userId uint64, cert vo.CertificationResp) dto.User {
	result := dto.User{
		ID:           userId,
		Status:       types.UserStatus__Actived,
		Username:     cert.User.Username,
		Area:         "",
		Notes:        "",
		AvatarKey:    "",
		LastLoginAt:  time.Now(),
		LastLoginIP:  "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		UpdateUserID: 0,
		Roles:        []dto.UserRole{},
		Whitelists:   []dto.UserWhitelist{},
		Tags:         []dto.UserTag{},
		Menu:         map[dto.ManagerMenuKey]struct{}{},
		AccountType:  cert.User.AccountType,
	}
	return result
}

func mockUserSeed() map[uint64]dto.User {
	var userMap map[uint64]dto.User /*建立集合 */
	if userMap == nil {
		userMap = make(map[uint64]dto.User)
	}
	return userMap
}
