package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProxmoxTask() *plugin.Table {
	return &plugin.Table{
		Name:        "proxmox_task",
		Description: "Proxmox node task/job history",
		List: &plugin.ListConfig{
			Hydrate: listProxmoxTasks,
		},
		Columns: []*plugin.Column{
			{Name: "node", Type: proto.ColumnType_STRING, Description: "Node the task ran on.", Transform: transform.FromField("Node")},
			{Name: "up_id", Type: proto.ColumnType_STRING, Description: "Unique process/task ID.", Transform: transform.FromField("UPID")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Task type.", Transform: transform.FromField("Type")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Task status (OK, error, running).", Transform: transform.FromField("Status")},
			{Name: "user_id", Type: proto.ColumnType_STRING, Description: "User who initiated the task.", Transform: transform.FromField("User")},
			{Name: "pid", Type: proto.ColumnType_INT, Description: "Process ID.", Transform: transform.FromField("PID")},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "Task start time (unix epoch).", Transform: transform.FromField("StartTime").Transform(transform.UnixToTimestamp)},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "Task end time (unix epoch, 0/unset while running).", Transform: transform.FromField("EndTime").Transform(epochToTimestampOrNil)},
		},
	}
}

// epochToTimestampOrNil converts a raw unix-epoch value to time.Time, treating
// 0 or nil as "not set" and returning nil rather than the 1970-01-01 epoch
// instant. This must check the raw numeric value BEFORE conversion -
// chaining transform.UnixToTimestamp followed by .NullIfZero() does not work,
// since NullIfZero() would then be comparing against time.Time{}'s zero value
// (0001-01-01), which a converted epoch-0 timestamp (1970-01-01) never equals.
func epochToTimestampOrNil(ctx context.Context, d *transform.TransformData) (interface{}, error) {
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
		return nil, nil
	}
	if epoch == 0 {
		return nil, nil
	}
	// Delegate to the SDK's own conversion once we know the value is real,
	// so both timestamp columns share the same underlying behavior.
	return transform.UnixToTimestamp(ctx, &transform.TransformData{Value: epoch})
}

func listProxmoxTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	client, err := connect(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_task.listProxmoxTasks", "connect_error", err)
		return nil, err
	}
	nodes, err := client.ListNodes()
	if err != nil {
		plugin.Logger(ctx).Error("proxmox_task.listProxmoxTasks", "api_error", err)
		return nil, err
	}

	for _, node := range nodes {
		tasks, err := client.ListTasks(node.Node)
		if err != nil {
			// A single unreachable or errored node shouldn't take down
			// results for every other healthy node in the cluster.
			plugin.Logger(ctx).Error("proxmox_task.listProxmoxTasks", "node", node.Node, "error", err)
			continue
		}

		for _, t := range tasks {
			d.StreamListItem(ctx, t)

			// Stop streaming early once the query's LIMIT is satisfied or
			// the context has been cancelled.
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}
