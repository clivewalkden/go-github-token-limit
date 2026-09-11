BINARY  := github-token-limit
OUTDIR  := bin
VERSION ?= $(shell svu next 2>/dev/null || echo "dev")
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS  = -s -w \
           -X 'main.Version=$(VERSION)' \
           -X 'main.Commit=$(COMMIT)' \
           -X 'main.Date=$(DATE)'

.PHONY: all build snapshot changelog release finish-release lint lint-install fmt vet test test-race clean help

all: build

build: ## Build the binary into $(OUTDIR)/$(BINARY)
	@mkdir -p $(OUTDIR)
	@echo "Building $(BINARY) $(VERSION) → $(OUTDIR)/$(BINARY)"
	@go build -v -ldflags "$(LDFLAGS)" -o $(OUTDIR)/$(BINARY) ./cmd

snapshot: ## Build a local snapshot with GoReleaser (no git tag required)
	goreleaser release --snapshot --clean

changelog: ## Regenerate CHANGELOG.md from all commits
	git-cliff --output CHANGELOG.md

release: ## Cut a release branch: bumps version, updates CHANGELOG, commits, and pushes
	go mod tidy
	git flow release start $(VERSION)
	git-cliff --tag $(VERSION) --unreleased --prepend CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore: update changelog for $(VERSION)"
	git push --set-upstream origin release/$(VERSION)

finish-release: ## Finish the current git-flow release and push branches/tags
	git flow release finish --fetch
	git push -u origin --all
	git push origin --tags
	git checkout develop

lint-install: ## Install golangci-lint compiled against the local Go toolchain (required for Go 1.26+)
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

lint: fmt vet ## Run linters (run make lint-install first if golangci-lint is missing or outdated)
	@golangci-lint run ./...

fmt: ## Check gofmt
	@unfmt=$$(gofmt -l .); \
	if [ -n "$$unfmt" ]; then \
		echo "Files need gofmt:"; echo "$$unfmt"; exit 1; \
	fi

vet: ## Run go vet
	@go vet ./...

test: ## Run tests
	@go test ./...

test-race: ## Run tests with the race detector enabled
	@go test -race ./...

clean: ## Remove built artifacts
	@rm -rf $(OUTDIR)

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "Targets:\n"} /^[a-zA-Z0-9_\-]+:.*##/ { printf "  %-14s %s\n", $$1, $$2 }' Makefile
