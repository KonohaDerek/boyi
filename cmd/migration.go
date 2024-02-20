package cmd

import (
	"context"
	"os"

	"boyi/pkg/infra/helper"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// SchedulerCmd 是此程式的Service入口點
var MigrateCmd = &cobra.Command{
	Run: runMigrate,
	Use: "migrate",
}

var (
	sqlPath string
)

func init() {
	MigrateCmd.PersistentFlags().StringVar(&sqlPath, "sql_path", "", "")
}

func runMigrate(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	app := fx.New(
		Module,
		RepositoryModule,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Error().Msg("start app")
		os.Exit(1)
		return
	}

	os.Exit(exitCode)
}
