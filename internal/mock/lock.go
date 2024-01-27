package mock

import (
	"context"
	"testing"

	"github.com/bsm/redislock"
	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

func NewRedisLocker(t *testing.T) *redislock.Client {
	m := gomock.NewController(t)

	redisClient := NewMockRedisClient(m)

	cmd := redis.NewCmd(context.Background())
	cmd.SetVal(int64(1))
	redisClient.EXPECT().EvalSha(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(cmd)

	boolCmd := redis.NewBoolCmd(context.Background())
	boolCmd.SetVal(true)
	redisClient.EXPECT().SetNX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(boolCmd)

	return redislock.New(redisClient)

}
