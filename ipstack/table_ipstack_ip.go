package ipstack

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableIpstackIP() *plugin.Table {
	return &plugin.Table{
		Name:        "ipstack_ip",
		Description: "ipstack provides IP to geolocation APIs and global IP database services.",
		List: &plugin.ListConfig{
			Hydrate: listIP,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ip"),
			Hydrate:    getIP,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_STRING, Transform: transform.FromField("IP"), Description: "Requested IP address."},
			{Name: "hostname", Type: proto.ColumnType_STRING},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "IP address type IPv4 or IPv6."},
			{Name: "continent_code", Type: proto.ColumnType_STRING, Description: "2-letter continent code associated with the IP."},
			{Name: "continent_name", Type: proto.ColumnType_STRING, Description: "Name of the continent associated with the IP."},
			{Name: "country_code", Type: proto.ColumnType_STRING, Description: "2-letter country code associated with the IP."},
			{Name: "country_name", Type: proto.ColumnType_STRING, Description: "name of the country associated with the IP."},
			{Name: "region_code", Type: proto.ColumnType_STRING, Description: "Region code of the region associated with the IP (e.g. CA for California)."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Description: "Name of the region associated with the IP."},
			{Name: "city", Type: proto.ColumnType_STRING, Description: "Name of the city associated with the IP."},
			{Name: "zip", Type: proto.ColumnType_STRING, Description: "ZIP code associated with the IP."},
			{Name: "latitude", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Latitide"), Description: "Latitude value associated with the IP."},
			{Name: "longitude", Type: proto.ColumnType_DOUBLE, Description: "Longitude value associated with the IP."},

			{Name: "location_geoname_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Location.GeonameID"), Description: "Unique geoname identifier in accordance with the Geonames Registry."},
			{Name: "location_capital", Type: proto.ColumnType_STRING, Transform: transform.FromField("Location.Capital"), Description: "Capital city of the country associated with the IP."},
			{Name: "location_languages", Type: proto.ColumnType_JSON, Transform: transform.FromField("Location.Langauges"), Description: "Object containing one or multiple sub-objects per language spoken in the country associated with the IP."},
			{Name: "location_country_flag_link", Type: proto.ColumnType_STRING, Transform: transform.FromField("Location.CountryFlagLink"), Description: "HTTP URL leading to an SVG-flag icon for the country associated with the IP."},
			{Name: "location_country_flag_emoji", Type: proto.ColumnType_STRING, Transform: transform.FromField("Location.CountryFlagEmoji"), Description: "Emoji icon for the flag of the country associated with the IP."},
			{Name: "location_country_flag_emoji_unicode", Type: proto.ColumnType_STRING, Transform: transform.FromField("Location.CountryFlagEmojiUnicode"), Description: "Unicode value of the emoji icon for the flag of the country associated with the IP (e.g. U+1F1F5 U+1F1F9 for the Portuguese flag)."},
			{Name: "location_calling_code", Type: proto.ColumnType_STRING, Transform: transform.FromField("Location.CallingCode"), Description: "Calling / dial code of the country associated with the IP. (e.g. 351) for Portugal."},
			{Name: "location_is_eu", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Location.IsEU"), Description: "True or false depending on whether or not the county associated with the IP is in the European Union."},

			{Name: "timezone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Timezone.ID"), Description: "ID of the time zone associated with the IP (e.g. America/Los_Angeles for PST)."},
			{Name: "timezone_current_time", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Timezone.CurrentTime"), Description: "Current date and time in the location associated with the IP (e.g. 2018-03-29T22:31:27-07:00)."},
			{Name: "timezone_gmt_offset", Type: proto.ColumnType_INT, Transform: transform.FromField("Timezone.GMTOffset"), Description: "GMT offset of the given time zone in seconds (e.g. -25200 for PST's -7h GMT offset)."},
			{Name: "timezone_code", Type: proto.ColumnType_STRING, Transform: transform.FromField("Timezone.Code"), Description: "Universal code of the given time zone."},
			{Name: "timezone_is_daylight_saving", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Timezone.IsDaylightSaving"), Description: "True or false depending on whether or not the given time zone is considered daylight saving time."},

			{Name: "currency_code", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency.Code"), Description: "3-letter code of the main currency associated with the IP."},
			{Name: "currency_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency.Name"), Description: "Name of the given currency."},
			{Name: "currency_plural", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency.Plural"), Description: "Plural name of the given currency."},
			{Name: "currency_symbol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency.Symbol"), Description: "Symbol letter of the given currency."},
			{Name: "currency_symbol_native", Type: proto.ColumnType_STRING, Transform: transform.FromField("Currency.SymbolNative"), Description: "Native symbol letter of the given currency."},

			{Name: "asn", Type: proto.ColumnType_INT, Transform: transform.FromField("Connection.ASN"), Description: "Autonomous System Number associated with the IP."},
			{Name: "isp", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.ISP"), Description: "Name of the ISP associated with the IP."},

			{Name: "is_proxy", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Security.IsProxy"), Description: "True or false depending on whether or not the given IP is associated with a proxy."},
			{Name: "proxy_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Security.ProxyType"), Description: "Type of proxy the IP is associated with."},
			{Name: "is_crawler", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Security.IsCrawler"), Description: "True or false depending on whether or not the given IP is associated with a crawler."},
			{Name: "crawler_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Security.CrawlerName"), Description: "Name of the crawler the IP is associated with."},
			{Name: "crawler_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Security.CrawlerType"), Description: "Type of crawler the IP is associated with."},
			{Name: "is_tor", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Security.IsTOR"), Description: "True or false depending on whether or not the given IP is associated with the anonymous Tor system."},
			{Name: "threat_level", Type: proto.ColumnType_STRING, Transform: transform.FromField("Security.ThreatLevel"), Description: "Type of threat level the IP is associated with."},
			{Name: "threat_types", Type: proto.ColumnType_JSON, Transform: transform.FromField("Security.ThreatTypes"), Description: "Object containing all threat types associated with the IP."},
		},
	}
}

func listIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
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

func getIP(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	ip := quals["ip"].GetStringValue()
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("ipstack_ip.getIP", "connection_error", err)
		return nil, err
	}
	result, err := conn.IP(ip)
	if err != nil {
		plugin.Logger(ctx).Error("ipstack_ip.getIP", "query_error", err)
		return nil, err
	}
	return result, nil
}
