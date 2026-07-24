package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxContainer() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_container",
		Description: "Proxmox LXC containers",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxContainers,
		},
		Columns: []*plugin.Column{
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node the container runs on.", Transform: transform.FromField("Node")},
			{Name: "vm_id", Type: proto.ColumnType_INT, Description: "Container ID.", Transform: transform.FromField("VMID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Container name.", Transform: transform.FromField("Name")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Container status.", Transform: transform.FromField("Status")},
			{Name: "cpus", Type: proto.ColumnType_INT, Description: "Number of CPUs.", Transform: transform.FromField("CPUs")},
			{Name: "mem", Type: proto.ColumnType_INT, Description: "Current memory usage.", Transform: transform.FromField("Mem")},
			{Name: "max_mem", Type: proto.ColumnType_INT, Description: "Maximum memory.", Transform: transform.FromField("MaxMem")},
			{Name: "max_disk", Type: proto.ColumnType_INT, Description: "Maximum disk.", Transform: transform.FromField("MaxDisk")},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "Uptime in seconds.", Transform: transform.FromField("Uptime")},
		},
	}
}

func listProxmoxContainers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		containers, err := client.ListContainers(node.Node)
		if err != nil {
			return nil, err
		}
		for _, ct := range containers {
			d.StreamListItem(ctx, ct)
		}
	}
	return nil, nil
}
