package ipstack

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type ipstackConfig struct {
	Token    *string `cty:"token"`
	Https    *bool   `cty:"https"`
	Security *bool   `cty:"security"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"https": {
		Type: schema.TypeBool,
	},
	"security": {
		Type: schema.TypeBool,
	},
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
