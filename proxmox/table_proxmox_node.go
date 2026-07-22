package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableProxmoxNode() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_node",
		Description: "Proxmox cluster nodes",

		List: &plugin.ListConfig{
			Hydrate: listProxmoxNodes,
		},

		Columns: []*plugin.Column{
			{
				Name:        "node",
				Type:        proto.ColumnType_STRING,
				Description: "Proxmox node name.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Node status.",
			},
			{
				Name:        "cpu",
				Type:        proto.ColumnType_DOUBLE,
				Description: "Current CPU usage ratio.",
			},
			{
				Name:        "maxcpu",
				Type:        proto.ColumnType_INT,
				Description: "Number of CPUs.",
			},
			{
				Name:        "mem",
				Type:        proto.ColumnType_INT,
				Description: "Current memory usage.",
			},
			{
				Name:        "maxmem",
				Type:        proto.ColumnType_INT,
				Description: "Maximum memory.",
			},
			{
				Name:        "disk",
				Type:        proto.ColumnType_INT,
				Description: "Current disk usage.",
			},
			{
				Name:        "maxdisk",
				Type:        proto.ColumnType_INT,
				Description: "Maximum disk.",
			},
			{
				Name:        "uptime",
				Type:        proto.ColumnType_INT,
				Description: "Node uptime in seconds.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Resource type.",
			},
		},
	}
}

func listProxmoxNodes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {

	// Load connection configuration
	config := d.Connection.Config.(*Config)

	// Create Proxmox API client
	client := NewClient(*config)

	// Get nodes from Proxmox API
	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}

	// Stream nodes as table rows
	for _, node := range nodes {
		d.StreamListItem(ctx, node)
	}

	return nil, nil
}
