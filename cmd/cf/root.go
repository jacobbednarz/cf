package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cf command [flags]",
	Short: "cf - a CLI tool for interacting with the Cloudflare API",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cloudflare")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	rootCmd.PersistentFlags().String("base-url", "https://api.cloudflare.com/client/v4", "base URL to use for HTTP interactions")
	rootCmd.PersistentFlags().String("api-key", "", "the API key to use for authentication")
	rootCmd.PersistentFlags().String("api-token", "", "the API token to use for authentication")
	rootCmd.PersistentFlags().String("email", "", "the email to use for authentication")
	rootCmd.PersistentFlags().BoolP("verbose", "v", viper.GetBool("VERBOSE"), "enable verbose debug logs")

	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("api-token", rootCmd.PersistentFlags().Lookup("api-token"))
	viper.BindPFlag("email", rootCmd.PersistentFlags().Lookup("email"))

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute: %s", err)
		os.Exit(1)
	}
}
