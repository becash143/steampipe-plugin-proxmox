package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxNetwork() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_network",
		Description: "Proxmox node network interfaces",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxNetworkInterfaces,
		},
		Columns: []*plugin.Column{
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node the interface belongs to.", Transform: transform.FromField("Node")},
			{Name: "iface", Type: proto.ColumnType_STRING, Description: "Interface name.", Transform: transform.FromField("Iface")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Interface type (bridge, bond, vlan, etc).", Transform: transform.FromField("Type")},
			{Name: "active", Type: proto.ColumnType_INT, Description: "Whether interface is active.", Transform: transform.FromField("Active")},
			{Name: "autostart", Type: proto.ColumnType_INT, Description: "Whether interface starts on boot.", Transform: transform.FromField("Autostart")},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "IP address.", Transform: transform.FromField("Address")},
			{Name: "netmask", Type: proto.ColumnType_STRING, Description: "Netmask.", Transform: transform.FromField("Netmask")},
			{Name: "gateway", Type: proto.ColumnType_STRING, Description: "Gateway address.", Transform: transform.FromField("Gateway")},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "Configuration method (static, dhcp, manual).", Transform: transform.FromField("Method")},
		},
	}
}

func listProxmoxNetworkInterfaces(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		ifaces, err := client.ListNetworkInterfaces(node.Node)
		if err != nil {
			return nil, err
		}
		for _, iface := range ifaces {
			d.StreamListItem(ctx, iface)
		}
	}
	return nil, nil
}
