package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxNode() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_node",
		Description: "Proxmox cluster nodes",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxNodes,
		},
		Columns: []*plugin.Column{
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Proxmox node name.", Transform: transform.FromField("Node")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Node status.", Transform: transform.FromField("Status")},
			{Name: "cpu", Type: proto.ColumnType_DOUBLE, Description: "Current CPU usage ratio.", Transform: transform.FromField("CPU")},
			{Name: "maxcpu", Type: proto.ColumnType_INT, Description: "Number of CPUs.", Transform: transform.FromField("MaxCPU")},
			{Name: "mem", Type: proto.ColumnType_INT, Description: "Current memory usage.", Transform: transform.FromField("Memory")},
			{Name: "maxmem", Type: proto.ColumnType_INT, Description: "Maximum memory.", Transform: transform.FromField("MaxMemory")},
			{Name: "disk", Type: proto.ColumnType_INT, Description: "Current disk usage.", Transform: transform.FromField("Disk")},
			{Name: "maxdisk", Type: proto.ColumnType_INT, Description: "Maximum disk.", Transform: transform.FromField("MaxDisk")},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "Node uptime in seconds.", Transform: transform.FromField("Uptime")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Resource type.", Transform: transform.FromField("Type")},
		},
	}
}

func listProxmoxNodes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		d.StreamListItem(ctx, node)
	}
	return nil, nil
}
