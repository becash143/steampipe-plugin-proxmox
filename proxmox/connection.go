package proxmox

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// connectMemoized wraps connectUncached so the built client is cached in the
// connection's ConnectionCache rather than rebuilt on every hydrate call.
// Per the SDK docs, Memoize'd functions must be invoked manually (via
// connect(), below) rather than assigned directly as a table's Hydrate.
var connectMemoized = plugin.HydrateFunc(connectUncached).Memoize(memoize.WithCacheKeyFunction(connectCacheKey))

// connectCacheKey scopes the cached client to the current connection name, so
// multiple Steampipe connections against different Proxmox clusters each get
// their own client rather than sharing one.
func connectCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return "proxmox_client." + d.Connection.Name, nil
}

// connectUncached builds a new API client from the connection config. Do not
// call this directly from table code - use connect() so the client is cached.
func connectUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := d.Connection.Config.(Config)
	return NewClient(config), nil
}

// connect returns the cached *Client for the current connection, building
// one only on the first call for that connection.
func connect(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (*Client, error) {
	clientRaw, err := connectMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return clientRaw.(*Client), nil
}
