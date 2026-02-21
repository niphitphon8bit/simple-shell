# simple-shell

A minimal interactive shell written in Go.

## Features

- Interactive prompt: `hostname\username\current-dir >`
- Built-in commands: `cd <path>`, `exit`
- Executes external commands found in `PATH`
- Ignores empty input safely
- Exits cleanly on `EOF` (for example, `Ctrl+D`)

## Project Structure

```text
simple-shell/
├── cmd/
│   └── simple-shell/
│       └── main.go
├── internal/
│   ├── errors.go
│   ├── exec.go
│   └── exec_test.go
├── .gitignore
├── go.mod
└── README.md
```

## Requirements

- Go `1.25.1` (from `go.mod`)

## Run

```bash
go run ./cmd/simple-shell
```

Or build a binary:

```bash
go build -o simple-shell ./cmd/simple-shell
./simple-shell
```

## Test

```bash
go test ./...
```

With coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Example

```text
my-host\my-user\simple-shell > ls
my-host\my-user\simple-shell > cd /tmp
my-host\my-user\tmp > exit
```

## Current Limitations

- Command parsing is basic (`strings.Fields`) and does not support quoted arguments
- No pipes (`|`) or redirects (`>`, `<`) yet
- `cd` without a path does not go to `$HOME`; it returns `path required`
- Arrow-key escape sequences are currently ignored
