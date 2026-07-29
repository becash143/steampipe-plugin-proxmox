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
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "User ID (e.g. user@pve).", Transform: transform.FromField("UserID")},
			{Name: "is_enabled", Type: proto.ColumnType_INT, Description: "Whether the account is enabled.", Transform: transform.FromField("Enable")},
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
	client, err := connect(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_user.listProxmoxUsers", "connect_error", err)
		return nil, err
	}
	users, err := client.ListUsers()
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_user.listProxmoxUsers", "api_error", err)
		return nil, err
	}
	for _, u := range users {
		d.StreamListItem(ctx, u)

		// Stop streaming early once the query's LIMIT is satisfied or the
		// context has been cancelled.
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}
