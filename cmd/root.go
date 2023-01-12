package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cqroot/todoapp/internal/app"
	"github.com/cqroot/todoapp/internal/server"
)

var (
	rootCmd = &cobra.Command{
		Use:   "todo-server",
		Short: "Run ToDo Server",
		Long:  "Run ToDo Server",
		Run:   runRootCmd,
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func runRootCmd(cmd *cobra.Command, args []string) {
	app.Bootstrap()

	err := server.Run()
	cobra.CheckErr(err)
}
