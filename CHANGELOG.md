## v0.9.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#45](https://github.com/turbot/steampipe-plugin-ipstack/pull/45))

## v0.9.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#43](https://github.com/turbot/steampipe-plugin-ipstack/pull/43))
- Recompiled plugin with Go version `1.21`. ([#43](https://github.com/turbot/steampipe-plugin-ipstack/pull/43))

## v0.8.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#37](https://github.com/turbot/steampipe-plugin-ipstack/pull/37))

## v0.7.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#34](https://github.com/turbot/steampipe-plugin-ipstack/pull/34))
- Recompiled plugin with Go version `1.19`. ([#34](https://github.com/turbot/steampipe-plugin-ipstack/pull/34))

## v0.6.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#29](https://github.com/turbot/steampipe-plugin-ipstack/pull/29))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#28](https://github.com/turbot/steampipe-plugin-ipstack/pull/28))

## v0.5.1 [2022-02-02]

_Bug fixes_

- Add the missing social image definition in the docs/index.md file ([#25](https://github.com/turbot/steampipe-plugin-ipstack/pull/25))

## v0.5.0 [2022-01-31]

_Warning_
- `https` option is now bool, not a string.
- Removed `IPSTACK_HTTPS` env var.

_Enhancements_
- Add `security` config (default false), since security data now requires a subscription.
- Recompiled plugin with [steampipe-plugin-sdk v1.8.3](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v183--2021-12-23) ([#22](https://github.com/turbot/steampipe-plugin-ipstack/pull/22))
- `docs/index.md` and `README.md` files have now been updated to the latest format

## v0.4.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#18](https://github.com/turbot/steampipe-plugin-ipstack/pull/18))
- Recompiled plugin with Go version 1.17 ([#18](https://github.com/turbot/steampipe-plugin-ipstack/pull/18))

## v0.3.1 [2021-10-06]

_Enhancements_

- Updated the README file to include GitHub clone link and license information ([#15](https://github.com/turbot/steampipe-plugin-ipstack/pull/15))
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)([#12](https://github.com/turbot/steampipe-plugin-ipstack/pull/12))

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
