---
title: "Steampipe Table: ipstack_my_ip - Query IPStack My IP using SQL"
description: "Allows users to query My IP details in IPStack, specifically the location data, providing insights into IP address geolocation."
---

# Table: ipstack_my_ip - Query IPStack My IP using SQL

IPStack is a service that provides geolocation and IP intelligence data. It allows you to identify user location, timezone, and ISP information using IP addresses. The IPStack service provides detailed and accurate data that can be used for a variety of purposes such as content personalization, fraud prevention, and network security.

## Table Usage Guide

The `ipstack_my_ip` table provides insights into the geolocation details of your IP address within IPStack. As a network engineer or security analyst, explore IP-specific details through this table, including location, ISP, and timezone. Utilize it to uncover information about your IP, such as its geolocation, the associated timezone, and the ISP details.

## Examples

### Get information about your current IP
Explore your current IP address details to understand its geographical location, ISP, and other related information. This can be useful for troubleshooting network issues or for security monitoring purposes.

```sql+postgres
select * from ipstack_my_ip;
```

```sql+sqlite
select * from ipstack_my_ip;
```