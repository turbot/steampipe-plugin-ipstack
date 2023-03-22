package ipstack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableIpstackMyIP() *plugin.Table {
	return &plugin.Table{
		Name:        "ipstack_my_ip",
		Description: "IP and geolocation information about your public IP.",
		List: &plugin.ListConfig{
			Hydrate: listMyIP,
		},

		Columns: ipstackIpColumns(),
	}
}

func listMyIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("ipstack_ip.listIP", "connection_error", err)
		return nil, err
	}
	result, err := conn.Me()
	if err != nil {
		plugin.Logger(ctx).Error("ipstack_ip.listIP", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
