# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.0] / 2020.05.12
### Added
- User settings for caching (pokeapi.CacheSettings)
  - CustomExpire: set a custom cache expiration (in minutes)
  - UseCache: turn caching on/off

## [1.3.2] / 2020.04.03
### Added
- Current version to README.
### Fixed
- Missing version link in CHANGELOG.

## [1.3.1] / 2020.04.03
### Changed
- Replaced CircleCI with GitHub Actions

## [1.3.0] / 2019.06.23
### Added
- Caching via [go-cache](https://github.com/patrickmn/go-cache)
- ClearCache function (to clear existing cache)
  
## [1.2.1] / 2019.06.23
### Changed
- Added additional unit tests for Search

## [1.2.0] / 2019.06.22
### Added
- Search function for filtering resource lists
  - Ability to filter by any string
  - Ability to filter by "starts with" using ^

## [1.1.1] / 2019.06.22
### Added
- Coverage via Codecov
- Badges for GoReport, GoDoc, license

## [1.1.0] / 2019.06.22
### Added
- Optional parameters for Resource calls:
  - Offset (defaults to zero)
  - Limit (defaults to twenty, per API)

## [1.0.0] / 2019.06.22
### Added
- Endpoint Categories
  - Berries (berries, firmness, flavors)
  - Contests (types, effects, etc.)
  - Encounters (methods, conditions, etc.)
  - Evolution (chains, triggers)
  - Games (generations, versions, etc.)
  - Locations (areas, regions, etc.)
  - Machines
  - Moves (moves, attributes, etc.)
  - Pokemon (abilities, egg groups, etc.)
  - Utility (languages)
- Resource endpoint for resource lists
- Unit tests for all endpoints

[Unreleased]: https://github.com/mtslzr/pokeapi-go/compare/v1.4.0...HEAD
[1.4.0]: https://github.com/mtslzr/pokeapi-go/compare/v1.3.2...v1.4.0
[1.3.2]: https://github.com/mtslzr/pokeapi-go/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/mtslzr/pokeapi-go/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/mtslzr/pokeapi-go/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/mtslzr/pokeapi-go/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/mtslzr/pokeapi-go/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/mtslzr/pokeapi-go/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/mtslzr/pokeapi-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/mtslzr/pokeapi-go/releases/tag/v1.0.0