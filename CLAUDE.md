# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A small Go CLI (`github-token-limit`) that checks the GitHub API rate limit for a token found in the environment, printing remaining requests or the reset time if exhausted.

## Commands

```sh
make fmt              # go fmt ./...
make vet              # go vet ./... (runs fmt first)
make lint             # golangci-lint run ./... (runs fmt first)
make build            # cross-compiles to bin/<VERSION>/ for freebsd/darwin/linux/windows (runs vet first)
go test ./...         # run all tests
go test ./cmd -run TestFetchRateLimit -v   # run a single test
```

`make build` enforces a minimum Go version via `MINIMUM_SUPPORTED_GO_MINOR_VERSION` in the Makefile — keep this in sync with the `go` directive in `go.mod` when bumping the Go version.

There is no `main` package build target for local iteration other than `make build` (which builds every platform); for a quick local build use `go build ./cmd/main.go`.

## Architecture

- `cmd/main.go` — entry point. Clears the screen, resolves a token via `githubapi.GetGithubTokenFromEnv()`, calls `githubapi.FetchRateLimit()`, and prints results via `internal/utils` notice helpers. Exit code `3` on fetch error or when GitHub's unauthenticated default limit (60) is detected, signalling no valid token was supplied.
- `internal/githubapi/githubapi.go` — all GitHub API interaction:
  - `TokenEnvNames` is an exhaustive list of environment variable names checked (in order) for a token — `GITHUB_TOKEN`, `GH_TOKEN`, and many token-name variants. `GetGithubTokenFromEnv()` returns the first non-empty one, or `""`.
  - `APIURL` is a package-level `var`, not a constant, specifically so tests can redirect it at an `httptest` server.
  - `Timestamp` wraps `time.Time` with a custom `UnmarshalJSON` that accepts either an RFC3339Nano string or a numeric Unix timestamp, since GitHub's rate-limit JSON has varied historically.
  - `FetchRateLimit(client, token)` does the HTTP round trip and JSON decode; it takes an `*http.Client` as a parameter (not a package global) for testability.
- `internal/utils/messaging.go` — colored, centered console notices (`InfoNotice`, `SuccessNotice`, `CautionNotice`, `ErrorNotice`) built on `github.com/fatih/color`, all centered to 80 columns via `helper.go`'s `CenterString`.
- `internal/utils/helper.go` — `CenterString` (pads a string to a fixed width) and `ObscureToken` (masks a token for display, showing first 14 and last 4 chars).

Tests mock the GitHub API by swapping `githubapi.APIURL` to point at an `httptest.NewServer`, then restoring it via `defer`. Follow this pattern for any new tests against `FetchRateLimit`.

## Release tooling

Version is injected at build time via `-ldflags -X main.version=...` (see `Makefile`'s `FLAGS`, using a hardcoded `VERSION` var) and also via GoReleaser (`.goreleaser.yaml`) using `{{.Version}}` from git tags — these are two separate versioning mechanisms, not yet unified. GoReleaser builds cross-platform archives and publishes to `clivewalkden/homebrew-taps`. CI runs in both GitHub Actions (`.github/workflows/ci.yml`: lint → test → semantic-release+goreleaser on push) and CircleCI (`.circleci/config.yml`: test only, via `gotestsum`).
