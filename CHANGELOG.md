## v0.3.0 [2021-03-18]

_What's new?_

- New tables added
  - [ipstack_my_ip](https://https://hub.steampipe.io/plugins/turbot/ipstack/tables/ipstack_my_ip) ([#10](https://github.com/turbot/steampipe-plugin-ipstack/pull/10))

_Enhancements_

- Updated examples for `ipstack_ip` table ([#10](https://github.com/turbot/steampipe-plugin-ipstack/pull/10))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

## v0.2.1 [2021-02-25]

_Bug fixes_

- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)

## v0.2.0 [2021-02-18]

_What's new?_

- Added support for [connection configuration](https://github.com/turbot/steampipe-plugin-ipstack/blob/main/docs/index.md#connection-configuration) arguments for the plugin. You may specify ipstack `token` for each connection in a configuration file.
