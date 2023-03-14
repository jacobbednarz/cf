package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultBaseURL = "https://api.cloudflare.com/client/v4"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Makes a HTTP call to the Cloudflare API",
	Example: heredoc.Doc(`
		# Make a HTTP call to /zones
		cf api "/zones"

		# Make a HTTP with authentication (using a flag)
		cf api --api-token "..." "/zones"

		# Make a HTTP with authentication (using environment variables)
		CLOUDFLARE_API_TOKEN="..." cf api "/zones"`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := http.Client{Timeout: time.Duration(1) * time.Second}
		req, err := http.NewRequest("GET", defaultBaseURL, nil)
		if err != nil {
			fmt.Printf("failed to build request: %s", err)
			return
		}

		req = buildHTTPRequestHeaders(req)

		resp, err := c.Do(req)
		if err != nil {
			fmt.Printf("failed to make request: %s", err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("failed to read body: %s", err)
			return
		}

		fmt.Println(string(body))
	},
}

func buildHTTPRequestHeaders(req *http.Request) *http.Request {
	apiKey := viper.GetString("api-key")
	email := viper.GetString("email")
	apiToken := viper.GetString("api-token")

	if apiKey != "" {
		req.Header.Add("x-auth-key", apiKey)
	}

	if email != "" {
		req.Header.Add("x-auth-email", email)
	}

	if apiToken != "" {
		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", apiToken))
	}

	return req
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
