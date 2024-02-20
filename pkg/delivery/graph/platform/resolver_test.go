package platform_test

import (
	"boyi/internal/graph/platform/generated"
	"boyi/internal/mock"
	"boyi/internal/test_fixture"
	"boyi/pkg/delivery/graph/platform"
	"boyi/pkg/hub"
	"boyi/pkg/repository"
	"boyi/pkg/service/audit_log"
	"boyi/pkg/service/auth"
	"boyi/pkg/service/menu"
	"boyi/pkg/service/role"
	"boyi/pkg/service/support"
	"boyi/pkg/service/tag"
	"boyi/pkg/service/user"
	"context"
	"os"
	"testing"

	"boyi/pkg/infra/zlog"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

type Suite struct {
	ctx              context.Context
	queryResolver    generated.QueryResolver
	mutationResolver generated.MutationResolver
	subResolver      generated.SubscriptionResolver
}

var suite Suite

func TestMain(m *testing.M) {
	var resolver *platform.Resolver
	if err := test_fixture.Initialize(
		repository.Module,
		hub.Module,
		fx.Options(
			auth.Module,
			menu.Module,
			user.Module,
			role.Module,
			tag.Module,
			audit_log.Module,
			support.Module,
		),
		fx.Provide(
			platform.NewResolver,
			mock.NewCacheRepo,
			mock.NewStorageSvc,
			mock.NewQQZengIP,
			mock.NewRedisLocker,
		),
		fx.Populate(&resolver),
		fx.Invoke(test_fixture.MigrationTestData),
	); err != nil {
		panic(err)
	}

	suite.mutationResolver = resolver.Mutation()
	suite.queryResolver = resolver.Query()
	suite.subResolver = resolver.Subscription()

	zlog.Init(&zlog.Config{Local: true}, nil)

	suite.ctx = log.Logger.WithContext(context.Background())
	e := m.Run()
	test_fixture.Close()
	os.Exit(e)
}
