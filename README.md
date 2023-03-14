# cf

The next generation Cloudflare CLI.

> **Note** This doesn't yet work but serves as a place holder for an experimental CLI for interacting with Cloudflare APIs.

```
cf - a CLI tool for interacting with the Cloudflare API

Usage:
  cf command [flags]
  cf [command]

Available Commands:
  api         Makes a HTTP call to the Cloudflare API
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of `cf`

Flags:
      --api-key string     the API key to use for authentication
      --api-token string   the API token to use for authentication
      --base-url string    base URL to use for HTTP interactions (default "https://api.cloudflare.com/client/v4")
      --email string       the email to use for authentication
  -h, --help               help for cf
  -v, --verbose            enable verbose debug logs

Use "cf [command] --help" for more information about a command.
```

The tool follows the pattern of `cf <product> <operation> <flags>`. I.e.
`cf dns get d3b07384d113edec49eaa6238ad5ff00`

## Commands

### `api`

Makes a HTTP call to an endpoint only passing the authentication headers.

Example: `cf api "/user"`
