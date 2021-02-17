package ipstack

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/plugin"

	"github.com/qioalice/ipstack"
)

func connect(_ context.Context, d *plugin.QueryData) (*ipstack.Client, error) {
	token := os.Getenv("IPSTACK_TOKEN")
	httpsEnvVar := os.Getenv("IPSTACK_HTTPS")

	ipstackConfig := GetConfig(d.Connection)
	if &ipstackConfig != nil {
		if ipstackConfig.Token != nil {
			token = *ipstackConfig.Token
		}

		if ipstackConfig.HttpsEnvVar != nil {
			httpsEnvVar = *ipstackConfig.HttpsEnvVar
		}

	}

	if token == "" {
		return nil, errors.New("IPSTACK_TOKEN environment variable must be set")
	}

	// httpsEnvVar, ok := os.LookupEnv("IPSTACK_HTTPS")
	httpsEnabled := false

	if strings.ToLower(httpsEnvVar) == "true" {
		httpsEnabled = true
	}

	return ipstack.New(
		ipstack.ParamToken(token),
		// Only do the "me" call if we want our own IP
		ipstack.ParamDisableFirstMeCall(),
		// HTTPS requires a subscription, and fails if it's used before that, so don't use it by default
		ipstack.ParamUseHTTPS(httpsEnabled),
		// Security requires a subscription, but doesn't fail, so leave it on
		ipstack.ParamEnableSecurity(true),
	)
}
