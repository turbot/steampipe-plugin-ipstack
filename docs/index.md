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

Download and install the latest ipstack plugin:

```bash
$ steampipe plugin install ipstack
Installing plugin ipstack...
$
```

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

ipstack requires an API token for all requests, but offers a free tier. Sign up
on the [ipstack website](https://ipstack.com) to get your free token. It looks like `e0067f483763d6132d549234f8a6de22`

### Configure API Token

The default connection. This uses standard Application Default Credentials (ADC) against the Slack

```hcl
 connection "ipstack" {
 plugin    = "ipstack"
 }
```

A connection to a specific account, using non default Credentials.

```hcl
connection "ipstack-enterprise" {
plugin    = "ipstack"
token   = "e0067f483763d6132d934864f8a6de22"
}
```

A connection to a HTPPS request, using non default Credentials.

```hcl
connection "ipstack-enterprise" {
plugin    = "ipstack"
token   = "e0067f483763d6132d934864f8a6de22"
IPSTACK_HTTPS   = "true"
}
```
