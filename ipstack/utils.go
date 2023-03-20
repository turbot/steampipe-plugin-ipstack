package ipstack

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/qioalice/ipstack"
)

func connect(_ context.Context, d *plugin.QueryData) (*ipstack.Client, error) {
	token := os.Getenv("IPSTACK_TOKEN")
	https := false
	security := false

	ipstackConfig := GetConfig(d.Connection)
	if ipstackConfig.Token != nil {
		token = *ipstackConfig.Token
	}
	if ipstackConfig.Https != nil {
		https = *ipstackConfig.Https
	}
	if ipstackConfig.Security != nil {
		security = *ipstackConfig.Security
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	return ipstack.New(
		ipstack.ParamToken(token),
		// Only do the "me" call if we want our own IP
		ipstack.ParamDisableFirstMeCall(),
		// Options enabled per config
		ipstack.ParamUseHTTPS(https),
		ipstack.ParamEnableSecurity(security),
	)
}
