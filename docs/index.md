---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/ipstack.svg"
brand_color: "#FF9900"
display_name: "ipstack"
name: "ipstack"
description: "Steampipe plugin for querying location, currency, timezone and security information about an IP address from ipstack."
og_description: "Query ipstack with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/ipstack-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"] 
---

# ipstack + Steampipe

[ipstack](https://ipstack.com/) provides IP to geolocation APIs and global IP database services.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  ip,
  country_code,
  region_name,
  latitude
from
  ipstack_ip
where
  ip = '99.84.45.75';
```

```
+-------------+--------------+-------------+--------------------+
| ip          | country_code | region_name | latitude           |
+-------------+--------------+-------------+--------------------+
| 99.84.45.75 | US           | New Jersey  | 40.738731384277344 |
+-------------+--------------+-------------+--------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/ipstack/tables)**

## Get started

### Install

Download and install the latest ipstack plugin:

```bash
steampipe plugin install ipstack
```

### Credentials

ipstack requires an API token for all requests, but offers a free tier. Sign up on the [ipstack website](https://ipstack.com) to get your free token. It looks like `e0067f483763d6132d549234f8a6de22`.

### Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest ipstack plugin will create a default connection named `ipstack` in the `~/.steampipe/config/ipstack.spc` file.  You must edit this connection to include your API token:

```hcl
connection "ipstack" {
  plugin = "ipstack"
  token  = "e0067f483763d6132d934864f8a6de22"
}
```

ipstack restricts HTTPS access and security data to subscribers. To enable:
```hcl
connection "ipstack" {
  plugin   = "ipstack"
  token    = "e0067f483763d6132d934864f8a6de22"
  https    = true
  security = true
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-ipstack
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
