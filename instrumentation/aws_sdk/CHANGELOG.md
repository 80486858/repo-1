# Release History: opentelemetry-instrumentation-aws_sdk

## [0.6.0](https://github.com/80486858/repo-1/compare/opentelemetry-instrumentation-aws_sdk-v0.5.2...opentelemetry-instrumentation-aws_sdk/v0.6.0) (2024-10-15)


### ⚠ BREAKING CHANGES

* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389))
* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3))

### Features

* aws-sdk instrumentation ([#1014](https://github.com/80486858/repo-1/issues/1014)) ([a37637c](https://github.com/80486858/repo-1/commit/a37637c241cacc1799dd0728248c2d1372986348))
* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389)) ([233dfd0](https://github.com/80486858/repo-1/commit/233dfd0dae81346e9687090f9d8dfb85215e0ba7))
* SQS / SNS messaging attributes and context propagation ([#1026](https://github.com/80486858/repo-1/issues/1026)) ([658440f](https://github.com/80486858/repo-1/commit/658440f2f749f151fde393ce1654803a816f1c38))


### Bug Fixes

* aws-sdk nitpicks ([#1023](https://github.com/80486858/repo-1/issues/1023)) ([b375af1](https://github.com/80486858/repo-1/commit/b375af1c2b737601bc3ec7037e6261470eb7f6d3))
* Base config options ([#499](https://github.com/80486858/repo-1/issues/499)) ([7304e86](https://github.com/80486858/repo-1/commit/7304e86e9a3beba5c20f790b256bbb54469411ca))
* broken test file requirements ([#1286](https://github.com/80486858/repo-1/issues/1286)) ([3ec7d8a](https://github.com/80486858/repo-1/commit/3ec7d8a456dbd3c9bbad7b397a3da8b8a311d8e3))
* regex non-match with obfuscation limit (issue [#486](https://github.com/80486858/repo-1/issues/486)) ([#488](https://github.com/80486858/repo-1/issues/488)) ([6a9c330](https://github.com/80486858/repo-1/commit/6a9c33088c6c9f39b2bc30247a3ed825553c07d4))
* Remove inline linter rules ([#608](https://github.com/80486858/repo-1/issues/608)) ([bc4a937](https://github.com/80486858/repo-1/commit/bc4a937ed2a0d1898f0f19ae45a2b3a0ef9a067c))
* RubyGems Fallback ([#1161](https://github.com/80486858/repo-1/issues/1161)) ([3b03ff7](https://github.com/80486858/repo-1/commit/3b03ff7ea66b69c85ba205a369b85c2c33b712fe))
* suppress invalid span attribute value type warning in aws-sdk instrumentation ([#64](https://github.com/80486858/repo-1/issues/64)) ([4db6509](https://github.com/80486858/repo-1/commit/4db6509a63e1dca36b492e0dd6c82e7bfae72562))


### Code Refactoring

* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3)) ([3e85d44](https://github.com/80486858/repo-1/commit/3e85d4436d338f326816c639cd2087751c63feb1))

### v0.5.3 / 2024-07-02

* DOCS: Fix CHANGELOGs to reflect a past breaking change

### v0.5.2 / 2024-04-30

* FIXED: Bundler conflict warnings

### v0.5.1 / 2024-02-08

* FIXED: Return nil for non-existant key in AwsSdk::MessageAttributeGetter

### v0.5.0 / 2023-09-07

* BREAKING CHANGE: Align messaging instrumentation operation names [#648](https://github.com/open-telemetry/opentelemetry-ruby-contrib/pull/648)

### v0.4.2 / 2023-08-03

* FIXED: Remove inline linter rules

### v0.4.1 / 2023-06-05

* FIXED: Base config options 

### v0.4.0 / 2023-04-17

* BREAKING CHANGE: Drop support for EoL Ruby 2.7 

* ADDED: Drop support for EoL Ruby 2.7 

### v0.3.2 / 2023-01-14

* DOCS: Fix gem homepage 
* DOCS: More gem documentation fixes 

### v0.3.1 / 2022-07-19

* FIXED: Suppress invalid span attribute value type warning in aws-sdk instrumentation 

### v0.3.0 / 2022-06-09

* Upgrading Base dependency version
* FIXED: Broken test file requirements 

### v0.2.3 / 2022-05-02

* FIXED: RubyGems Fallback 

### v0.2.2 / 2022-01-26

* (No significant changes)

### v0.2.1 / 2022-01-21

* ADDED: attach HTTP status code to AWS spans

### v0.2.0 / 2022-01-20

* ADDED: SQS / SNS messaging attributes and context propagation

### v0.1.0 / 2021-12-01

* Initial release.
