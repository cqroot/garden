package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cqroot/garden/internal/app"
)

var (
	rootCmd = &cobra.Command{
		Use:   "garden",
		Short: "A self-hosted todo application",
		Long:  "A self-hosted todo application",
		Run:   runRootCmd,
	}
)

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func runRootCmd(cmd *cobra.Command, args []string) {
	cobra.CheckErr(app.Run())
}
