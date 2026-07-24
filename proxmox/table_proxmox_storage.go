package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxStorage() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_storage",
		Description: "Proxmox storage pools",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxStorage,
		},
		Columns: []*plugin.Column{
			{Name: "storage", Type: proto.ColumnType_STRING, Description: "Storage identifier.", Transform: transform.FromField("Storage")},
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node this storage status was reported from.", Transform: transform.FromField("Node")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Storage type (dir, nfs, zfs, lvm, etc).", Transform: transform.FromField("Type")},
			{Name: "content", Type: proto.ColumnType_STRING, Description: "Allowed content types.", Transform: transform.FromField("Content")},
			{Name: "is_active", Type: proto.ColumnType_BOOL, Description: "Whether the storage pool is active.", Transform: transform.FromField("Active").Transform(intToBool)},
			{Name: "used", Type: proto.ColumnType_INT, Description: "Used space in bytes.", Transform: transform.FromField("Used")},
			{Name: "avail", Type: proto.ColumnType_INT, Description: "Available space in bytes.", Transform: transform.FromField("Avail")},
			{Name: "total", Type: proto.ColumnType_INT, Description: "Total space in bytes.", Transform: transform.FromField("Total")},
			{Name: "is_shared", Type: proto.ColumnType_BOOL, Description: "Whether the storage pool is shared across nodes.", Transform: transform.FromField("Shared").Transform(intToBool)},
		},
	}
}

func listProxmoxStorage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	client, err := connect(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_storage.listProxmoxStorage", "connect_error", err)
		return nil, err
	}
	nodes, err := client.ListNodes()
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_storage.listProxmoxStorage", "api_error", err)
		return nil, err
	}
	for _, n := range nodes {
		storages, err := client.ListStorageStatus(n.Node)
		if err != nil {
			// A single unreachable or errored node shouldn't take down results
			// for every other healthy node in the cluster. Log and continue.
			plugin.Logger(ctx).Error("proxmox_storage.listProxmoxStorage", "node", n.Node, "error", err)
			continue
		}
		for _, s := range storages {
			d.StreamListItem(ctx, s)

			// Stop streaming early once the query's LIMIT is satisfied or the
			// context has been cancelled.
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}
