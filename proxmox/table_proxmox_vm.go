package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxVM() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_vm",
		Description: "Proxmox QEMU virtual machines",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxVMs,
		},
		Columns: []*plugin.Column{
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node the VM runs on.", Transform: transform.FromField("Node")},
			{Name: "vmid", Type: proto.ColumnType_INT, Description: "VM ID.", Transform: transform.FromField("VMID")},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "VM name.", Transform: transform.FromField("Name")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "VM status.", Transform: transform.FromField("Status")},
			{Name: "cpus", Type: proto.ColumnType_INT, Description: "Number of CPUs.", Transform: transform.FromField("CPUs")},
			{Name: "mem", Type: proto.ColumnType_INT, Description: "Current memory usage.", Transform: transform.FromField("Mem")},
			{Name: "maxmem", Type: proto.ColumnType_INT, Description: "Maximum memory.", Transform: transform.FromField("MaxMem")},
			{Name: "maxdisk", Type: proto.ColumnType_INT, Description: "Maximum disk.", Transform: transform.FromField("MaxDisk")},
			{Name: "uptime", Type: proto.ColumnType_INT, Description: "Uptime in seconds.", Transform: transform.FromField("Uptime")},
		},
	}
}

func listProxmoxVMs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		vms, err := client.ListVMs(node.Node)
		if err != nil {
			return nil, err
		}
		for _, vm := range vms {
			d.StreamListItem(ctx, vm)
		}
	}
	return nil, nil
}
