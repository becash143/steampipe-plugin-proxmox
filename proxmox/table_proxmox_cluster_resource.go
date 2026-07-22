package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxClusterResource() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_cluster_resource",
		Description: "Unified view of all Proxmox cluster resources (nodes, VMs, containers, storage)",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxClusterResources,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Resource ID.", Transform: transform.FromField("ID")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Resource type (node, qemu, lxc, storage, pool).", Transform: transform.FromField("Type")},
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node name.", Transform: transform.FromField("Node")},
			{Name: "vmid", Type: proto.ColumnType_INT, Description: "VM/CT ID, if applicable.", Transform: transform.FromField("VMID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Resource name.", Transform: transform.FromField("Name")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Resource status.", Transform: transform.FromField("Status")},
			{Name: "cpu", Type: proto.ColumnType_DOUBLE, Description: "CPU usage ratio.", Transform: transform.FromField("CPU")},
			{Name: "maxcpu", Type: proto.ColumnType_INT, Description: "Max CPUs.", Transform: transform.FromField("MaxCPU")},
			{Name: "mem", Type: proto.ColumnType_INT, Description: "Current memory usage.", Transform: transform.FromField("Mem")},
			{Name: "maxmem", Type: proto.ColumnType_INT, Description: "Max memory.", Transform: transform.FromField("MaxMem")},
			{Name: "disk", Type: proto.ColumnType_INT, Description: "Current disk usage.", Transform: transform.FromField("Disk")},
			{Name: "maxdisk", Type: proto.ColumnType_INT, Description: "Max disk.", Transform: transform.FromField("MaxDisk")},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "Uptime in seconds.", Transform: transform.FromField("Uptime")},
			{Name: "pool", Type: proto.ColumnType_STRING, Description: "Resource pool.", Transform: transform.FromField("Pool")},
		},
	}
}

func listProxmoxClusterResources(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	resources, err := client.ListClusterResources()
	if err != nil {
		return nil, err
	}
	for _, r := range resources {
		d.StreamListItem(ctx, r)
	}
	return nil, nil
}
