---
Title: ipstack
Organization: Turbot HQ, Inc
iconURL: https://turbot.com/images/turbot-icon-original.png

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

## Configure API Token

ipstack requires an API token for all requests, but offers a free tier. Sign up
on the [ipstack website](https://ipstack.com) to get your free token.

Set ipstack API credentials as environment variables (Mac, Linux):

```bash
export IPSTACK_TOKEN="e0067f483763d6132d549234f8a6de22"
```

Run a query:

```bash
$ steampipe query
Welcome to Steampipe v0.0.12
Type ".inspect" for more information.
> select ip, country_code, region_name, latitude, longitude, location_calling_code from ipstack_ip;
+----------------+--------------+-------------+-------------------+--------------------+-----------------------+
|       ip       | country_code | region_name |     latitude      |     longitude      | location_calling_code |
+----------------+--------------+-------------+-------------------+--------------------+-----------------------+
| 173.54.210.135 | US           | New Jersey  | 40.72555923461914 | -74.25717163085938 |                     1 |
+----------------+--------------+-------------+-------------------+--------------------+-----------------------+
>
```

## Using HTTPS (requires ipstack subscription)

ipstack restricts HTTPS requests to subscribers. This plugin uses HTTP by
default for convenience in the free tier. If you wish to use HTTPS please
set the environment variable:

```bash
export IPSTACK_HTTPS="true"
```

