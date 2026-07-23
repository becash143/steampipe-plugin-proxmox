package proxmox

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type Config struct {
	Endpoint  string `cty:"endpoint"`
	APIToken  string `cty:"api_token"`
	APISecret string `cty:"api_secret"`
	Insecure  *bool  `cty:"insecure"`
}

func ConfigInstance() interface{} {
	return &Config{}
}

var ConfigSchema = &plugin.ConnectionConfigSchema{
	NewInstance: ConfigInstance,
	Schema: map[string]*schema.Attribute{
		"endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"api_token": {
			Type:     schema.TypeString,
			Required: true,
		},
		"api_secret": {
			Type:     schema.TypeString,
			Required: true,
		},
		"insecure": {
			Type: schema.TypeBool,
		},
	},
}
