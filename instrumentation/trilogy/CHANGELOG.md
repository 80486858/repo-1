# Release History: opentelemetry-instrumentation-trilogy

## [0.60.0](https://github.com/80486858/repo-1/compare/opentelemetry-instrumentation-trilogy-v0.59.3...opentelemetry-instrumentation-trilogy/v0.60.0) (2024-10-15)


### ⚠ BREAKING CHANGES

* Separate logical MySQL host from connected host ([#487](https://github.com/80486858/repo-1/issues/487))
* Add database name for trilogy traces ([#484](https://github.com/80486858/repo-1/issues/484))
* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389))
* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3))

### Features

* Add database name for trilogy traces ([#484](https://github.com/80486858/repo-1/issues/484)) ([e8f57c6](https://github.com/80486858/repo-1/commit/e8f57c6ea4d29a0ea9a61ad9352970c29a1d8516))
* Add Obfuscation Limit Option to Trilogy ([#477](https://github.com/80486858/repo-1/issues/477)) ([234738c](https://github.com/80486858/repo-1/commit/234738c5fbd8d630d543f61d84fcefcf948756f1))
* add option to configure span name for trilogy ([#348](https://github.com/80486858/repo-1/issues/348)) ([518e108](https://github.com/80486858/repo-1/commit/518e108c08486f5755d637d63bc0f7771889f271))
* Add Trilogy Auto Instrumentation ([#1063](https://github.com/80486858/repo-1/issues/1063)) ([8261f22](https://github.com/80486858/repo-1/commit/8261f2286f6379061ee389e538d6310648cf18be))
* add with_attributes context propagation to Trilogy instrumentation ([#347](https://github.com/80486858/repo-1/issues/347)) ([64a2b10](https://github.com/80486858/repo-1/commit/64a2b101c901117d8555708a1ffbbb6bab0376b3))
* Drop support for EoL Ruby 2.7 ([#389](https://github.com/80486858/repo-1/issues/389)) ([233dfd0](https://github.com/80486858/repo-1/commit/233dfd0dae81346e9687090f9d8dfb85215e0ba7))
* Separate logical MySQL host from connected host ([#487](https://github.com/80486858/repo-1/issues/487)) ([f6df818](https://github.com/80486858/repo-1/commit/f6df818e7e92b1357314dae0f9c1b0877d04878e))


### Bug Fixes

* Base config options ([#499](https://github.com/80486858/repo-1/issues/499)) ([7304e86](https://github.com/80486858/repo-1/commit/7304e86e9a3beba5c20f790b256bbb54469411ca))
* broken test file requirements ([#1286](https://github.com/80486858/repo-1/issues/1286)) ([3ec7d8a](https://github.com/80486858/repo-1/commit/3ec7d8a456dbd3c9bbad7b397a3da8b8a311d8e3))
* Ensure encoding errors handled during SQL obfuscation for Trilogy ([#345](https://github.com/80486858/repo-1/issues/345)) ([1a5972f](https://github.com/80486858/repo-1/commit/1a5972f449e920bd3b54fc1033121d72f906c771))
* regex non-match with obfuscation limit (issue [#486](https://github.com/80486858/repo-1/issues/486)) ([#488](https://github.com/80486858/repo-1/issues/488)) ([6a9c330](https://github.com/80486858/repo-1/commit/6a9c33088c6c9f39b2bc30247a3ed825553c07d4))
* regex non-match with obfuscation limit (issue [#486](https://github.com/80486858/repo-1/issues/486)) ([#488](https://github.com/80486858/repo-1/issues/488)) ([b31a4cb](https://github.com/80486858/repo-1/commit/b31a4cbb20ba7ee4a3422ce65f948a7fa3f43f85))
* Remove inline linter rules ([#608](https://github.com/80486858/repo-1/issues/608)) ([bc4a937](https://github.com/80486858/repo-1/commit/bc4a937ed2a0d1898f0f19ae45a2b3a0ef9a067c))
* Trilogy Driver Options ([#1080](https://github.com/80486858/repo-1/issues/1080)) ([01b6790](https://github.com/80486858/repo-1/commit/01b679017cf8a5327c79aae06912cb784f53c381))


### Code Refactoring

* Remove parent repo libraries ([#3](https://github.com/80486858/repo-1/issues/3)) ([3e85d44](https://github.com/80486858/repo-1/commit/3e85d4436d338f326816c639cd2087751c63feb1))

### v0.59.3 / 2024-04-30

* FIXED: Bundler conflict warnings

### v0.59.2 / 2024-02-20

* FIXED: Dup string if frozen in trilogy query

### v0.59.1 / 2024-02-08

* FIXED: Add missing requires for sql-helpers to mysql, pg, and trilogy instrumentation

### v0.59.0 / 2024-02-08

* BREAKING CHANGE: Move shared sql behavior to helper gems

* ADDED: Propagate context to Vitess

### v0.58.0 / 2024-01-06

* BREAKING CHANGE: Change db.mysql.instance.address to db.instance.id

* ADDED: Change db.mysql.instance.address to db.instance.id
* FIXED: Trilogy only set db.instance.id attribute if there is a value

### v0.57.0 / 2023-10-27

* ADDED: Instrument connect and ping

### v0.56.3 / 2023-08-03

* FIXED: Remove inline linter rules

### v0.56.2 / 2023-07-14

* ADDED: `db.user` attribute (recommended connection-level attribute)

### v0.56.1 / 2023-06-05

* FIXED: Base config options 

### v0.56.0 / 2023-06-02

* BREAKING CHANGE: Separate logical MySQL host from connected host 

* ADDED: Separate logical MySQL host from connected host 

### v0.55.1 / 2023-06-01

* FIXED: Regex non-match with obfuscation limit (issue #486) 

### v0.55.0 / 2023-05-31

* BREAKING CHANGE: Add database name for trilogy traces 

* ADDED: Add database name for trilogy traces 

### v0.54.0 / 2023-05-25

* ADDED: Add Obfuscation Limit Option to Trilogy 

### v0.53.0 / 2023-04-17

* BREAKING CHANGE: Drop support for EoL Ruby 2.7 

* ADDED: Drop support for EoL Ruby 2.7 

### v0.52.0 / 2023-03-06

* ADDED: Add with_attributes context propagation to Trilogy instrumentation 
* ADDED: Add option to configure span name for trilogy 
* FIXED: Ensure encoding errors handled during SQL obfuscation for Trilogy 

### v0.51.1 / 2023-01-14

* DOCS: Fix gem homepage 
* DOCS: More gem documentation fixes 

### v0.51.0 / 2022-06-09

* Upgrading Base dependency version
* FIXED: Broken test file requirements 

### v0.50.2 / 2022-05-05

* (No significant changes)

### v0.50.1 / 2022-01-07

* FIXED: Trilogy Driver Options 

### v0.50.0 / 2021-12-31

* Initial release.
