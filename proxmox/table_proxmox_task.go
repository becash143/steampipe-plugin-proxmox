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
			{Name: "upid", Type: proto.ColumnType_STRING, Description: "Unique process/task ID.", Transform: transform.FromField("UPID")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Task type.", Transform: transform.FromField("Type")},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Task status (OK, error, running).", Transform: transform.FromField("Status")},
			{Name: "user", Type: proto.ColumnType_STRING, Description: "User who initiated the task.", Transform: transform.FromField("User")},
			{Name: "pid", Type: proto.ColumnType_INT, Description: "Process ID.", Transform: transform.FromField("PID")},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "Task start time (unix epoch).", Transform: transform.FromField("StartTime")},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "Task end time (unix epoch).", Transform: transform.FromField("EndTime")},
		},
	}
}

func listProxmoxTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	config := d.Connection.Config.(Config)
	client := NewClient(config)

	nodes, err := client.ListNodes()
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		tasks, err := client.ListTasks(node.Node)
		if err != nil {
			return nil, err
		}
		for _, t := range tasks {
			d.StreamListItem(ctx, t)
		}
	}
	return nil, nil
}
