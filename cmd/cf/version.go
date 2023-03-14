package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of `cf`",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("cf (%s)", version))
	},
}
