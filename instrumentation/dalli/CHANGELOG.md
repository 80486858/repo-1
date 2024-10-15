# Release History: opentelemetry-instrumentation-dalli

## [0.26.0](https://github.com/80486858/repo-1/compare/opentelemetry-instrumentation-dalli-v0.25.3...opentelemetry-instrumentation-dalli/v0.26.0) (2024-10-15)


### ⚠ BREAKING CHANGES

* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389))
* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3))

### Features

* add dalli obfuscation for db_statement ([#1013](https://github.com/80486858/repo-1/issues/1013)) ([a17f05b](https://github.com/80486858/repo-1/commit/a17f05be19228a8b33a58ecd899daddfc27284be))
* add db.operation attribute for dalli ([#458](https://github.com/80486858/repo-1/issues/458)) ([f631a26](https://github.com/80486858/repo-1/commit/f631a26222b5cc3b20224c081b6d594568089044))
* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389)) ([233dfd0](https://github.com/80486858/repo-1/commit/233dfd0dae81346e9687090f9d8dfb85215e0ba7))
* Validate Using Enums ([#1132](https://github.com/80486858/repo-1/issues/1132)) ([7cd4b10](https://github.com/80486858/repo-1/commit/7cd4b10ba516cecbb15a40dbe3bd5ed3860b1f88))


### Bug Fixes

* `format_command`'s terrible performance ([#207](https://github.com/80486858/repo-1/issues/207)) ([950c7b2](https://github.com/80486858/repo-1/commit/950c7b283a0343fef6a3396c9eb542e173b9e3da))
* Base config options ([#499](https://github.com/80486858/repo-1/issues/499)) ([7304e86](https://github.com/80486858/repo-1/commit/7304e86e9a3beba5c20f790b256bbb54469411ca))
* broken test file requirements ([#1286](https://github.com/80486858/repo-1/issues/1286)) ([3ec7d8a](https://github.com/80486858/repo-1/commit/3ec7d8a456dbd3c9bbad7b397a3da8b8a311d8e3))
* regex non-match with obfuscation limit (issue [#486](https://github.com/80486858/repo-1/issues/486)) ([#488](https://github.com/80486858/repo-1/issues/488)) ([6a9c330](https://github.com/80486858/repo-1/commit/6a9c33088c6c9f39b2bc30247a3ed825553c07d4))
* resolve Dalli::Server deprecation in 3.0+ ([#1015](https://github.com/80486858/repo-1/issues/1015)) ([9428b4e](https://github.com/80486858/repo-1/commit/9428b4e7df915f6bf062445fa6de7fb921cc7a44))


### Code Refactoring

* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3)) ([3e85d44](https://github.com/80486858/repo-1/commit/3e85d4436d338f326816c639cd2087751c63feb1))

### v0.25.3 / 2024-06-18

* FIXED: Relax otel common gem constraints

### v0.25.2 / 2024-05-09

* FIXED: Untrace entire request

### v0.25.1 / 2024-04-30

* FIXED: Bundler conflict warnings

### v0.25.0 / 2023-10-16

* BREAKING CHANGE: Obfuscation for mysql2, dalli and postgresql as default option for db_statement

* ADDED: Obfuscation for mysql2, dalli and postgresql as default option for db_statement

### v0.24.2 / 2023-07-21

* ADDED: Update `opentelemetry-common` from [0.19.3 to 0.20.0](https://github.com/open-telemetry/opentelemetry-ruby-contrib/pull/537)

### v0.24.1 / 2023-06-05

* FIXED: Base config options 

### v0.24.0 / 2023-05-15

* ADDED: Add db.operation attribute for dalli

### v0.23.0 / 2023-04-17

* BREAKING CHANGE: Drop support for EoL Ruby 2.7 

* ADDED: Drop support for EoL Ruby 2.7 

### v0.22.2 / 2023-01-14

* DOCS: Fix gem homepage 
* DOCS: More gem documentation fixes 

### v0.22.1 / 2022-11-28

* FIXED: `format_command`'s terrible performance 

### v0.22.0 / 2022-06-09

* Upgrading Base dependency version
* FIXED: Broken test file requirements 

### v0.21.0 / 2022-05-02

* ADDED: Validate Using Enums 

### v0.20.0 / 2021-12-01

* ADDED: Add dalli obfuscation for db_statement 
* FIXED: Resolve Dalli::Server deprecation in 3.0+ 

### v0.19.1 / 2021-09-29

* (No significant changes)

### v0.19.0 / 2021-08-12

* ADDED: Add configuration options for dalli 
* DOCS: Update docs to rely more on environment variable configuration 

### v0.18.1 / 2021-06-23

* (No significant changes)

### v0.18.0 / 2021-05-21

* ADDED: Updated API depedency for 1.0.0.rc1

### v0.17.0 / 2021-04-22

* (No significant changes)

### v0.16.0 / 2021-03-17

* FIXED: Example scripts now reference local common lib
* DOCS: Replace Gitter with GitHub Discussions

### v0.15.0 / 2021-02-18

* ADDED: Add instrumentation config validation

### v0.14.0 / 2021-02-03

* (No significant changes)

### v0.13.0 / 2021-01-29

* (No significant changes)

### v0.12.0 / 2020-12-24

* (No significant changes)

### v0.11.0 / 2020-12-11

* ADDED: Add peer service config to dalli
* ADDED: Move utf8 encoding to common utils
* FIXED: Copyright comments to not reference year

### v0.10.0 / 2020-12-03

* (No significant changes)

### v0.9.0 / 2020-11-27

* BREAKING CHANGE: Add timeout for force_flush and shutdown

* ADDED: Add timeout for force_flush and shutdown

### v0.8.0 / 2020-10-27

* (No significant changes)

### v0.7.0 / 2020-10-07

* DOCS: Standardize toplevel docs structure and readme

### v0.6.0 / 2020-09-10

* Initial release.
