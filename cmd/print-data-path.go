package cmd

import (
	"fmt"

	"github.com/cqroot/garden/internal/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printDataPathCmd)
}

var Msg string
var printDataPathCmd = &cobra.Command{
	Use:   "print-data-path",
	Short: "Print data path",
	Long:  "Print data path",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.Bootstrap()
		cobra.CheckErr(err)

		fmt.Println(app.Config().DataPath())
	},
}
