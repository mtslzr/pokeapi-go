# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Search function (can filter resource lists)

## [1.1.1] - 2019-06-22
### Added
- Coverage via Codecov
- Badges for GoReport, GoDoc, license

## [1.1.0] - 2019-06-22
### Added
- Optional parameters for Resource calls:
  - Offset (defaults to zero)
  - Limit (defaults to twenty, per API)

## [1.0.0] - 2019-06-22
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

[Unreleased]: https://github.com/mtslzr/pokeapi-go/compare/v1.1.1...HEAD
[1.1.1]: https://github.com/mtslzr/pokeapi-go/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/mtslzr/pokeapi-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/mtslzr/pokeapi-go/releases/tag/v1.0.0