package mock

import (
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// NewRepository ...
func NewRepository(t *testing.T) iface.IRepository {
	m := gomock.NewController(t)
	mock := NewMockIRepository(m)

	mock.EXPECT().Transaction(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.AssignableToTypeOf(&dto.User{}), gomock.Any()).AnyTimes().Return(nil).SetArg(2, &dto.User{
		ID: 1,
	})

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.AssignableToTypeOf(&dto.UserWhitelist{}), gomock.Any()).AnyTimes().Return(nil).SetArg(2, &dto.UserWhitelist{
		ID:     1,
		UserID: 2,
	})

	mock.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.AssignableToTypeOf(&dto.User{}), gomock.AssignableToTypeOf(&option.UserWhereOption{})).AnyTimes().Return(nil).SetArg(2, dto.User{
		ID:       1,
		Username: "username_test_1",
	})

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().CreateOrUpdate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().CreateIfNotExists(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return mock
}
