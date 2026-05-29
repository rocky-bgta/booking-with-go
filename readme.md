# Booking

A Go web application using the chi router.

## Prerequisites

- Go 1.21+

## Running the Application

### Option 1 — Batch script (cmd)

```bat
run.bat
```

### Option 2 — PowerShell

```powershell
.\run.bat
```

### Option 3 — Go directly (no build step)

```powershell
go run ./cmd/web/
```

The server starts on **http://localhost:8080**.

## Build Only

```powershell
go build -o bookings.exe ./cmd/web/
.\bookings.exe
```

## Tests

```powershell
# All packages
go test ./...

# Single package
go test ./internal/forms/

# Verbose
go test -v ./...
```

## Other Commands

```powershell
go fmt ./...   # format code
go vet ./...   # lint
go mod tidy    # tidy dependencies
```
