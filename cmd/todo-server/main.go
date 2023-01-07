package main

import (
	"github.com/spf13/cobra"

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

func runRootCmd(cmd *cobra.Command, args []string) {
	server.Run()
}

func main() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
