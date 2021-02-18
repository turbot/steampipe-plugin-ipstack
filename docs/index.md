---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/ipstack.svg"
brand_color: "#FF9900"
display_name: "ipstack"
name: "ipstack"
description: "Steampipe plugin for querying location, currency, timezone and security information about an IP address from ipstack."
---

# ipstack

ipstack provides IP to geolocation APIs and global IP database services.

## Installation

To download and install the latest ipstack plugin:

```bash
$ steampipe plugin install ipstack
Installing plugin ipstack...
$
```

## Credentials

ipstack requires an API token for all requests, but offers a free tier. Sign up on the [ipstack website](https://ipstack.com) to get your free token. It looks like `e0067f483763d6132d549234f8a6de22`.


## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest ipstack plugin will create a default connection named `ipstack` in the `~/.steampipe/config/ipstack.spc` file.  You must edit this connection to include your API token:

```hcl
connection "ipstack" {
  plugin  = "ipstack"
  token   = "e0067f483763d6132d934864f8a6de22"
}
```

ipstack restricts HTTPS requests to subscribers. This plugin uses HTTP by default for convenience in the free tier. If you wish to use HTTPS please set the `https` argument:
```hcl
connection "ipstack" {
  plugin  = "ipstack"
  token   = "e0067f483763d6132d934864f8a6de22"
  https   = "true"
}
```
