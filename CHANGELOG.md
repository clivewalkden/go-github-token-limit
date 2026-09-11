# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.0] - 2026-09-11

### Dependencies

- Update dependencies

### Documentation

- Add CLAUDE.md project guidance
- Add CHANGELOG.md generation via git-cliff

### Features

- Updated golang sys from version v0.18.0 to v0.21.0

### Fixes

- Corrected dependabot target-branch
- Resolve golangci-lint findings

## [1.4.0] - 2026-09-11

### Dependencies

- Update dependencies

### Documentation

- Add CLAUDE.md project guidance

### Features

- Updated golang sys from version v0.18.0 to v0.21.0

### Fixes

- Corrected dependabot target-branch
- Resolve golangci-lint findings

## [1.3.0] - 2024-06-22

### Documentation

- Added Usage instructions

### Features

- Search multiple ENV variables for the GitHub Access Token.
- Output the token used so that users can easily identify the token used.

### Fixes

- If limit of 60 returned we know a token hasn't been specified
- Added line spacing after error messages to match success line spacing.

### Tests

- Added test for return value having a limit of 60
- CircleCI code coverage tests
- Removed codecoverage for future release

## [1.2.1] - 2024-06-15

### Documentation

- Added the token name in ENV that the software is looking for
- Update Homebrew taps to new name

### Fixes

- CenterString function breaks if the length is shorter than the string.

### Tests

- Helper test added

## [1.2.0] - 2024-06-15

### Documentation

- Updated docs to include the Homebrew installation instructions

### Features

- Simplified the string centering and updated quota messages.

## [1.1.2] - 2024-06-14

### Fixes

- GoReleaser Config build for Homebrew tap

## [1.1.1] - 2024-06-14

### Fixes

- GoReleaser Config

## [1.1.0] - 2024-06-14

### Documentation

- Added build badges

### Features

- Clear screen before output is started

### Tests

- Added tests
- Added Circle-ci test config
- Fix linter reported issue
- Updated tests to use gotestsum and output results in xml as per circle-ci docs

[1.4.0]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.1.2...v1.2.0
[1.1.2]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/clivewalkden/go-github-token-limit/compare/v1.0.0...v1.1.0

