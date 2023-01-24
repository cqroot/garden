package cmd

import (
	"fmt"

	"github.com/cqroot/garden/internal/config"
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
		config, err := config.New()
		cobra.CheckErr(err)

		fmt.Println(config.DataPath())
	},
}
