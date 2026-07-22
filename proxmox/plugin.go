package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "proxmox",

		ConnectionConfigSchema: ConfigSchema,

		TableMap: map[string]*plugin.Table{},
	}
}
