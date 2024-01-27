package main

import (
	"boyi/cmd"
	"fmt"
	"os"
	"runtime"

	_ "boyi/pkg/infra/graph/value"

	_ "boyi/docs"

	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/mailru/easyjson/gen"
	"github.com/spf13/cobra"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var rootCmd = &cobra.Command{Use: "server scheduler migrate"}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	rootCmd.AddCommand(cmd.ServerCmd, cmd.MigrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
