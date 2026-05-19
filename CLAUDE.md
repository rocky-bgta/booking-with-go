# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Environment

Non-essential model calls are disabled in this project.
Run this before starting: `export DISABLE_NON_ESSENTIAL_MODEL_CALLS=1`

## Compact

When compacting, preserve:
- Current task goal
- Files changed
- Commands already run
- Failing tests and exact errors
- Decisions made
- Next action list

Drop:
- Old exploration paths
- Repeated logs
- Irrelevant discussion

[//]: # (## Model)

[//]: # ()
[//]: # (Always use claude-haiku-4-5 for all tasks unless explicitly asked otherwise.)

## Commands

```bash
# Run development server (from repo root)
go run ./cmd/web/

# Build binary
go build -o booking.exe ./cmd/web/

# Run tests
go test ./...

# Run single package tests
go test ./pkg/handlers/

# Format and vet
go fmt ./...
go vet ./...

# Tidy dependencies
go mod tidy
```

The server starts on `:8080`. Templates and static files are resolved from the working directory, so commands must be run from the repo root.

## Architecture

This is a Go web application using chi router with a layered package structure.

**Request flow:**
`Chi router` → `NoSurf (CSRF)` + `SCS SessionLoad` middleware → `pkg/handlers` → `pkg/render` → response

**Key packages:**
- `cmd/web/` — entry point (`main.go`), middleware wiring (`middleware.go`), route registration (`routes.go`)
- `pkg/config/` — `AppConfig` struct that is threaded through the entire app as a dependency injection container; holds the template cache, session manager, and environment flags
- `pkg/handlers/` — `Repository` struct holds a pointer to `AppConfig`; all HTTP handlers are methods on `Repository`; the package-level `Repo` variable is set via `NewRepo`/`NewHandlers`
- `pkg/render/` — template rendering; reads from `AppConfig.TemplateCache` when `app.UseCache = true`, otherwise rebuilds the cache from disk on every request (development mode)
- `models/` — shared data structs passed to templates (`TemplateData`)

**Configuration flags in `AppConfig`:**
- `InProduction bool` — controls cookie security; set to `false` in `main.go`
- `UseCache bool` — `false` in development (templates reloaded per request), `true` in production

**Template naming conventions:**
- Pages: `*.page.gohtml`
- Layouts: `*.layout.gohtml`
- The render package globs for these patterns to build the template cache

**Session management** uses `alexedwards/scs/v2` with a 24-hour lifetime and is initialized in `main.go`, stored on `AppConfig.Session`.