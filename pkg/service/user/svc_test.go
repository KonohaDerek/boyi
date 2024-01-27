package user

import (
	"boyi/internal/mock"
	"boyi/internal/test_fixture"
	"boyi/pkg/iface"
	"boyi/pkg/repository"
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type Suite struct {
	ctx  context.Context
	svc  iface.IUserService
	repo iface.IRepository
}

var suite Suite

func TestMain(m *testing.M) {
	err := test_fixture.Initialize(
		repository.Module,
		fx.Provide(
			New,
			mock.NewStorageSvc,
			mock.NewCacheRepo,
			mock.NewQQZengIP,
			mock.NewSupportSvc,
			mock.NewRedisLocker,
		),
		fx.Populate(&suite.svc, &suite.repo),
		fx.Invoke(test_fixture.MigrationTestData),
	)

	if err != nil {
		panic(err)
	}

	ctx := log.Logger.WithContext(context.Background())
	suite.ctx = ctx
	e := m.Run()
	test_fixture.Close()
	os.Exit(e)
}
