![image](https://hub.steampipe.io/images/plugins/turbot/ipstack-social-graphic.png)

# ipstack Plugin for Steampipe

Use SQL to query IP address information including geolocation and more from ipstack.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/ipstack)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/ipstack/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-ipstack/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install ipstack
```

Run a query:

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
> select ip, country_code, region_name, latitude from ipstack_ip where ip = '99.84.45.75' ;
+-------------+--------------+-------------+--------------------+
| ip          | country_code | region_name | latitude           |
+-------------+--------------+-------------+--------------------+
| 99.84.45.75 | US           | New Jersey  | 40.738731384277344 |
+-------------+--------------+-------------+--------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-ipstack.git
cd steampipe-plugin-ipstack
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/ipstack.spc
```

Try it!

```
steampipe query
> .inspect ipstack
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-ipstack/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [ipstack Plugin](https://github.com/turbot/steampipe-plugin-ipstack/labels/help%20wanted)