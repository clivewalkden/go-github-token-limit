# Check GitHub Token Limit

[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/clivewalkden/go-github-token-limit/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/clivewalkden/go-github-token-limit/?branch=main)
[![Build Status](https://scrutinizer-ci.com/g/clivewalkden/go-github-token-limit/badges/build.png?b=main)](https://scrutinizer-ci.com/g/clivewalkden/go-github-token-limit/build-status/main)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/clivewalkden/go-github-token-limit/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/clivewalkden/go-github-token-limit/tree/main)

This executable returns the number of requests remaining for the GitHub API. If there are no tokens left it gives the reset time.

The application assumes the ENV variable `GITHUB_TOKEN` is set with a valid GitHub token and uses that to return the quota data for.

## Installation

### Pre-built binaries
Pre-built binaries are available on the [releases page](https://github.com/clivewalkden/go-github-token-limit/releases/latest).

Simply download the binary for your platform and run it.

### Homebrew

Install with Homebrew on macOS (or Linux with Homebrew installed):
```shell
brew tap clivewalkden/taps
brew install github-token-limit
```

## Usage
The application is a command line application and will try to find the GitHub token in the following environment variables:
* GITHUB_TOKEN
* GITHUB_TOKEN
* GH_TOKEN
* GITHUB_ACCESS_TOKEN
* GH_ACCESS_TOKEN
* GITHUB_OAUTH_TOKEN
* GH_OAUTH_TOKEN
* GITHUB_PAT
* GH_PAT
* GITHUB_AUTH_TOKEN
* GH_AUTH_TOKEN
* GITHUB_API_TOKEN
* GH_API_TOKEN
* GITHUB_API_KEY
* GH_API_KEY
* GITHUB_PERSONAL_ACCESS_TOKEN
* GH_PERSONAL_ACCESS_TOKEN
* GITHUB_PERSONAL_TOKEN
* GH_PERSONAL_TOKEN
* GITHUB_PERSONAL_API_TOKEN
* GH_PERSONAL_API_TOKEN
* GITHUB_PERSONAL_API_KEY
* GH_PERSONAL_API_KEY
* GITHUB_APP_TOKEN
* GH_APP_TOKEN
* GITHUB_APP_KEY
* GH_APP_KEY

If the token is not found in any of these environment variables, the application will exit with an error.

![No Token](./docs/github-token-checker-no-token.png)

If the token is found, the application will check the quota and return the number of requests remaining.

![Quote remaining](./docs/github-token-checker-success.png)

If the token has no requests remaining, the application will return the reset time.

![Quote reset](./docs/github-token-checker-none-left.png)

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see
the [tags on this repository](https://github.com/clivewalkden/go-wasabi-cleanup/tags) or [CHANGELOG.md](./CHANGELOG.md).

## Authors

* **Clive Walkden** - *Initial work* - [SOZO Design Ltd](https://github.com/sozo-design)

See also the list of [contributors](https://github.com/clivewalkden/go-wasabi-cleanup/contributors) who participated in
this project.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details