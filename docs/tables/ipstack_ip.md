
# Table: ipstack_ip

Query location, currency, timezone and security information about an IP address.

## Examples

### Get information about your current IP

```sql
select * from ipstack_ip;
```

### Query details for any IPv4 address

```sql
select * from ipstack_ip where ip = '99.84.45.75';
```

### Query details for any IPv6 address

```sql
select * from ipstack_ip where ip = '2001:0db8:85a3:0000:0000:8a2e:0370:7334';
```
