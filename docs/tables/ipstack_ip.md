---
title: "Steampipe Table: ipstack_ip - Query IP Geolocation using SQL"
description: "Allows users to query IP Geolocation, specifically providing detailed information about any IP address such as location, connection, and security details."
---

# Table: ipstack_ip - Query IP Geolocation using SQL

IP Geolocation is a service within that allows you to get detailed information about any IP address. It provides location details such as country, region, city, zip code, latitude, and longitude. Additionally, it offers connection information like ISP and ASN, as well as security data including whether the IP is a proxy, crawler, or threat.

## Table Usage Guide

The `ipstack_ip` table provides insights into IP addresses using IP Geolocation. As a network administrator or cybersecurity professional, you can explore detailed information about any IP address, including its location, connection, and security details. Use this table to gain insights into network traffic patterns, identify potential security threats, or troubleshoot network issues.

**Important Notes**
- Security data is only available to subscribers. Set the security config option to true to use it.

## Examples

### Query details for any IPv4 address
Explore the geographical and connection details for a specific IPv4 address. This can be useful in understanding the origin of web traffic or troubleshooting network issues.

```sql+postgres
select * from ipstack_ip where ip = '99.84.45.75';
```

```sql+sqlite
select * from ipstack_ip where ip = '99.84.45.75';
```

### Query details for any IPv6 address
Gain insights into the geographical and network-related information for a specific IPv6 address. This can be useful for network troubleshooting, security audits, or understanding user demographics.

```sql+postgres
select * from ipstack_ip where ip = '2001:4860:4860::8888';
```

```sql+sqlite
select * from ipstack_ip where ip = '2001:4860:4860::8888';
```


### Lookup IP details for elastic IPs
Analyze the geographical details of Elastic IP addresses to understand their distribution worldwide. This is beneficial for assessing the spread of your network resources and identifying potential areas of vulnerability or opportunity.

```sql+postgres
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

```sql+sqlite
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