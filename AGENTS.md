# Agent Instructions

## Build/Lint/Test Commands

- `make all` – Build openapi and lambdas
- `make openapi-generated` – Generate OpenAPI spec from TypeSpec
- `make lambdas` – Build all Lambda binaries
- `make server` – Run local dev server (adapters/inbound/local)
- `make test` – Run all core tests with coverage (generates mocks first)
- Single test: `cd core && go test ./usecases -run TestName`
- Coverage check: `make check-coverage` (100% thresholds)
- Regenerate mocks: `go generate ./core/mocks`
- Format: `go fmt ./...` and `go vet ./...` (no explicit lint target)

## Code Style Guidelines

- Go 1.24.4, workspace with go.work
- Imports: standard, external, internal groups
- Naming: PascalCase exported, camelCase private, suffixes like `Usecases`, `Handler`
- Error handling: Often ignored (`_`) in adapters; use `respondWithJSON` helper
- Use repository interface with generics (Repository[T])
- Testing: gomock for mocks, factory pattern (testhelpers/factories.go)
- Coverage: 100% required for files, packages, total (enforced by .testcoverage.yaml)
- No Cursor/Copilot rules found
- File naming: lowercase, underscores only in directory names (e.g., `non-electronic-games`)