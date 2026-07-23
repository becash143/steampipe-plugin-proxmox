package proxmox

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxUser() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_user",
		Description: "Proxmox access control users",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxUsers,
		},
		Columns: []*plugin.Column{
			{Name: "userid", Type: proto.ColumnType_STRING, Description: "User ID (e.g. user@pve).", Transform: transform.FromField("UserID")},
			{Name: "enable", Type: proto.ColumnType_INT, Description: "Whether the account is enabled.", Transform: transform.FromField("Enable")},
			{Name: "expire", Type: proto.ColumnType_TIMESTAMP, Description: "Account expiration (unix epoch, 0 = never).", Transform: transform.FromField("Expire").Transform(expireToTimestamp)},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address.", Transform: transform.FromField("Email")},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment/description.", Transform: transform.FromField("Comment")},
		},
	}
}

// expireToTimestamp converts the raw Proxmox "expire" unix-epoch value into a time.Time,
// treating 0 (Proxmox's "never expires" sentinel) as nil instead of an invalid timestamp.
func expireToTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	var epoch int64
	switch v := d.Value.(type) {
	case int64:
		epoch = v
	case int:
		epoch = int64(v)
	case float64:
		epoch = int64(v)
	default:
		// Unexpected type, don't fail the whole query
		return nil, nil
	}

	if epoch == 0 {
		return nil, nil
	}

	return time.Unix(epoch, 0), nil
}

func listProxmoxUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)
	users, err := client.ListUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		d.StreamListItem(ctx, u)
	}
	return nil, nil
}
