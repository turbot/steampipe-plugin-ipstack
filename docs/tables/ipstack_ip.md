
# Table: ipstack_ip

Query location, currency, timezone and security information about any IP address.

Notes:
* Security data is only available to subscribers. Set the security config option to true to use it.

## Examples

### Query details for any IPv4 address

```sql
select * from ipstack_ip where ip = '99.84.45.75';
```

### Query details for any IPv6 address

```sql
select * from ipstack_ip where ip = '2001:4860:4860::8888';
```


### Lookup IP details for elastic IPs

```sql
select
  allocation_id,
  public_ip,
  i.city,
  i.zip,
  i.region_name,
  i.country_name,
  i.continent_name,
  i.latitude,
  i.longitude
from
  aws_vpc_eip as eip,
  ipstack_ip as i
where
  eip.public_ip = i.ip;
```
