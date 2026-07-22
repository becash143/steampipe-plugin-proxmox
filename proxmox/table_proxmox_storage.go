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
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Storage type (dir, nfs, zfs, lvm, etc).", Transform: transform.FromField("Type")},
			{Name: "content", Type: proto.ColumnType_STRING, Description: "Allowed content types.", Transform: transform.FromField("Content")},
			{Name: "active", Type: proto.ColumnType_INT, Description: "Whether storage is active.", Transform: transform.FromField("Active")},
			{Name: "used", Type: proto.ColumnType_INT, Description: "Used space in bytes.", Transform: transform.FromField("Used")},
			{Name: "avail", Type: proto.ColumnType_INT, Description: "Available space in bytes.", Transform: transform.FromField("Avail")},
			{Name: "total", Type: proto.ColumnType_INT, Description: "Total space in bytes.", Transform: transform.FromField("Total")},
			{Name: "shared", Type: proto.ColumnType_INT, Description: "Whether storage is shared across nodes.", Transform: transform.FromField("Shared")},
		},
	}
}

func listProxmoxStorage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	storages, err := client.ListStorage()
	if err != nil {
		return nil, err
	}
	for _, s := range storages {
		d.StreamListItem(ctx, s)
	}
	return nil, nil
}
