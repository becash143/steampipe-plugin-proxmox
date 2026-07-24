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
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "Whether the interface is active.", Transform: transform.FromField("Active").Transform(intToBool)},
			{Name: "is_autostart", Type: proto.ColumnType_BOOL, Description: "Whether the interface starts on boot.", Transform: transform.FromField("Autostart").Transform(intToBool)},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "IP address.", Transform: transform.FromField("Address")},
			{Name: "netmask", Type: proto.ColumnType_STRING, Description: "Netmask.", Transform: transform.FromField("Netmask")},
			{Name: "gateway", Type: proto.ColumnType_STRING, Description: "Gateway address.", Transform: transform.FromField("Gateway")},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "Configuration method (static, dhcp, manual).", Transform: transform.FromField("Method")},
		},
	}
}

// intToBool converts Proxmox's 0/1 integer-style booleans (received as int,
// int64, or float64 depending on how the underlying client unmarshals JSON)
// into a real bool for BOOL-typed columns. Returns nil for a nil source value
// rather than false, so "unset" and "explicitly false" remain distinguishable
// in query results.
func intToBool(_ context.Context, d *transform.TransformData) (any, error) {
	switch v := d.Value.(type) {
	case int:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case nil:
		return nil, nil
	}
	return d.Value, nil
}

func listProxmoxNetworkInterfaces(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	client, err := connect(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_network.listProxmoxNetworkInterfaces", "connect_error", err)
		return nil, err
	}
	nodes, err := client.ListNodes()
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_network.listProxmoxNetworkInterfaces", "api_error", err)
		return nil, err
	}
	for _, node := range nodes {
		ifaces, err := client.ListNetworkInterfaces(node.Node)
		if err != nil {
			// A single unreachable or errored node shouldn't take down results
			// for every other healthy node in the cluster. Log and continue.
			plugin.Logger(ctx).Error("proxmox_network.listProxmoxNetworkInterfaces", "node", node.Node, "error", err)
			continue
		}
		for _, iface := range ifaces {
			d.StreamListItem(ctx, iface)

			// Stop streaming early once the query's LIMIT is satisfied or the
			// context has been cancelled.
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}
