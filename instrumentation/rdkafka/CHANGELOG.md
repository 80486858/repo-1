# Release History: opentelemetry-instrumentation-rdkafka

## [0.5.0](https://github.com/80486858/repo-1/compare/opentelemetry-instrumentation-rdkafka-v0.4.6...opentelemetry-instrumentation-rdkafka/v0.5.0) (2024-10-15)


### ⚠ BREAKING CHANGES

* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389))
* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3))

### Features

* Adds instrumentation for rdkafka ([#978](https://github.com/80486858/repo-1/issues/978)) ([a84067b](https://github.com/80486858/repo-1/commit/a84067bb6c8404a4c784b7291e16985ab859010d))
* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389)) ([233dfd0](https://github.com/80486858/repo-1/commit/233dfd0dae81346e9687090f9d8dfb85215e0ba7))


### Bug Fixes

* Base config options ([#499](https://github.com/80486858/repo-1/issues/499)) ([7304e86](https://github.com/80486858/repo-1/commit/7304e86e9a3beba5c20f790b256bbb54469411ca))
* broken test file requirements ([#1286](https://github.com/80486858/repo-1/issues/1286)) ([3ec7d8a](https://github.com/80486858/repo-1/commit/3ec7d8a456dbd3c9bbad7b397a3da8b8a311d8e3))
* regex non-match with obfuscation limit (issue [#486](https://github.com/80486858/repo-1/issues/486)) ([#488](https://github.com/80486858/repo-1/issues/488)) ([6a9c330](https://github.com/80486858/repo-1/commit/6a9c33088c6c9f39b2bc30247a3ed825553c07d4))
* skip recording non-utf8 kafka keys in racecar and rdkafka ([#392](https://github.com/80486858/repo-1/issues/392)) ([d5a7487](https://github.com/80486858/repo-1/commit/d5a74878e657efad2f6de6d5bc6dc25db0b631e3))


### Code Refactoring

* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3)) ([3e85d44](https://github.com/80486858/repo-1/commit/3e85d4436d338f326816c639cd2087751c63feb1))

### v0.4.7 / 2024-07-02

* DOCS: Fix CHANGELOGs to reflect a past breaking change

### v0.4.6 / 2024-06-18

* FIXED: Relax otel common gem constraints

### v0.4.5 / 2024-05-09

* FIXED: Untrace entire request

### v0.4.4 / 2024-04-30

* FIXED: Bundler conflict warnings

### v0.4.3 / 2024-04-05

* FIXED: Suppress deprecation warning in Rdkafka Instrumentation

### v0.4.2 / 2023-11-23

* FIXED: Retry Release of 0.4.1 [#730](https://github.com/open-telemetry/opentelemetry-ruby-contrib/issues/730)

### v0.4.1 / 2023-11-22

* FIXED: Get Rdkafka version from VERSION contant

### v0.4.0 / 2023-09-07

* BREAKING CHANGE: Align messaging instrumentation operation names [#648](https://github.com/open-telemetry/opentelemetry-ruby-contrib/pull/648)

### v0.3.2 / 2023-07-21

* ADDED: Update `opentelemetry-common` from [0.19.3 to 0.20.0](https://github.com/open-telemetry/opentelemetry-ruby-contrib/pull/537)

### v0.3.1 / 2023-06-05

* FIXED: Base config options 

### v0.3.0 / 2023-04-17

* BREAKING CHANGE: Drop support for EoL Ruby 2.7 

* ADDED: Drop support for EoL Ruby 2.7 

### v0.2.3 / 2023-03-24

* FIXED: Skip recording non-utf8 kafka keys in racecar and rdkafka

### v0.2.2 / 2023-01-14

* DOCS: Fix gem homepage 
* DOCS: More gem documentation fixes 

### v0.2.1 / 2022-11-14

* Loosen dependency on opentelemetry-api

### v0.2.0 / 2022-06-09

* Upgrading Base dependency version
* FIXED: Broken test file requirements

### v0.1.0 / 2022-05-02

* Initial release.
