package mock

import (
	"boyi/pkg/iface"
	context "context"
	"testing"
	time "time"

	"boyi/pkg/Infra/errors"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

var (
	tmpKey map[string]string
)

// 建立 CacheRepository Mock
func NewCacheRepo(t *testing.T) iface.ICacheRepository {
	tmpKey = make(map[string]string)
	m := gomock.NewController(t)
	mock := NewMockICacheRepository(m)

	mock.EXPECT().FlushAllCache(gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().Get(gomock.Any(), gomock.Any()).
		AnyTimes().
		DoAndReturn(func(_ context.Context, key string) (string, error) {
			v, ok := tmpKey[key]
			if !ok {
				return "", errors.ErrResourceNotFound
			}

			return v, nil
		})
	mock.EXPECT().Exists(gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	mock.EXPECT().SetEX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		AnyTimes().
		DoAndReturn(func(_ context.Context, key string, value interface{}, _ time.Duration) error {
			var tmp string
			switch value := value.(type) {
			case []uint8:
				b := make([]byte, len(value))
				for i, v := range value {
					b[i] = byte(v)
				}
				tmp = string(b)
			}
			tmpKey[key] = tmp

			return nil
		})
	mock.EXPECT().SetEXWithJson(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().Scan(gomock.Any(), gomock.Any()).AnyTimes().Return([]string{}, nil)
	mock.EXPECT().SetTTL(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().SetNX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(true, nil)
	mock.EXPECT().Del(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ZRem(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ZRangeWithScore(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return([]redis.Z{}, nil)
	mock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	mock.EXPECT().ZRangeByScore(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().
		AnyTimes().
		DoAndReturn(func(_ context.Context, key string, value interface{}) ([]string, error) {
			_, ok := tmpKey[key]
			if !ok {
				return nil, errors.ErrResourceNotFound
			}
			return nil, nil
		})

	mock.EXPECT().UserOnlineMap(gomock.Any()).AnyTimes().Return(nil, nil)

	mock.EXPECT().AddUserOnline(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	mock.EXPECT().ZAddNX(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	return mock
}
