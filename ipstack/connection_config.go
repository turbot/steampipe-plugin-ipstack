package ipstack

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type ipstackConfig struct {
	Token    *string `hcl:"token"`
	Https    *bool   `hcl:"https"`
	Security *bool   `hcl:"security"`
}

func ConfigInstance() interface{} {
	return &ipstackConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) ipstackConfig {
	if connection == nil || connection.Config == nil {
		return ipstackConfig{}
	}
	config, _ := connection.Config.(ipstackConfig)
	return config
}
