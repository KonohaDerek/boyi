package repository

import (
	"boyi/internal/mock"
	"boyi/internal/test_fixture"
	"boyi/pkg/iface"
	"context"
	"os"
	"testing"

	"go.uber.org/fx"
)

type Suite struct {
	ctx  context.Context
	repo iface.IRepository
}

var suite Suite

func TestMain(m *testing.M) {
	err := test_fixture.Initialize(
		Module,
		fx.Populate(&suite.repo),
		fx.Provide(
			mock.NewCacheRepo,
		),
		fx.Invoke(test_fixture.MigrationTestData),
	)

	if err != nil {
		return
	}

	suite.ctx = context.Background()
	e := m.Run()
	test_fixture.Close()
	os.Exit(e)
}
