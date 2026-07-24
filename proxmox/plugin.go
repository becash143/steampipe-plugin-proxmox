package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name:                   "proxmox",
		DefaultTransform:       transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: ConfigSchema,
		TableMap: map[string]*plugin.Table{
			"proxmox_node":             tableProxmoxNode(),
			"proxmox_vm":               tableProxmoxVM(),
			"proxmox_container":        tableProxmoxContainer(),
			"proxmox_cluster_resource": tableProxmoxClusterResource(),
			"proxmox_storage":          tableProxmoxStorage(),
			"proxmox_network":          tableProxmoxNetwork(),
			"proxmox_task":             tableProxmoxTask(),
			"proxmox_user":             tableProxmoxUser(),
			"proxmox_pool":             tableProxmoxPool(),
		},
	}
}
