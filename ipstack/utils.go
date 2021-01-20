package ipstack

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/qioalice/ipstack"
)

func connect(_ context.Context) (*ipstack.Client, error) {
	token, ok := os.LookupEnv("IPSTACK_TOKEN")
	if !ok || token == "" {
		return nil, errors.New("IPSTACK_TOKEN environment variable must be set")
	}
	httpsEnvVar, ok := os.LookupEnv("IPSTACK_HTTPS")
	httpsEnabled := false
	if ok {
		if strings.ToLower(httpsEnvVar) == "true" {
			httpsEnabled = true
		}
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
