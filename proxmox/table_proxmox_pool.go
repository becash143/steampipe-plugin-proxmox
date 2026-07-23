package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxPool() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_pool",
		Description: "Proxmox resource pools",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxPools,
		},
		Columns: []*plugin.Column{
			{Name: "poolid", Type: proto.ColumnType_STRING, Description: "Pool identifier.", Transform: transform.FromField("PoolID")},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Pool comment/description.", Transform: transform.FromField("Comment")},
		},
	}
}

func listProxmoxPools(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	pools, err := client.ListPools()
	if err != nil {
		return nil, err
	}
	for _, p := range pools {
		d.StreamListItem(ctx, p)
	}
	return nil, nil
}
